<script>
    import { enhance } from "$app/forms"
    import { goto } from "$app/navigation"

    import * as base64 from "base64-arraybuffer"
    import * as webauthnJson from "@github/webauthn-json/browser-ponyfill"

    import * as cryptography from "$lib/js/cryptography.js"
    import * as validations from "$lib/js/validations.js"
    import * as storage from "$lib/js/storage.js"

    import RegularButton from "$lib/components/buttons/RegularButton.svelte"
    import PrimaryButton from "$lib/components/buttons/PrimaryButton.svelte"
    import EmailInput from "$lib/components/inputs/EmailInput.svelte"
    import PasswordInput from "$lib/components/inputs/PasswordInput.svelte"
    import TextInput from "$lib/components/inputs/TextInput.svelte"

    let view = "signin"
    let viewData = ""
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
        } else if (response.status === 200) {
            view = json.challengeType
            viewData = json[view + "Id"]
        }

        cancel()
    }

    async function emailChallenge({ data, cancel }) {
        submitError = undefined

        const code = data.get("code")

        if (code === "" || code === null || code === undefined) {
            cancel()
            return
        }

        const response = await fetch("/api/v1/email/signinChallenge", {
            method: "POST",
            headers: {
                "Content-type": "application/json",
            },
            body: JSON.stringify({
                "emailChallengeId": viewData,
                "code": code,
            })
        })
        const json = await response.json()

        if (response.status !== 200) {
            submitError = json.error.code
        } else if (response.status === 200) {
            const databaseKey = await cryptography.unprotectDatabaseKey(masterKey, {
                key: base64.decode(json.protectedDatabaseKey),
                iv: base64.decode(json.protectedDatabaseKeyIv),
            })
            const exportedDatabaseKey = await cryptography.exportKey(databaseKey)

            storage.setDatabaseKey(exportedDatabaseKey)
            storage.setAuthToken(json.authToken)

            goto("/home")
        }

        cancel()
    }

    async function webauthnChallenge({ cancel }) {
        submitError = undefined

        const response = await fetch(`/api/v1/webauthn/signinChallenge?webauthnChallengeId=${viewData}`, {
            method: "GET",
            headers: {
                "Content-type": "application/json",
            },
        })
        const json = await response.json()

        if (response.status !== 200) {
            submitError = json.error.code
        } else if (response.status === 200) {
            const webauthnChallengeId = json.webauthnChallengeId
            const credential = await webauthnJson.get(webauthnJson.parseRequestOptionsFromJSON(json.options))

            const finishResponse = await fetch("/api/v1/webauthn/signinChallenge", {
                method: "POST",
                headers: {
                    "Content-type": "application/json",
                },
                body: JSON.stringify({
                    webauthnChallengeId,
                    credential: credential,
                })
            })
            const finishJson = await finishResponse.json()

            if (finishResponse.status !== 200) {
                submitError = json.error.code
            } else if (finishResponse.status === 200) {
                const databaseKey = await cryptography.unprotectDatabaseKey(masterKey, {
                    key: base64.decode(finishJson.protectedDatabaseKey),
                    iv: base64.decode(finishJson.protectedDatabaseKeyIv),
                })
                const exportedDatabaseKey = await cryptography.exportKey(databaseKey)

                storage.setDatabaseKey(exportedDatabaseKey)
                storage.setAuthToken(finishJson.authToken)

                goto("/home")
            }
        }

        cancel()
    }
</script>

<svelte:head>
    <title>Auth</title>
</svelte:head>

<main>
    <div class="outer">
        {#if view === "signin"}
            <div class="inner">
                <h1>Sign in</h1>
                <p>
                    To sign up, click the <span class="text-gray-500">Sign up</span> button to be redirected to the correct page.
                    <br/>
                    <br/>
                    To sign in, enter your details and click the <span class="text-blue-500">Sign in</span> button.
                </p>
            </div>
        
            <form method="POST" class="inner space-y-5" use:enhance={signin}>
                <EmailInput classList="flex-grow" label="Email" name="email" description="Enter your email address."/>
                <PasswordInput classList="flex-grow" label="Password" name="password" description="Enter a secure password."/>

                {#if submitError !== undefined}
                    <div class="text-red-500 text-sm">• {submitError}</div>
                {/if}

                <div class="flex flex-row space-x-5">
                    <PrimaryButton class="flex-grow" submit variant="accent">Sign in</PrimaryButton>  
                    <RegularButton class="flex-grow" href="/auth/signup">Sign up</RegularButton>          
                </div>
            </form>
        {:else if view === "emailChallenge"}
            <div class="inner">
                <h1>Email Verification</h1>
                <p>
                    A code has been sent to your email address so we can verify you.
                    <br/>
                    <br/>
                    Enter the code provided in the email and click the <span class="text-blue-500">Verify</span> button.
                </p>
            </div>
        
            <form method="POST" class="inner space-y-5" use:enhance={emailChallenge}>
                <TextInput class="flex-grow" label="Code" name="code" description="Enter the code from your email."/>

                {#if submitError !== undefined}
                    <div class="text-red-500 text-sm">• {submitError}</div>
                {/if}

                <div class="flex flex-row space-x-5">
                    <PrimaryButton class="flex-grow" submit variant="accent">Verify</PrimaryButton>  
                </div>
            </form>
        {:else if view === "webauthnChallenge"}
            <div class="inner">
                <h1>Webauthn Verification</h1>
                <p>
                    To start Webauthn verification click the <span class="text-blue-500">Start Webauthn</span> button.
                </p>
            </div>
        
            <form method="POST" class="inner space-y-5" use:enhance={webauthnChallenge}>
                {#if submitError !== undefined}
                    <div class="text-red-500 text-sm">• {submitError}</div>
                {/if}

                <div class="flex flex-row space-x-5">
                    <PrimaryButton class="flex-grow" submit variant="accent">Start Webauthn</PrimaryButton>  
                </div>
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
        @apply bg-white border border-black w-full h-fit md:w-fit duration-100 p-5 rounded-2xl min-w-[25%];
    }

    h1 {
        @apply text-xl font-bold h-fit px-5 py-2.5 bg-black text-white rounded-t-[15px];
    }

    p {
        @apply p-5 m-0 font-sans w-full md:max-w-[300px] text-sm;
    }
</style>