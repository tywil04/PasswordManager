<script>
    import { enhance } from "$app/forms"
    import { goto } from "$app/navigation"

    import * as base64 from "base64-arraybuffer"
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
    let challengeId

    let masterKey
    let submitError 

    async function signup({ data, cancel }) {
        submitError = undefined

        const email = data.get("email")
        const password = data.get("password")
        const passwordConfirm = data.get("passwordConfirm")

        if (!validations.validateEmail(email) || !validations.validatePassword(password) || password !== passwordConfirm) {
            cancel()
            return
        }

        masterKey = await cryptography.generateMasterKey(password, email)
        let masterHash = await cryptography.generateMasterHash(password, masterKey) 
        let databaseKey = await cryptography.generateDatabaseKey()
        let protectedDatabaseKey = await cryptography.protectDatabaseKey(masterKey, databaseKey)

        const response = await fetch("/api/v1/auth/signup", {
            method: "POST",
            headers: {
                "Content-type": "application/json",
            },
            body: JSON.stringify({
                email: email,
                masterHash: base64.encode(masterHash),
                protectedDatabaseKey:  base64.encode(protectedDatabaseKey.key),
                protectedDatabaseKeyIv: base64.encode(protectedDatabaseKey.iv),
            })
        })
        const json = await response.json()

        if (response.status !== 200) {
            submitError = json.error.code
        } else if (response.status === 200) {
            challengeId = json.challengeId
            await startEmailChallenge()
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
</script>

<svelte:head>
    <title>Password Manager - Sign up</title>
</svelte:head>

<main>
    <div class="outer">
        {#if view === "default"}
            <div class="inner">
                <p>
                    To sign in, click the <span class="text-gray-500">Sign in</span> button to be redirected to the correct page.
                    <br/><br/>
                    To sign up, ender your details and click the <span class="text-blue-500">Sign up</span> button. 
                    <br/><br/>
                    Your password must have at least a single lowercase letter, uppercase letter, numbers and special character.
                    <br/><br/>
                    Allowed special characters:
                    <br/>
                    <code class="text-gray-500">!#$%^'"`&*-=_+&gt;&lt;?;:()&#123;&#125;[].,@</code>
                    <br/><br/>
                    Your password must also have a length of 8 characters or more.
                </p>
            </div>
        
            <form method="POST" class="inner space-y-5" use:enhance={signup}>
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
                <PasswordInput tabIndex="10" class="flex-grow" label="Confirm Password" name="passwordConfirm" description="Confirm your secure password." invalidMsg="Enter a valid secure password.">
                    <svelte:fragment slot="left">
                        <Key size="30"strokeWidth="1"/>
                    </svelte:fragment>
                </PasswordInput>

                {#if submitError !== undefined}
                    <div class="text-red-500 text-sm">• {submitError}</div>
                {/if}

                <div class="flex flex-row space-x-5">
                    <Button tabIndex="10" class="flex-grow" variant="accent" type="submit">Sign up</Button>  
                    <Button tabIndex="10" class="flex-grow" href="/auth/signin">Sign in</Button>          
                </div>
            </form>
        {:else if view === "emailChallenge"}
            <div class="inner">
                <p>
                    A code has been sent to your email address so we can verify you and finish creating your account.
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

    h1 {
        @apply text-xl font-bold h-fit px-5 py-2.5 bg-gray-300/80 border-b border-black text-gray-800/80 rounded-t-[15px];
    }

    p {
        @apply p-5 m-0 font-sans w-full md:max-w-[300px] text-sm;
    }
</style>