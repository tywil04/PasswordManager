<script>
    import * as validations from "../../lib/validations.js"
    import * as utils from "../../lib/utils.js"

    import * as base64 from "base64-arraybuffer"

    import PrimaryButton from "../../components/buttons/PrimaryButton.svelte";
    import RegularButton from "../../components/buttons/RegularButton.svelte";
    import EmailInput from "../../components/inputs/EmailInput.svelte";
    import PasswordInput from "../../components/inputs/PasswordInput.svelte";
    import TextInput from "../../components/inputs/TextInput.svelte";

    let visibleForm = "webauthnChallenge"
    let visibleFormData = ""
    let submitError

    async function signin(e) {
        submitError = undefined
        data = new FormData(e.target)

        const email = data.get("email")
        const password = data.get("password")

        if (!validations.validateEmail(email) || !validations.validatePassword(password)) {
            return
        }

        let masterKey = await crypto.generateMasterKey(password, email)
        let masterHash = await crypto.generateMasterHash(password, masterKey) 

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
            visibleForm = json.challengeType
            visibleFormData = json[visibleForm + "Id"]

            const databaseKey = await crypto.unprotectDatabaseKey(masterKey, {
                key: base64.decode(json.ProtectedDatabaseKey.Key),
                iv: base64.decode(json.ProtectedDatabaseKey.Iv),
            })

            const exportedDatabaseKey = await crypto.exportKey(databaseKey)
            sessionStorage.setItem("PasswordManager4:databaseKey", base64.encode(exportedDatabaseKey))
        }
    }

    async function emailChallenge(e) {
        submitError = undefined
        data = new FormData(e.target)

        const code = data.get("code")

        if (code === "" || code === null || code === undefined) {
            return
        }

        const response = await fetch("/api/v1/email/signinChallenge", {
            method: "POST",
            headers: {
                "Content-type": "application/json",
            },
            body: JSON.stringify({
                "emailChallengeId": visibleFormData,
                "code": code,
            })
        })
        const json = await response.json()

        if (response.status !== 200) {
            sessionStorage.setItem("PasswordManager4:databaseKey", "")
            submitError = json.error.code
        } else if (response.status === 200) {
            sessionStorage.setItem("PasswordManager4:authToken", json.authToken)
            window.location = "/"
        }
    }

    async function webauthnChallenge() {
        submitError = undefined

        const response = await fetch(`/api/v1/webauthn/signinChallenge?webauthnChallengeId=${visibleFormData}`, {
            method: "GET",
            headers: {
                "Content-type": "application/json",
            },
        })
        const json = await response.json()

        if (response.status !== 200) {
            sessionStorage.setItem("PasswordManager4:databaseKey", "")
            submitError = json.error.code
        } else if (response.status === 200) {
            console.log(json)

            const webauthnChallengeId = json.webauthnChallengeId
            const credential = await navigator.credentials.get(json.options)
            const jsonCredential = {
                authenticatorAttachment: credential.authenticatorAttachment,
                id: credential.id,
                rawId: credential.id,
                response: {
                    authenticatorData: utils.base64ToUrlBase64(base64.encode(credential.response.authenticatorData)).replaceAll("=", ""),
                    clientDataJSON: utils.base64ToUrlBase64(base64.encode(credential.response.clientDataJSON)).replaceAll("=", ""),
                    signature: utils.base64ToUrlBase64(base64.encode(credential.response.signature)).replaceAll("=", ""),
                },
                type: credential.type,
            }

            const finishResponse = await fetch("/api/v1/webauthn/signinChallenge", {
                method: "POST",
                headers: {
                    "Content-type": "application/json",
                },
                body: JSON.stringify({
                    webauthnChallengeId,
                    credential: jsonCredential,
                })
            })
            const finishJson = await finishResponse.json()

            if (finishResponse.status !== 200) {
                sessionStorage.setItem("PasswordManager4:databaseKey", "")
                submitError = json.error.code
            } else if (finishResponse.status === 200) {
                sessionStorage.setItem("PasswordManager4:authToken", finishJson.authToken)
                window.location = "/"
            }
        }
    }
</script>

<main class="w-full h-full bg-blue-400 flex flex-col justify-center">
    <div class="w-full h-fit flex flex-col justify-center space-y-5 p-16 md:p-0 md:space-x-5 md:space-y-0 md:flex-row">    
        {#if visibleForm == "signin"}
            <div class="bg-white border border-black w-full h-fit md:w-fit duration-100 rounded-2xl">
                <h1 class="text-xl font-bold h-fit px-5 py-2.5 bg-black text-white rounded-t-[15px]">Sign in</h1>
                <p class="p-5 m-0 font-sans w-full md:max-w-[300px] text-sm">
                    To sign up, click the <span class="text-gray-500">Sign up</span> button to be redirected to the correct page.
                    <br/>
                    <br/>
                    To sign in, enter your details and click the <span class="text-blue-500">Sign in</span> button.
                </p>
            </div>

            <form method="POST" class="bg-white border border-black w-full h-fit md:w-fit duration-100 p-5 rounded-2xl space-y-5 min-w-[30%]" on:submit|preventDefault={signin}>
                <EmailInput class="flex-grow" label="Email" name="email" description="Enter your email address."/>
                <PasswordInput class="flex-grow" label="Password" name="password" description="Enter a secure password."/>

                {#if submitError !== undefined}
                    <div class="text-red-500 text-sm">• {submitError}</div>
                {/if}

                <div class="flex flex-row space-x-5">
                    <PrimaryButton class="flex-grow" submit>Sign in</PrimaryButton>  
                    <RegularButton class="flex-grow" href="#/auth/signup">Sign up</RegularButton>        
                </div>
            </form>
        {:else if visibleForm == "emailChallenge"}
            <div class="bg-white border border-black w-full h-fit md:w-fit duration-100 rounded-2xl">
                <h1 class="text-xl font-bold h-fit px-5 py-2.5 bg-black text-white rounded-t-[15px]">Email Verification</h1>
                <p class="p-5 m-0 font-sans w-full md:max-w-[300px] text-sm">
                    A code has been sent to your email address so we can verify you.
                    <br/>
                    <br/>
                    Enter the code provided in the email and click the <span class="text-blue-500">Verify</span> button.
                </p>
            </div>

            <form method="POST" class="bg-white border border-black w-full h-fit md:w-fit duration-100 p-5 rounded-2xl space-y-5 min-w-[30%]" on:submit|preventDefault={emailChallenge}>
                <TextInput class="flex-grow" label="Code" name="code" description="Enter the code from your email."/>

                {#if submitError !== undefined}
                    <div class="text-red-500 text-sm">• {submitError}</div>
                {/if}

                <div class="flex flex-row space-x-5">
                    <PrimaryButton class="flex-grow" submit>Verify</PrimaryButton>         
                </div>
            </form>
        {:else if visibleForm == "webauthnChallenge"}
            <div class="bg-white border border-black w-full h-fit md:w-fit duration-100 rounded-2xl">
                <h1 class="text-xl font-bold h-fit px-5 py-2.5 bg-black text-white rounded-t-[15px]">Webauthn Verification</h1>
                <p class="p-5 m-0 font-sans w-full md:max-w-[300px] text-sm">
                    To start Webauthn verification click the <span class="text-blue-500">Start Webauthn</span> button.
                </p>
            </div>

            <form method="POST" class="bg-white border border-black w-full h-fit md:w-fit duration-100 p-5 rounded-2xl space-y-5 min-w-[30%]" on:submit|preventDefault={webauthnChallenge}>
                {#if submitError !== undefined}
                    <div class="text-red-500 text-sm">• {submitError}</div>
                {/if}

                <div class="flex flex-row space-x-5">
                    <PrimaryButton class="flex-grow" submit>Start Webauthn</PrimaryButton>         
                </div>
            </form>
        {/if}
    </div>
</main>