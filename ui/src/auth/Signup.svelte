<script>
    import * as validations from "../../lib/validations.js"
    import * as cryptography from "../../lib/cryptography.js"

    import * as base64 from "base64-arraybuffer"

    import PrimaryButton from "../../components/buttons/PrimaryButton.svelte";
    import RegularButton from "../../components/buttons/RegularButton.svelte";
    import EmailInput from "../../components/inputs/EmailInput.svelte";
    import PasswordInput from "../../components/inputs/PasswordInput.svelte";
    import TextInput from "../../components/inputs/TextInput.svelte";

    let urlParams = new URLSearchParams(window.location.hash.replace("#", ""));

    let signupData = {
        email: "",
        password: "",
        passwordConfirm: "",
    }

    let emailChallengeData = {
        code: "",
    }

    let submitError

    async function signup() {
        submitError = undefined

        const { email, password, passwordConfirm } = signupData

        if (!validations.validateEmail(email) || !validations.validatePassword(password) || password !== passwordConfirm) {
            return
        }

        let masterKey = await cryptography.generateMasterKey(password, email)
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

        console.log(json)

        if (response.status !== 200) {
            submitError = json.error.code
        } else if (response.status === 200) {
            window.location.hash = `#/auth/signup?view=emailChallenge&data=${json.emailChallengeId}`
            urlParams = new URLSearchParams(window.location.hash.replace("#", ""))
        }
    }

    async function emailChallenge() {
        submitError = undefined

        const { code } = emailChallengeData

        if (code === "" || code === null || code === undefined) {
            return
        }

        const response = await fetch("/api/v1/email/signupChallenge", {
            method: "POST",
            headers: {
                "Content-type": "application/json",
            },
            body: JSON.stringify({
                "emailChallengeId": urlParams.get("data"),
                "code": code,
            })
        })
        const json = await response.json()

        if (response.status !== 200) {
            submitError = json.error.code
        } else if (response.status === 200) {
            window.location = "#/auth/signin"
        }
    }
</script>

<main class="w-full h-full bg-blue-400 flex flex-col justify-center">
    <div class="w-full h-fit flex flex-col justify-center space-y-5 p-16 md:p-0 md:space-x-5 md:space-y-0 md:flex-row">
        {#if urlParams.get("view") == "signup"}
            <div class="bg-white border border-black w-full h-fit md:w-fit duration-100 rounded-2xl">
                <h1 class="text-xl font-bold h-fit px-5 py-2.5 bg-black text-white rounded-t-[15px]">Sign up</h1> 
                <p class="p-5 m-0 font-sans w-full md:max-w-[300px] text-sm">
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
                    Your password must also have a length of 12 characters or more.
                </p>
            </div>
        
            <form class="bg-white border border-black w-full h-fit md:w-fit duration-100 p-5 rounded-2xl space-y-5 min-w-[30%]" on:submit|preventDefault={signup}>
                <EmailInput bind:value={signupData.email} class="flex-grow" label="Email" name="email" description="Enter your email address."/>
                <PasswordInput bind:value={signupData.password} class="flex-grow" label="Password" name="password" description="Enter a secure password."/>
                <PasswordInput bind:value={signupData.passwordConfirm} class="flex-grow" label="Confirm Password" name="passwordConfirm" description="Confirm your secure password."/>

                {#if submitError !== undefined}
                    <div class="text-red-500 text-sm">• {submitError}</div>
                {/if}

                <div class="flex flex-row space-x-5">
                    <PrimaryButton class="flex-grow" submit>Sign up</PrimaryButton>  
                    <RegularButton class="flex-grow" href="#/auth/signin">Sign in</RegularButton>   
                </div>
            </form>
        {:else if urlParams.get("view") == "emailChallenge"}
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
                <TextInput bind:value={emailChallengeData.code} class="flex-grow" label="Code" name="code" description="Enter the code from your email."/>

                {#if submitError !== undefined}
                    <div class="text-red-500 text-sm">• {submitError}</div>
                {/if}

                <div class="flex flex-row space-x-5">
                    <PrimaryButton class="flex-grow" submit>Verify</PrimaryButton>         
                </div>
            </form>
        {/if}
    </div>
</main>