<script>
    import { enhance } from "$app/forms"

    import * as base64 from "base64-arraybuffer"

    import * as cryptography from "$lib/js/cryptography.js"
    import * as utils from "$lib/js/utils.js"
    import * as storage from "$lib/js/storage.js"
    import * as validations from "$lib/js/validations.js"

    import Button from "$lib/components/buttons/Button.svelte"
    import PasswordInput from "$lib/components/inputs/PasswordInput.svelte"
    import TextInput from "$lib/components/inputs/TextInput.svelte"
    import ColourInput from "$lib/components/inputs/ColourInput.svelte";

    let additionalFields = []
    let urls = []

    let submitError

    const newVault = async ({data, cancel}) => {
        submitError = undefined

        const name = data.get("name")
        const colour = data.get("colour").replaceAll("#", "")

        if (!validations.validateHexColour("#" + colour)) 
            cancel()

        const databaseKey = await storage.getDatabaseKey()
        const authToken = await storage.getAuthToken()

        const encryptedName = await cryptography.encrypt(databaseKey, utils.stringToArrayBuffer(name))
        const encryptedColour = await cryptography.encrypt(databaseKey, utils.stringToArrayBuffer(colour))

        const { status, json } = await utils.postJson("/api/v1/vaults", {
            body: {
                name: base64.encode(encryptedName.encrypted),
                nameIv: base64.encode(encryptedName.iv),
                colour: base64.encode(encryptedColour.encrypted),
                colourIv: base64.encode(encryptedColour.iv),
            },
            headers: {
                "Authorization": authToken,
            }
        })

        console.log(status)
        console.log(json)

        cancel()
    }

    const newPassword = async ({ data, cancel }) => {
        submitError = undefined

        const name = data.get("name")
        const username = data.get("username")
        const password = data.get("password")
        const colour = data.get("colour").replaceAll("#", "")

        const vaultId = data.get("vaultId")  // temp
        
        if (!username || !password || !validations.validateHexColour("#" + colour))
            cancel()

        const databaseKey = await storage.getDatabaseKey()
        const authToken = await storage.getAuthToken()

        const encryptedName = await cryptography.encrypt(databaseKey, utils.stringToArrayBuffer(name))
        const encryptedUsername = await cryptography.encrypt(databaseKey, utils.stringToArrayBuffer(username))
        const encryptedPassword = await cryptography.encrypt(databaseKey, utils.stringToArrayBuffer(password))
        const encryptedColour = await cryptography.encrypt(databaseKey, utils.stringToArrayBuffer(colour))

        const processedAdditionalFields = []
        const processedUrls = []
        let lastKey
        
        for (const [ key, value ] of data) {
            if (lastKey !== undefined && key === "value") {
                const encryptedKey = await cryptography.encrypt(databaseKey, utils.stringToArrayBuffer(lastKey))
                const encryptedValue = await cryptography.encrypt(databaseKey, utils.stringToArrayBuffer(value))
            
                processedAdditionalFields.push({
                    key: base64.encode(encryptedKey.encrypted),
                    keyIv: base64.encode(encryptedKey.iv),
                    value: base64.encode(encryptedValue.encrypted),
                    valueIv: base64.encode(encryptedValue.iv)
                })
            }

            if (key === "key") {
                lastKey = value
            }

            if (key === "url") {
                const encryptedUrl = await cryptography.encrypt(databaseKey, utils.stringToArrayBuffer(value))
                processedUrls.push({
                    url: base64.encode(encryptedUrl.encrypted),
                    urlIv: base64.encode(encryptedUrl.iv)
                })
            }
        }

        const { status, json } = await utils.postJson("/api/v1/vaults/passwords", {  
            body: {
                vaultId: vaultId,
                name: base64.encode(encryptedName.encrypted),
                nameIv: base64.encode(encryptedName.iv),
                username: base64.encode(encryptedUsername.encrypted),
                usernameIv: base64.encode(encryptedUsername.iv),
                password: base64.encode(encryptedPassword.encrypted),
                passwordIv: base64.encode(encryptedPassword.iv),
                colour: base64.encode(encryptedColour.encrypted),
                colourIv: base64.encode(encryptedColour.iv),
                additionalFields: processedAdditionalFields,
                urls: processedUrls,
            },
            headers: {
                "Authorization": authToken,
            }
        })

        if (status !== 200) {
            submitError = json.error.code
        } else if (status === 200) {
            console.log(json)
        }
        
        cancel()
    }
</script>

<main>
    <div class="outer">
        <div class="inner">
            <form method="POST" class="inner space-y-5" use:enhance={newVault}>
                <TextInput class="flex-grow" label="Name" name="name" description="Enter a name."/>
                <ColourInput class="flex-grow" label="Colour" name="colour" description="Find a colour." invalidMsg="Find a valid colour."/>
                <Button type="submit">Submit</Button>
            </form> 

            <form method="POST" class="inner space-y-5" use:enhance={newPassword}>
                <TextInput class="flex-grow" label="Vault Id" name="vaultId" description="Enter a vault id."/>
                <TextInput class="flex-grow" label="Name" name="name" description="Enter a name."/>
                <TextInput class="flex-grow" label="Username" name="username" description="Enter a username."/>
                <PasswordInput checkValid={false} class="flex-grow" label="Password" name="password" description="Enter a password."/>
                <ColourInput class="flex-grow" label="Colour" name="colour" description="Find a colour." invalidMsg="Find a valid colour."/>

                <div class="flex flex-col">
                    {#each additionalFields as additionalField}
                        <div class="flex flex-row">
                            <TextInput required={false} class="flex-grow" name="key" label="Key" bind:value={additionalField.key}/>
                            <TextInput required={false} class="flex-grow" name="value" label="Value" bind:value={additionalField.value}/>
                        </div>
                    {/each}
                </div>

                <div class="flex flex-col">
                    {#each urls as url}
                        <div class="flex flex-row">
                            <TextInput required={false} class="flex-grow" name="url" label="Url" bind:value={url}>
                                <svelte:fragment slot="left">
                                    <span>https://</span>
                                </svelte:fragment>
                            </TextInput>
                        </div>
                    {/each}
                </div>
            
                <Button on:click={() => {
                    additionalFields = [...additionalFields, { key: "", value: "" }]
                }}>additionalField</Button>

                <Button on:click={() => {
                    urls = [...urls, ""]
                }}>url</Button>
            
                <Button type="submit">Submit</Button>
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