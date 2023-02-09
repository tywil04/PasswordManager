<script>
    import { writable } from 'svelte/store';

    import * as storage from "$lib/js/storage.js"
    import * as utils from "$lib/js/utils.js"
    import * as base64 from "base64-arraybuffer"
    import * as cryptography from "$lib/js/cryptography.js"

    export let data

    var names = writable([])

    async function decryptNames(namesStore) {
        const databaseKey = await storage.getDatabaseKey()
        let internalNames = []
        for (let i = 0; i < data.passwords.length; i++) {
            internalNames.push({
                name: utils.arrayBufferToString(await cryptography.decrypt(databaseKey, {
                    encrypted: base64.decode(data.passwords[i].name),
                    iv: base64.decode(data.passwords[i].nameIv),
                })),
                index: i,
            })
        }
        namesStore.set(internalNames)
    }
    decryptNames(names)
</script>

<div class="flex flex-col h-full w-full bg-blue-500 p-5">
    {#each $names as name}
        <div class="bg-white border border-black w-fit h-fit duration-100 p-5 rounded">
            {name.name}
        </div>
    {/each}
</div>