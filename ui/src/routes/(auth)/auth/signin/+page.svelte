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
    <title>Password Manager - Sign in</title>
</svelte:head>

<main>
    <div class="outer">
        {#if view === "default"}
            <div class="inner">
                <p>
                    To sign up, click the <span class="text-gray-500">Sign up</span> button to be redirected to the correct page.
                    <br/>
                    <br/>
                    To sign in, enter your details and click the <span class="text-blue-500">Sign in</span> button.
                </p>
            </div>
        
            <form method="POST" class="inner space-y-5" use:enhance={signin}>
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
                    <div class="text-red-500 text-sm">• {submitError}</div>
                {/if}

                <div class="flex flex-row space-x-5">
                    <Button tabIndex="10" class="flex-grow" variant="accent" type="submit">Sign in</Button>  
                    <Button tabIndex="10" class="flex-grow" href="/auth/signup">Sign up</Button>          
                </div>
            </form>
        {:else if view === "emailChallenge"}
            <div class="inner">
                <p>
                    A code has been sent to your email address so we can verify you.
                    <br/>
                    <br/>
                    Enter the code provided in the email and click the <span class="text-blue-500">Verify</span> button.
                </p>
            </div>
        
            <form method="POST" class="inner space-y-5" use:enhance={emailChallenge}>
                <TextInput tabIndex="10" class="flex-grow" label="Code" name="code" description="Enter the code from your email."/>

                {#if submitError !== undefined}
                    <div class="text-red-500 text-sm">• {submitError}</div>
                {/if}

                <div class="flex flex-row space-x-5">
                    <Button tabIndex="10" class="flex-grow" variant="accent" type="submit">Verify</Button>  
                </div>
            </form>
        {:else if view === "totpChallenge"}
            <div class="inner">
                <p>
                    A code has been sent to your email address so we can verify you.
                    <br/>
                    <br/>
                    Enter the code provided in the email and click the <span class="text-blue-500">Verify</span> button.
                </p>
            </div>
        
            <form method="POST" class="inner space-y-5" use:enhance={totpChallenge}>
                <TextInput tabIndex="10" class="flex-grow" label="Code" name="code" description="Enter the code from your email."/>

                {#if submitError !== undefined}
                    <div class="text-red-500 text-sm">• {submitError}</div>
                {/if}

                <div class="flex flex-row space-x-5">
                    <Button tabIndex="10" class="flex-grow" variant="accent" type="submit">Verify</Button>  
                </div>
            </form>
        {:else if view === "webauthnChallenge"}
            <div class="inner">
                <p>
                    To start Webauthn verification click the <span class="text-blue-500">Start Webauthn</span> button.
                </p>
            </div>
        
            <form method="POST" class="inner space-y-5" use:enhance={webauthnChallenge}>
                {#if submitError !== undefined}
                    <div class="text-red-500 text-sm">• {submitError}</div>
                {/if}

                <div class="flex flex-row space-x-5">
                    <Button tabIndex="10" class="flex-grow" variant="accent" type="submit">Start Webauthn</Button>  
                </div>
            </form>
        {:else if view === "challengeSelector"}
            <div class="inner">
                <p>
                    Select a 2FA method.
                </p>
            </div>
        
            <form method="POST" class="inner space-y-5" on:submit|preventDefault>
                {#if availableChallenges.indexOf("email") !== -1}
                    <Button tabIndex="10" class="flex-grow" variant="accent" type="submit" on:click={startEmailChallenge}>Email</Button>
                {/if}

                {#if availableChallenges.indexOf("totp") !== -1}
                    <Button tabIndex="10" class="flex-grow" variant="accent" type="submit" on:click={startTotpChallenge}>Email</Button>
                {/if}

                {#if availableChallenges.indexOf("webauthn") !== -1}
                    <Button tabIndex="10" class="flex-grow" variant="accent" type="submit" on:click={startWebauthnChallenge}>Email</Button>
                {/if}
            </form>
        {/if}
    </div>
</main>

<style lang="postcss">
    main {
        @apply w-full h-full bg-blue-400 flex flex-col justify-center;
    }

    div.outer {
        @apply w-full h-fit flex flex-col justify-center space-y-5 p-16 md:p-0 md:space-x-5 md:space-y-0 md:flex-row;
    }

    div.inner {
        @apply bg-white border border-black w-full h-fit md:w-fit duration-100 rounded-2xl;
    }

    form.inner {
        @apply bg-white border border-black md:w-fit h-fit duration-100 p-5 rounded-2xl w-full min-w-[25%];
    }

    p {
        @apply p-5 m-0 font-sans w-full md:max-w-[300px] text-sm;
    }
</style>