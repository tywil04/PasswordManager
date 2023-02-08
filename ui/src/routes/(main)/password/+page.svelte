<script>
    import { enhance } from "$app/forms"

    import * as base64 from "base64-arraybuffer"

    import * as cryptography from "$lib/js/cryptography.js"
    import * as utils from "$lib/js/utils.js"
    import * as storage from "$lib/js/storage.js"

    import RegularButton from "$lib/components/buttons/RegularButton.svelte"
    import PasswordInput from "$lib/components/inputs/PasswordInput.svelte"
    import TextInput from "$lib/components/inputs/TextInput.svelte"

    let additionalFields = []

    let submitError

    const submit = async ({ data, cancel }) => {
        submitError = undefined

        const name = data.get("name")
        const username = data.get("username")
        const password = data.get("password")
        
        if (!username || !password)
            cancel()

        const databaseKey = storage.getDatabaseKey()
        const authToken = storage.getAuthToken()
        const importedDatabaseKey = await cryptography.importDatabaseKey(databaseKey)

        const encryptedName = await cryptography.encrypt(importedDatabaseKey, utils.stringToArrayBuffer(name))
        const encryptedUsername = await cryptography.encrypt(importedDatabaseKey, utils.stringToArrayBuffer(username))
        const encryptedPassword = await cryptography.encrypt(importedDatabaseKey, utils.stringToArrayBuffer(password))

        const processedAdditionalFields = []
        let lastWasField = false
        let skip = false
        let temp = {}

        for (const [ key, value ] of data) {
            if (!lastWasField && key === "key") {
                if (value.trim() === "") {
                    skip = true
                } else {
                    const encryptedKey = await cryptography.encrypt(importedDatabaseKey, utils.stringToArrayBuffer(value))
                    temp.Key = base64.encode(encryptedKey.encrypted)
                    temp.KeyIv = base64.encode(encryptedKey.iv)
                }
            } else if (lastWasField && key === "value" && !skip) {
                if (value.trim() !== "") {
                    const encryptedValue = await cryptography.encrypt(importedDatabaseKey, utils.stringToArrayBuffer(value))
                    temp.Value = base64.encode(encryptedValue.encrypted)
                    temp.ValueIv = base64.encode(encryptedValue.iv)
                    processedAdditionalFields.push(temp)
                }
                temp = {}
            } else if (skip) {
                skip = false
                temp = {}
            }
            lastWasField = !lastWasField
        }

        const response = await fetch("/api/v1/password", {
            method: "POST",
            headers: {
                "Content-type": "application/json",
                "Authorization": authToken,
            },
            body: JSON.stringify({
                name: base64.encode(encryptedName.encrypted),
                nameIv: base64.encode(encryptedName.iv),
                username: base64.encode(encryptedUsername.encrypted),
                usernameIv: base64.encode(encryptedUsername.iv),
                password: base64.encode(encryptedPassword.encrypted),
                passwordIv: base64.encode(encryptedPassword.iv),
                additionalFields: processedAdditionalFields,
            })
        })
        const json = await response.json()

        if (response.status !== 200) {
            submitError = json.error.code
        } else if (response.status === 200) {
            console.log(json)
        }
        
        cancel()
    }
</script>

<main>
    <div class="outer">
        <div class="inner">
            <form method="POST" class="inner space-y-5" use:enhance={submit}>
                <TextInput classList="flex-grow" label="Name" name="name" description="Enter a namae."/>
                <TextInput classList="flex-grow" label="Username" name="username" description="Enter a username."/>
                <PasswordInput verifiyValidity={false} classList="flex-grow" label="Password" name="password" description="Enter a password."/>
            
                <div class="flex flex-col">
                    {#each additionalFields as additionalField}
                        <div class="flex flex-row">
                            <TextInput required={false} classList="flex-grow" name="key" label="Key" bind:value={additionalField.Key}/>
                            <TextInput required={false} classList="flex-grow" name="value" label="Value" bind:value={additionalField.Value}/>
                        </div>
                    {/each}
                </div>
            
                <RegularButton onClick={() => {
                    additionalFields = [...additionalFields, { Key: "", Value: "" }]
                }}>Click me!</RegularButton>
            
                <RegularButton submit>Submit</RegularButton>
            </form>            
        </div>
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
        @apply bg-white border border-black w-full h-fit md:w-fit duration-100 p-5 rounded-2xl;
    }
</style>