<script>
    import { Plus, ChevronDown, Key, DocumentText, ArrowLeftOnRectangle, ArrowPath, Eye, EyeSlash, Clipboard, Pencil, Trash, Minus } from "svelte-heros-v2"

    import * as server from "$lib/js/server.js"
    import * as cryptography from "$lib/js/cryptography.js"

    export let data

    let selectedVaultId = "all"
    let selectedType = "password"
    let selectedId = ""
    let selectedState = {}
    let passwordElement

    let editMode = false
    let beforeEditHash = undefined
    let editModeDeletedUrls = []
    let editModeDeletedAdditionalFields = []

    async function toggleEditSelected() {
        let selected
        if (selectedType === "password") {
            selected = data.passwords[selectedId]
        } else if (selectedType === "note") {
            selected = data.notes[selectedId]
        }  

        const recursiveTrimmer = (selected) => {
            for (let [ key, value ] of Object.entries(selected)) {
                if (typeof value === "string") {
                    selected[key] = value.trim()
                } else if (typeof value === "object") {
                    selected[key] = recursiveTrimmer(value)
                }
            }
            return selected
        }
        selected = recursiveTrimmer(selected)

        if (editMode) {
            for (let urlIndex of editModeDeletedUrls) {
                data.passwords[selectedId].urls.splice(urlIndex, urlIndex + 1)
            }

            for (let afIndex of editModeDeletedAdditionalFields) {
                data.passwords[selectedId].additionalFields.splice(afIndex, afIndex + 1)
            }

            if (await cryptography.quickHash(JSON.stringify(selected)) != beforeEditHash) {
                beforeEditHash = undefined

                if (selectedType === "password") {
                    server.updatePassword(selectedId, data.passwords[selectedId])
                } else if (selectedType === "note") {
                    server.updateNote(selectedId, data.notes[selectedId])
                }
            }
            editMode = false
        } else if (!editMode) {
            beforeEditHash = await cryptography.quickHash(JSON.stringify(selected))
            editMode = true
        } 
    }

    function deleteSelected() {
        if (selectedType === "password") {
            server.deletePassword(selectedId, data.passwords[selectedId].vaultId)
            document.getElementById(selectedId).remove()
            selectedId = ""
            selectedState = {}
        } else if (selectedType === "note") {
            server.deleteNote(selectedId, data.notes[selectedId].vaultId)
            document.getElementById(selectedId).remove()
            selectedId = ""
            selectedState = {}
        }
    }

    function deleteUrlFromSelected(index) {
        editModeDeletedUrls.push(index)
        document.getElementById(`url-${index}`).remove()
        document.getElementById(`url-${index}-button`).remove()
    }

    function deleteAdditionalFieldFromSelected(index) {
        editModeDeletedAdditionalFields.push(index)
        document.getElementById(`af-${index}`).remove()
        document.getElementById(`af-${index}-button`).remove()
    }
</script>

<main class="flex flex-row w-100 h-full">
	<div class="!bg-gray-100 flex flex-col w-[20%]">
        <div class="bg-gray-200 border-b border-gray-300 py-2.5 px-2.5 flex flex-row">
            <b class="py-1.5 px-1.5 text-lg">Password Manager</b>
        </div> 

        <button on:click={() => { selectedVaultId = "all"; selectedType = "password", selectedId = ""; editMode = false }} class="hover:bg-gray-200 group-open:border-b mt-3 transition-opacity duration-100 border-gray-300 flex flex-row py-1 px-2.5 cursor-pointer">
            <Key class="w-6 focus:outline-none" strokeWidth="1"/>
            <p class="ml-3">All Passwords</p>
        </button>

        <button on:click={() => { selectedVaultId = "all"; selectedType = "note", selectedId = ""; editMode = false }} class="hover:bg-gray-200 group-open:border-b transition-opacity duration-100 border-gray-300 flex flex-row py-1 px-2.5 cursor-pointer">
            <DocumentText class="w-6 focus:outline-none" strokeWidth="1"/>
            <p class="ml-3">All Notes</p>
        </button>

        <div class="my-2"/>

        {#if data.vaults}
            {#each data.vaults as vault}
                <details class="flex flex-col open:mb-3 group duration-100">
                    <summary class="hover:bg-gray-200 transition-opacity duration-100 border-gray-300 flex flex-row py-2 px-2.5 cursor-pointer">
                        <div class="w-6 flex flex-col justify-center">
                            <div style:background={vault.colour} class="h-6 w-6 rounded-full"/>
                        </div>
                        <b class="ml-3">{vault.name}</b>
                        <ChevronDown class="ml-auto group-open:-rotate-90 transition-transform duration-100 focus:outline-none" strokeWidth="1"/>
                    </summary>
                    
                    <button on:click={() => { selectedVaultId = vault.id; selectedType = "password"; editMode = false }} class="hover:bg-gray-200 duration-100 cursor-pointer list-none py-1 px-2 flex flex-row w-full">
                        <Key class="w-6 focus:outline-none" strokeWidth="1"/>
                        <p class="ml-3 select-none">Passwords</p>
                    </button>
        
                    <button on:click={() => { selectedVaultId = vault.id; selectedType = "note"; editMode = false }} class="hover:bg-gray-200 duration-100 cursor-pointer list-none py-1 px-2 flex flex-row w-full">
                        <DocumentText class="w-6 focus:outline-none" strokeWidth="1"/>
                        <p class="ml-3 select-none">Notes</p>
                    </button>
                </details>
            {/each}
        {/if}
    </div>

    <div class="w-[1px] bg-gray-200"/>

	<div class="!bg-white !overflow-auto max-h-full flex flex-col w-[25%]">
        <div class="bg-gray-200 border-b border-gray-300 py-2.5 px-2.5 flex flex-row">
            <abbr class="w-full !no-underline" title={selectedType === "password" ? "Create New Password": "Create New Note"}>
                <button class="bg-gray-100 w-full flex flex-row justify-center py-2 rounded-[4px] hover:bg-gray-300 duration-100" on:click={() => {data.passwords = [...data.passwords, {colour:"#000042", name:"blue"}]}}>
                    <Plus strokeWidth="1"/>
                    <p class="ml-2 leading-[22.5px]">
                        {selectedType === "password" ? "New Password": "New Note"}
                    </p>
                </button>
            </abbr>
        </div>
        <div class="flex flex-col !overflow-auto max-h-full">
            {#if data.passwords}
                {#key selectedVaultId}
                    {#each selectedType === "password" ? Object.entries(data.passwords): selectedType === "note" ? Object.entries(data.notes): [] as [ id, data ]}
                        {#if selectedVaultId === "all" || data.vaultId === selectedVaultId}
                            <div id={id} on:keydown={() => {}} on:click={() => { selectedId = id; selectedState = {}; passwordElement.type = "password"; editMode = false }} class="border-b border-gray-200 flex flex-row hover:bg-gray-200/60 duration-100 cursor-pointer odd:bg-gray-50">
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
    </div>

    <div class="w-[1px] bg-gray-200"/>

	<div class="!bg-gray-100 w-[65%]">
        <div class="bg-gray-200 border-b border-gray-300 py-2.5 px-2.5 flex flex-row">
            <div class="flex flex-row mr-auto space-x-2">
                <abbr class="!no-underline" title="Synchronise Client Database">
                    <button on:click={server.syncClientData} class="bg-gray-100 w-fit flex flex-row justify-center py-2 px-3 rounded-[4px] hover:bg-blue-600 hover:text-white duration-100">
                        <ArrowPath class="focus:outline-none" strokeWidth="1"/>
                        <p class="ml-1 leading-[22.5px]">Sync</p>
                    </button>
                </abbr>

                {#if selectedId}
                    <abbr class="!no-underline" title="Edit">
                        <button data-editmode={editMode} on:click={toggleEditSelected} class="bg-gray-100 w-fit flex flex-row justify-center py-2 px-3 rounded-[4px] hover:bg-gray-300 duration-100 data-[editmode=true]:bg-red-600 data-[editmode=true]:text-white">
                            <Pencil class="focus:outline-none" strokeWidth="1"/>
                            <p class="ml-1 leading-[22.5px]">{editMode ? "Finish Edit": "Edit"}</p>
                        </button>
                    </abbr>

                    <abbr class="!no-underline" title="Delete">
                        <button data-editmode={editMode} on:click={deleteSelected} class="bg-gray-100 w-fit flex flex-row justify-center py-2 px-3 rounded-[4px] hover:bg-red-600 hover:text-white duration-100">
                            <Trash class="focus:outline-none" strokeWidth="1"/>
                            <p class="ml-1 leading-[22.5px]">Delete</p>
                        </button>
                    </abbr>
                {/if}
            </div>

            <div class="flex flex-row ml-auto">
                <button on:click={server.signout} class="bg-gray-100 w-fit flex flex-row justify-center py-2 px-3 rounded-[4px] hover:bg-red-600 hover:text-white duration-100">
                    <ArrowLeftOnRectangle class="focus:outline-none" strokeWidth="1"/>
                    <p class="ml-1 leading-[22.5px]">Signout</p>
                </button>
            </div>
        </div> 

        {#if !selectedId}
            <!-- I know this tag is deprecated however I was just lazy -->
            <center class="w-full p-10 py-20 text-gray-600/80">
                No {selectedType} selected
            </center>
        {:else}
            <div class="py-2 px-3 flex flex-col space-y-3">
                {#if selectedType === "password"}
                    <div class="flex flex-col">
                        <label for="name" class="pl-[1px] text-sm">Name</label>
                        <input id="name" type="text" readonly={!editMode} bind:value={data.passwords[selectedId].name} class="bg-gray-50 rounded-[4px] !border-gray-200 py-2 px-3 !outline-none !ring-0"/>
                    </div>

                    <div class="flex flex-col">
                        <label for="username" class="pl-[1px] text-sm">Username</label>
                        <div class="flex flex-row">
                            <input id="username" type="text" readonly={!editMode} bind:value={data.passwords[selectedId].username} class="bg-gray-50 rounded-[4px] rounded-r-none !border-gray-200 py-2 px-3 !outline-none !ring-0 flex-grow"/>
                            <abbr class="!no-underline" title="Copy username to clipboard">
                                <button on:click={() => navigator.clipboard.writeText(data.passwords[selectedId].username) } class="bg-gray-50 rounded-[4px] rounded-l-none border !border-l-0 border-gray-200 hover:bg-gray-200 py-2 px-3 duration-100">
                                    <Clipboard class="focus:outline-none" strokeWidth="1"/>
                                </button>
                            </abbr>
                        </div>
                    </div>

                    <div class="flex flex-col">
                        <label for="password" class="pl-[1px] text-sm">Password</label>
                        <div class="flex flex-row">
                            <input id="password" type="password" readonly={!editMode} bind:this={passwordElement} bind:value={data.passwords[selectedId].password} class="bg-gray-50 rounded-[4px] rounded-r-none !border-gray-200 !border-r-0 py-2 px-3 !outline-none !ring-0 flex-grow"/>
                            <abbr class="!no-underline" title="Copy password to clipboard">
                                <button on:click={() => navigator.clipboard.writeText(data.passwords[selectedId].password) } class="bg-gray-50 border !border-r-0 border-gray-200 hover:bg-gray-200 py-2 px-3 duration-100">
                                    <Clipboard class="focus:outline-none" strokeWidth="1"/>
                                </button>
                            </abbr>
                            <abbr class="!no-underline" title="Toggle visibility">
                                <button on:click={() => { selectedState.visible === true ? passwordElement.type = "password": passwordElement.type = "text"; selectedState.visible = !selectedState.visible }} class="bg-gray-50 rounded-[4px] rounded-l-none border border-gray-200 hover:bg-gray-200 py-2 px-3 duration-100">
                                    {#key selectedState.visible}
                                        {#if !selectedState.visible}
                                            <Eye class="focus:outline-none" strokeWidth="1"/>
                                        {:else if selectedState.visible}
                                            <EyeSlash class="focus:outline-none" strokeWidth="1"/>
                                        {/if}
                                    {/key}
                                </button>
                            </abbr>
                        </div>
                    </div>

                    <div class="flex flex-col">
                        <div class="flex flex-row">
                            <label for="af" class="pl-[1px] text-sm">Additional Fields</label>
                            {#if editMode}
                                <abbr class="ml-auto h-[20px] p-0 m-0" title="New Additional Field">
                                    <button>
                                        <Plus class="h-[20px]  focus:outline-none" strokeWidth="1"/>
                                    </button>
                                </abbr>
                            {/if}
                        </div>
                        <div class="flex flex-row">
                            {#key data.passwords[selectedId].additionalFields}
                                <div class="flex flex-col w-full">
                                    {#if data.passwords[selectedId].additionalFields.length > 0}
                                        {#each data.passwords[selectedId].additionalFields as additionalField, index}
                                            <div class="flex flex-row group" id={"af-" + index}>
                                                <input type="text" readonly={!editMode} bind:value={additionalField.key} class="bg-gray-50 !border-gray-200 py-2 px-3 !outline-none !ring-0 flex-grow group-first:rounded-t-[4px] group-last:rounded-b-[4px] border border-b-0 group-last:!border-b group-first:group-last:!border-b"/>
                                                <p class="py-2 px-1.5">:</p>
                                                <input type="text" readonly={!editMode} bind:value={additionalField.value} class="bg-gray-50 !border-gray-200 py-2 px-3 !outline-none !ring-0 flex-grow group-first:rounded-t-[4px] group-last:rounded-b-[4px] border border-b-0 group-last:!border-b group-first:group-last:!border-b"/>
                                            </div>    
                                        {/each}
                                    {:else}
                                        <input id="af" type="af" readonly={!editMode} value="No additional fields." class="bg-gray-50 rounded-[4px] border !border-gray-200 py-2 px-3 !outline-none !ring-0"/>
                                    {/if}
                                </div>

                                {#if editMode}
                                    <div class="flex flex-col w-fit ml-3">
                                        {#each data.passwords[selectedId].additionalFields as _, index}
                                            <button id={"af-" + index + "-button"} on:click={() => deleteAdditionalFieldFromSelected(index)} class="bg-gray-50 border border-gray-200 hover:bg-red-600 hover:text-white py-2 px-3 duration-100 first:rounded-t-[4px] last:rounded-b-[4px] !border-b-0 last:!border-b first:last:!border-b">
                                                <Trash class="focus:outline-none" strokeWidth="1"/>
                                            </button>
                                        {/each}
                                    </div>
                                {/if}
                            {/key}
                        </div>
                    </div>

                    <div class="flex flex-col">
                        <div class="flex flex-row">
                            <label for="url" class="pl-[1px] text-sm">URLs</label>
                            {#if editMode}
                                <abbr class="ml-auto h-[20px] p-0 m-0" title="New URL">
                                    <button>
                                        <Plus class="h-[20px] focus:outline-none" strokeWidth="1"/>
                                    </button>
                                </abbr>
                            {/if}
                        </div>
                        <div class="flex flex-row">
                            {#key data.passwords[selectedId].urls}
                                <div class="flex flex-col group w-full">
                                    {#if data.passwords[selectedId].urls.length > 0}
                                        {#each data.passwords[selectedId].urls as url, index}
                                            <input data-editmode={editMode} id={"url-" + index} type="text" readonly={!editMode} bind:value={url.url} class="bg-gray-50 !border-gray-200 py-2 px-3 !outline-none !ring-0 first:rounded-t-[4px] last:rounded-b-[4px] border border-b-0 last:!border-b first:last:!border-b"/>
                                        {/each}
                                    {:else}
                                        <input id="url" type="url" readonly={!editMode} value="No urls." class="bg-gray-50 rounded-[4px] !border-gray-200 py-2 px-3 !outline-none !ring-0"/>
                                    {/if}
                                </div>
        
                                {#if editMode}
                                    <div class="flex flex-col group w-fit ml-3">
                                        {#if data.passwords[selectedId].urls.length > 0}
                                            {#each data.passwords[selectedId].urls as _, index}
                                                <button id={"url-" + index + "-button"} on:click={() => deleteUrlFromSelected(index)} class="bg-gray-50 border border-gray-200 hover:bg-red-600 hover:text-white py-2 px-3 duration-100 first:rounded-t-[4px] last:rounded-b-[4px] !border-b-0 last:!border-b first:last:!border-b">
                                                    <Trash class="focus:outline-none" strokeWidth="1"/>
                                                </button>
                                            {/each}
                                        {:else}
                                            <input id="url" type="url" readonly={!editMode} value="No urls." class="bg-gray-50 rounded-[4px] !border-gray-200 py-2 px-3 !outline-none !ring-0"/>
                                        {/if}
                                    </div>
                                {/if}
                            {/key}
                        </div>
                    </div>
                {/if}
            </div>
        {/if}
    </div>
</main>