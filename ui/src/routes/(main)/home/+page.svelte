<script>
    // import { writable } from 'svelte/store';

    // import * as storage from "$lib/js/storage.js"
    // import * as utils from "$lib/js/utils.js"
    // import * as base64 from "base64-arraybuffer"
    // import * as cryptography from "$lib/js/cryptography.js"

    // export let data

    // var names = writable([])

    // async function decryptNames(namesStore) {
    //     const databaseKey = await storage.getDatabaseKey()
    //     let internalNames = []
    //     for (let i = 0; i < data.passwords.length; i++) {
    //         internalNames.push({
    //             name: utils.arrayBufferToString(await cryptography.decrypt(databaseKey, {
    //                 encrypted: base64.decode(data.passwords[i].name),
    //                 iv: base64.decode(data.passwords[i].nameIv),
    //             })),
    //             index: i,
    //         })
    //     }
    //     namesStore.set(internalNames)
    // }
    // decryptNames(names)

    import { Pane, Splitpanes } from "svelte-splitpanes"
    import { Plus, ChevronDown, Key, DocumentText, ArrowLeftOnRectangle, ArrowPath, Eye, EyeSlash } from "svelte-heros-v2"

    import Dialog from "$lib/components/dialog/Dialog.svelte"

    import * as server from "$lib/js/server.js"

    export let data

    let selectedVaultId = "all"
    let selectedType = "password"
    let selectedId = ""
    let selectedState = {}
    let passwordElement
</script>

<!-- <div class="flex flex-col h-full w-full bg-blue-500 p-5">
    {#each $names as name}
        <div class="bg-white border border-black w-fit h-fit duration-100 p-5 rounded">
            {name.name}
        </div>
    {/each}
</div> -->

<Dialog title="Webauthn" open>
    <div class="flex flex-col">
        <label class="pl-[1px] text-sm">Name</label>
        <input type="text" class="bg-gray-100 rounded-[4px] !border-gray-300 py-2 px-3 !outline-none !ring-0"/>
    </div>
</Dialog>

<Splitpanes>
	<Pane class="!bg-gray-100 flex flex-col" minSize={20} size={20}>
        <div class="bg-gray-200 border-b border-gray-300 py-2.5 px-2.5 flex flex-row">
            <b class="py-2">Password Manager</b>
        </div> 

        <div class="pb-0.5 pt-2.5 mx-2.5 flex flex-row"> 
            <b>Vaults</b>
            <abbr class="ml-auto" title="New Vault">
                <button>
                    <Plus class="focus:outline-none" strokeWidth="1"/>
                </button>
            </abbr>
        </div>

        <button on:click={() => { selectedVaultId = "all"; selectedType = "password"; selectedId = "" }} class="hover:bg-gray-200 group-open:border-b transition-opacity duration-100 border-gray-300 flex flex-row py-1 px-2.5 cursor-pointer">
            <Key class="w-5 focus:outline-none" strokeWidth="1"/>
            <p class="ml-1 leading-[22px]">All Passwords</p>
        </button>

        <button on:click={() => { selectedVaultId = "all"; selectedType = "note"; selectedId = "" }} class="hover:bg-gray-200 group-open:border-b transition-opacity duration-100 border-gray-300 flex flex-row py-1 px-2.5 cursor-pointer">
            <DocumentText class="w-5 focus:outline-none" strokeWidth="1"/>
            <p class="ml-1 leading-[22px]">All Notes</p>
        </button>

        <div class="my-1.5"/>

        {#if data.vaults}
            {#each data.vaults as vault}
                <details class="flex flex-col open:mb-2 group duration-100">
                    <summary class="hover:bg-gray-200 group-open:border-b transition-opacity duration-100 border-gray-300 flex flex-row py-1 px-2.5 cursor-pointer">
                        <div class="w-5 flex flex-col justify-center">
                            <div style:background={vault.colour} class="h-5 w-5 rounded-full"/>
                        </div>
                        <p class="ml-2">{vault.name}</p>
                        <ChevronDown class="ml-auto group-open:-rotate-90 transition-transform duration-100 focus:outline-none" strokeWidth="1"/>
                    </summary>

                    <button on:click={() => { selectedVaultId = vault.id; selectedType = "password" }} class="hover:bg-gray-200 duration-100 cursor-pointer list-none py-0.5 px-2.5 flex flex-row w-full">
                        <Key class="w-5 focus:outline-none" strokeWidth="1"/>
                        <i class="ml-2 select-none">Passwords</i>
                    </button>
        
                    <button on:click={() => { selectedVaultId = vault.id; selectedType = "note" }} class="hover:bg-gray-200 duration-100 cursor-pointer list-none py-0.5 px-2.5 flex flex-row w-full">
                        <DocumentText class="w-5 focus:outline-none" strokeWidth="1"/>
                        <i class="ml-2 select-none">Notes</i>
                    </button>
                </details>
            {/each}
        {/if}
    </Pane>
	<Pane class="!bg-white !overflow-auto max-h-full flex flex-col" minSize={20} size={25}>
        <div class="bg-gray-200 border-b border-gray-300 py-2.5 px-2.5 flex flex-row">
            <abbr class="w-full !no-underline" title={selectedType === "password" ? "Create New Password": "Create New Note"}>
                <button class="bg-gray-100 w-full flex flex-row justify-center py-2 rounded-[4px] hover:bg-gray-300 duration-100" on:click={() => data.passwords = [...data.passwords, {colour:"#000042", name:"blue"}]}>
                    <Plus strokeWidth="1"/>
                    <p class="ml-2 leading-[22.5px]">{selectedType === "password" ? "New Password": "New Note"}</p>
                </button>
            </abbr>
        </div>
        <div class="flex flex-col !overflow-auto max-h-full">
            {#if data.passwords}
                {#key selectedVaultId}
                    {#each selectedType === "password" ? Object.entries(data.passwords): selectedType === "note" ? Object.entries(data.notes): [] as [ id, data ]}
                        {#if selectedVaultId === "all" || data.vaultId === selectedVaultId}
                            <div on:click={() => { selectedId = id; selectedState = {}; passwordElement.type = "password" }} class="border-b border-gray-200 flex flex-row hover:bg-gray-200/60 duration-100 cursor-pointer odd:bg-gray-50">
                                <div style:background={data.colour} class="w-10 h-10 rounded-full my-1.5 mx-2.5"/>
                                <div class="flex flex-col justify-center">
                                    {data.name}
                                </div>
                            </div>
                        {/if}
                    {/each}
                {/key}
            {/if}
        </div>
    </Pane>
	<Pane class="!bg-gray-100" minSize={20} size={65}>
        <div class="bg-gray-200 border-b border-gray-300 py-2.5 px-2.5 flex flex-row">
            <div class="flex flex-row mr-auto">
                <abbr class="!no-underline" title="Synchronise Client Database">
                    <button on:click={server.syncClientData} class="bg-gray-100 w-fit flex flex-row justify-center py-2 px-3 rounded-[4px] hover:bg-gray-300 duration-100">
                        <ArrowPath strokeWidth="1"/>
                        <p class="ml-1 leading-[22.5px]">Sync</p>
                    </button>
                </abbr>
            </div>

            <div class="flex flex-row ml-auto">
                <button class="bg-gray-100 w-fit flex flex-row justify-center py-2 px-3 rounded-[4px] hover:bg-gray-300 duration-100">
                    <ArrowLeftOnRectangle strokeWidth="1"/>
                    <p class="ml-1 leading-[22.5px]">Signout</p>
                </button>
            </div>
        </div> 

        {#if !selectedId}
            <div class="w-full h-full flex flex-col justify-center">
                <div class="w-full h-fit flex flex-row justify-center">
                    Not found.
                </div>
            </div>
        {:else}
            <div class="py-2 px-3 flex flex-col space-y-2">
                {#if selectedType === "password"}
                    <div class="flex flex-col">
                        <label for="name" class="pl-[1px] text-sm">Name</label>
                        <input id="name" type="text" readonly bind:value={data.passwords[selectedId].name} class="bg-gray-50 rounded-[4px] !border-gray-200 py-2 px-3 !outline-none !ring-0"/>
                    </div>

                    <div class="flex flex-col">
                        <label for="username" class="pl-[1px] text-sm">Username</label>
                        <input id="username" type="text" readonly bind:value={data.passwords[selectedId].username} class="bg-gray-50 rounded-[4px] !border-gray-200 py-2 px-3 !outline-none !ring-0"/>
                    </div>

                    <div class="flex flex-col">
                        <label for="password" class="pl-[1px] text-sm">Password</label>
                        <div class="flex flex-row">
                            <input id="password" type="password" readonly bind:this={passwordElement} bind:value={data.passwords[selectedId].password} class="bg-gray-50 rounded-[4px] rounded-r-none !border-gray-200 !border-r-0 py-2 px-3 !outline-none !ring-0 flex-grow"/>
                            <button on:click={() => { selectedState.visible === true ? passwordElement.type = "password": passwordElement.type = "text"; selectedState.visible = !selectedState.visible }} class="bg-gray-50 rounded-[4px] rounded-l-none border border-gray-200 hover:bg-gray-200 py-2 px-3 duration-100">
                                {#key selectedState.visible}
                                    {#if !selectedState.visible}
                                        <Eye class="focus:outline-none" strokeWidth="1"/>
                                    {:else if selectedState.visible}
                                        <EyeSlash class="focus:outline-none" strokeWidth="1"/>
                                    {/if}
                                {/key}
                            </button>
                        </div>
                    </div>

                    <div class="flex flex-col">
                        <div class="flex flex-row">
                            <label for="af" class="pl-[1px] text-sm">Additional Fields</label>
                            <abbr class="ml-auto" title="New Additional Field">
                                <Plus class="w-5" strokeWidth="1"/>
                            </abbr>
                        </div>
                        {#if data.passwords[selectedId].additionalFields.length > 0}
                            {#each data.passwords[selectedId].additionalFields as additionalField, index}
                                <div class="flex flex-row">
                                    <input id={"af-" + index + "-key"} type="text" readonly bind:value={additionalField.key} class="bg-gray-50 rounded-[4px] !border-gray-200 py-2 px-3 !outline-none !ring-0 flex-grow"/>
                                    <input id={"af-" + index + "-value"} type="text" readonly bind:value={additionalField.value} class="bg-gray-50 rounded-[4px] !border-gray-200 py-2 px-3 !outline-none !ring-0 flex-grow"/>
                                </div>    
                            {/each}
                        {:else}
                            <input id="af" type="af" readonly value="No additional fields." class="bg-gray-50 rounded-[4px] border !border-gray-200 py-2 px-3 !outline-none !ring-0"/>
                        {/if}
                    </div>

                    <div class="flex flex-col">
                        <div class="flex flex-row">
                            <label for="url" class="pl-[1px] text-sm">URLs</label>
                            <abbr class="ml-auto" title="New URL">
                                <Plus class="w-5" strokeWidth="1"/>
                            </abbr>
                        </div>
                        {#if data.passwords[selectedId].urls.length > 0}
                            {#each data.passwords[selectedId].urls as url, index}
                                <input id={"url-" + index} type="text" readonly bind:value={url.url} class="bg-gray-50 rounded-[4px] !border-gray-200 py-2 px-3 !outline-none !ring-0"/>
                            {/each}
                        {:else}
                            <input id="url" type="url" readonly value="No urls." class="bg-gray-50 rounded-[4px] !border-gray-200 py-2 px-3 !outline-none !ring-0"/>
                        {/if}
                    </div>
                {/if}

                {selectedId}
                {selectedVaultId}
                {selectedType}
                {selectedState.visible}
            </div>
        {/if}
    </Pane>
</Splitpanes>

<style>
    :global(.splitpanes__pane) {
        scrollbar-width: thin;
        @apply border-0;
    }

    :global(.splitpanes__splitter) {
        @apply !relative !border-0 !w-[1px] !bg-gray-300;
    }

    :global(.splitpanes__splitter::before), :global(.splitpanes__splitter::after) {
        @apply !invisible
    }
</style>