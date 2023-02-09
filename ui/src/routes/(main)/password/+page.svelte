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
    let urls = []

    let submitError

    const submit = async ({ data, cancel }) => {
        submitError = undefined

        const name = data.get("name")
        const username = data.get("username")
        const password = data.get("password")
        const colour = data.get("colour")
        
        if (!username || !password)
            cancel()

        const databaseKey = await storage.getDatabaseKey()
        const authToken = await storage.getAuthToken()

        const encryptedName = await cryptography.encrypt(databaseKey, utils.stringToArrayBuffer(name))
        const encryptedUsername = await cryptography.encrypt(databaseKey, utils.stringToArrayBuffer(username))
        const encryptedPassword = await cryptography.encrypt(databaseKey, utils.stringToArrayBuffer(password))

        const processedAdditionalFields = []
        const processedUrls = []

        let aFLastWasField = false
        let skip = false
        let temp = {}

        for (const [ key, value ] of data) {
            if (!aFLastWasField && key === "key") {
                if (value.trim() === "") {
                    skip = true
                } else {
                    const encryptedKey = await cryptography.encrypt(databaseKey, utils.stringToArrayBuffer(value))
                    temp.key = base64.encode(encryptedKey.encrypted)
                    temp.keyIv = base64.encode(encryptedKey.iv)
                }
            } else if (aFLastWasField && key === "value" && !skip) {
                if (value.trim() !== "") {
                    const encryptedValue = await cryptography.encrypt(databaseKey, utils.stringToArrayBuffer(value))
                    temp.value = base64.encode(encryptedValue.encrypted)
                    temp.valueIv = base64.encode(encryptedValue.iv)
                    processedAdditionalFields.push(temp)
                }
                temp = {}
            } else if (skip) {
                skip = false
                temp = {}
            }
            aFLastWasField = !aFLastWasField

            if (key === "url") {
                const encryptedUrl = await cryptography.encrypt(databaseKey, utils.stringToArrayBuffer(value))
                processedUrls.push({
                    url: base64.encode(encryptedUrl.encrypted),
                    urlIv: base64.encode(encryptedUrl.iv)
                })
            }
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
                colour: colour.replaceAll("#", ""),
                additionalFields: processedAdditionalFields,
                urls: processedUrls,
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
                <TextInput classList="flex-grow" label="Name" name="name" description="Enter a name."/>
                <TextInput classList="flex-grow" label="Username" name="username" description="Enter a username."/>
                <PasswordInput verifiyValidity={false} classList="flex-grow" label="Password" name="password" description="Enter a password."/>
            
                <input type="colour" value="0000ff" name="colour"/>

                <div class="flex flex-col">
                    {#each additionalFields as additionalField}
                        <div class="flex flex-row">
                            <TextInput required={false} classList="flex-grow" name="key" label="Key" bind:value={additionalField.Key}/>
                            <TextInput required={false} classList="flex-grow" name="value" label="Value" bind:value={additionalField.Value}/>
                        </div>
                    {/each}
                </div>

                <div class="flex flex-col">
                    {#each urls as url}
                        <div class="flex flex-row">
                            <TextInput required={false} classList="flex-grow" name="url" label="Url" bind:value={url}/>
                        </div>
                    {/each}
                </div>
            
                <RegularButton onClick={() => {
                    additionalFields = [...additionalFields, { Key: "", Value: "" }]
                    urls = [...urls, ""]
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