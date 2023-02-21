<script>
    import { enhance } from "$app/forms"
    import { goto } from "$app/navigation"

    import * as base64 from "base64-arraybuffer"
    import * as webauthnJson from "@github/webauthn-json/browser-ponyfill"
    import { Envelope, Key } from "svelte-heros-v2"

    import * as cryptography from "$lib/js/cryptography.js"
    import * as validations from "$lib/js/validations.js"
    import * as storage from "$lib/js/storage.js"
    import * as utils from "$lib/js/utils.js"

    import Button from "$lib/components/buttons/Button.svelte"
    import EmailInput from "$lib/components/inputs/EmailInput.svelte"
    import PasswordInput from "$lib/components/inputs/PasswordInput.svelte"
    import TextInput from "$lib/components/inputs/TextInput.svelte"

    let view = "default"
    let availableChallenges
    let challengeId

    let masterKey
    let submitError 

    async function signin({ data, cancel }) {
        submitError = undefined

        const email = data.get("email")
        const password = data.get("password")

        if (!validations.validateEmail(email) || !validations.validatePassword(password)) {
            cancel()
            return
        }

        masterKey = await cryptography.generateMasterKey(password, email)
        let masterHash = await cryptography.generateMasterHash(password, masterKey) 

        const response = await fetch("/api/v1/auth/signin", {
            method: "POST",
            headers: {
                "Content-type": "application/json",
            },
            body: JSON.stringify({
                email: email,
                masterHash: base64.encode(masterHash),
            })
        })
        const json = await response.json()

        if (response.status !== 200) {
            submitError = json.error.code
        } else if (response.status === 403) {
            submitError = "Incorrect credentials"
        } else if (response.status === 200) {
            availableChallenges = json.availableChallenges
            challengeId = json.challengeId

            if (availableChallenges.length === 1 && availableChallenges[0] === "email") {
                startEmailChallenge()
            } else {
                view = "challengeSelector"
            }
        }

        cancel()
    }


    //


    const startEmailChallenge = async () => {
        submitError = undefined
        view = "emailChallenge"
        await utils.getJson(`/api/v1/email/challenge?challengeId=${challengeId}`)
    } 

    const emailChallenge = async ({ data, cancel }) => {
        const code = data.get("code")

        if (code === "" || code === null || code === undefined) {
            cancel()
            return
        }

        const postResponse = await fetch("/api/v1/email/challenge", {
            method: "POST",
            headers: {
                "Content-type": "application/json",
            },
            body: JSON.stringify({
                "challengeId": challengeId,
                "code": code,
            })
        })
        const postJson = await postResponse.json()

        if (postResponse.status !== 200 && postResponse.status !== 403) {
            submitError = postJson.error.code
        } else if (postResponse.status === 403) {
            submitError = "Incorrect code"
        } else if (postResponse.status === 200) {
            const databaseKey = await cryptography.unprotectDatabaseKey(masterKey, {
                key: base64.decode(postJson.protectedDatabaseKey),
                iv: base64.decode(postJson.protectedDatabaseKeyIv),
            })
            const exportedDatabaseKey = await cryptography.exportKey(databaseKey)

            storage.setDatabaseKey(exportedDatabaseKey)
            storage.setAuthToken(postJson.authToken)

            goto("/home")
        }

        cancel()
    }


    //
     

    const startTotpChallenge = async () => {
        submitError = undefined
        view = "totpChallenge"
    }

    const totpChallenge = async ({ data, cancel }) => {
        submitError = undefined

        const code = data.get("code")

        if (code === "" || code === null || code === undefined) {
            cancel()
            return
        }

        const postResponse = await fetch("/api/v1/totp/challenge", {
            method: "POST",
            headers: {
                "Content-type": "application/json",
            },
            body: JSON.stringify({
                "challengeId": challengeId,
                "code": code,
            })
        })
        const postJson = await postResponse.json()

        if (postResponse.status !== 200 && postResponse.status !== 403) {
            submitError = postJson.error.code
        } else if (postResponse.status === 403) {
            submitError = "Incorrect code"
        } else if (postResponse.status === 200) {
            const databaseKey = await cryptography.unprotectDatabaseKey(masterKey, {
                key: base64.decode(postJson.protectedDatabaseKey),
                iv: base64.decode(postJson.protectedDatabaseKeyIv),
            })
            const exportedDatabaseKey = await cryptography.exportKey(databaseKey)

            storage.setDatabaseKey(exportedDatabaseKey)
            storage.setAuthToken(postJson.authToken)

            goto("/home")
        }

        cancel()
    }


    //


    const startWebauthnChallenge = async () => {
        submitError = undefined
        view = "webauthnChallenge"
    }

    async function webauthnChallenge({ cancel }) {
        submitError = undefined

        const getResponse = await fetch(`/api/v1/webauthn/challenge?challengeId=${challengeId}`, {
            method: "GET",
            headers: {
                "Content-type": "application/json",
            },
        })
        const getJson = await getResponse.json()

        if (getResponse.status !== 200) {
            submitError = getJson.error.code
        } else if (getResponse.status === 200) {
            const credential = await webauthnJson.get(webauthnJson.parseRequestOptionsFromJSON(getJson.options))
            const postResponse = await fetch("/api/v1/webauthn/challenge", {
                method: "POST",
                headers: {
                    "Content-type": "application/json",
                },
                body: JSON.stringify({
                    challengeId: challengeId,
                    credential: credential,
                })
            })
            const postJson = await postResponse.json()

            if (postResponse.status !== 200 && postResponse.status !== 403) {
                submitError = postJson.error.code
            } else if (postResponse.status === 403) {
                submitError = "Invalid webauthn"
            } else if (postResponse.status === 200) {
                const databaseKey = await cryptography.unprotectDatabaseKey(masterKey, {
                    key: base64.decode(postJson.protectedDatabaseKey),
                    iv: base64.decode(postJson.protectedDatabaseKeyIv),
                })
                const exportedDatabaseKey = await cryptography.exportKey(databaseKey)

                storage.setDatabaseKey(exportedDatabaseKey)
                storage.setAuthToken(postJson.authToken)

                goto("/home")
            }
        }

        cancel()
    }
</script>

<svelte:head>
    <title>Password Manager - Signin</title>
</svelte:head>

<main>
    <div class="outer">
        {#if view === "default"}
            <div class="innerLeft">
                Signin
            </div>

            <div class="innerRight">
                <form method="POST" class="innerForm" use:enhance={signin}>
                    <EmailInput tabIndex="10" class="flex-grow" label="Email" name="email" description="Enter your email address." invalidMsg="Enter a valid email address.">
                        <svelte:fragment slot="left">
                            <Envelope size="30"strokeWidth="1"/>
                        </svelte:fragment>
                    </EmailInput>
                    <PasswordInput tabIndex="10" class="flex-grow" label="Password" name="password" description="Enter a secure password." invalidMsg="Enter a valid secure password.">
                        <svelte:fragment slot="left">
                            <Key size="30"strokeWidth="1"/>
                        </svelte:fragment>
                    </PasswordInput>
    
                    {#if submitError !== undefined}
                        <div class="errorText">• {submitError}</div>
                    {/if}
    
                    <div class="buttonRow">
                        <Button tabIndex="10" class="flex-grow" variant="accent" type="submit">Sign in</Button>  
                        <Button tabIndex="10" class="flex-grow" href="/auth/signup">Sign up</Button>          
                    </div>
                </form>
            </div>
        {:else if view === "emailChallenge"}
            <div class="innerLeft">
                Email Verification
            </div>
        
            <div class="innerRight">
                <p class="additionalText">
                    A verification email has been sent to your inbox with a code, please enter the code and click verify.
                </p>

                <form method="POST" class="innerForm" use:enhance={emailChallenge}>
                    <TextInput tabIndex="10" class="flex-grow" label="Code" name="code" description="Enter the code from your email."/>
    
                    {#if submitError !== undefined}
                        <div class="errorText">• {submitError}</div>
                    {/if}
    
                    <div class="buttonRow">
                        <Button tabIndex="10" class="flex-grow" variant="accent" type="submit">Verify</Button>  
                    </div>
                </form>
            </div>
        {:else if view === "totpChallenge"}
            <div class="innerLeft">
                Totp Verification
            </div>
        
            <div class="innerRight">
                <p class="additionalText">
                    Use your authenticator app to generate a code. Once you have the code, please enter it and click verify.
                </p>

                <form method="POST" class="innerForm" use:enhance={totpChallenge}>
                    <TextInput tabIndex="10" class="flex-grow" label="Code" name="code" description="Enter the code from your authenticator app."/>

                    {#if submitError !== undefined}
                        <div class="errorText">• {submitError}</div>
                    {/if}

                    <div class="buttonRow">
                        <Button tabIndex="10" class="flex-grow" variant="accent" type="submit">Verify</Button>  
                    </div>
                </form>
            </div>
        {:else if view === "webauthnChallenge"}
            <div class="innerLeft">
                Webauthn Verification
            </div>
    
            <div class="innerRight">
                <p class="additionalText">
                    Click the start 2FA button to start the 2FA process. You must use a credential you have already registered.
                </p>

                {#if submitError !== undefined}
                    <div class="">• {submitError}</div>
                {/if}

                <form method="POST" class="innerForm" use:enhance={webauthnChallenge}>
                    <div class="buttonRow">
                        <Button tabIndex="10" class="flex-grow" variant="accent" type="submit">Start 2FA</Button>  
                    </div>
                </form>
            </div>
        {:else if view === "challengeSelector"}
            <div class="innerLeft">
                2FA Selector
            </div>

            <div class="innerRight">
                <p class="additionalText">
                    Choose a method of 2FA to verify yourself with.
                </p>

                <form method="POST" class="innerForm" on:submit|preventDefault>
                    <div class="selectorColumn">
                        {#if availableChallenges.indexOf("email") !== -1}
                            <div class="buttonRow">
                                <Button tabIndex="10" class="flex-grow" type="submit" on:click={startEmailChallenge}>Email</Button>
                            </div>
                        {/if}
        
                        {#if availableChallenges.indexOf("totp") !== -1}
                            <div class="buttonRow">
                                <Button tabIndex="10" class="flex-grow" type="submit" on:click={startTotpChallenge}>Authenticator App</Button>
                            </div>
                        {/if}
        
                        {#if availableChallenges.indexOf("webauthn") !== -1}
                            <div class="buttonRow">
                                <Button tabIndex="10" class="flex-grow" type="submit" on:click={startWebauthnChallenge}>Webauthn</Button>
                            </div>
                        {/if}                    
                    </div>
                </form>
            </div>
        {/if}
    </div>
</main>

<style lang="postcss">
    main {
        @apply w-full h-full bg-blue-400 flex flex-col justify-center;
    }

    .outer {
        @apply w-full h-fit flex flex-row justify-center;
    }

    .innerLeft {
        writing-mode: vertical-rl;
        @apply rotate-180 flex flex-col-reverse py-5 px-2 font-bold text-2xl text-gray-800 w-fit rounded-r-2xl border border-black bg-gray-100;
    }

    .innerRight {
        @apply bg-white border border-l-0 border-black h-full duration-100 rounded-r-2xl md:max-w-[40%] lg:max-w-[30%] w-[100%] flex flex-col;
    }

    .additionalText {
        @apply p-5 pb-2 text-sm;
    }

    :global(.innerForm > *) {
        @apply space-y-5 m-5 flex flex-col justify-end flex-grow;
    }

    .selectorColumn {
        @apply flex flex-col space-y-2;
    }

    :global(.buttonRow) {
        @apply flex flex-row;
    }

    :global(.buttonRow > *:not(:last-child)) {
        @apply mr-5;
    }

    .errorText {
        @apply text-red-500 text-sm;
    }
</style>