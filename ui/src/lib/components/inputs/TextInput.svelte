<script>
    export let value = ""
    export let name = ""
    export let label = ""
    export let description = ""
    export let required = true
    export let autocomplete = ""
    export let type = ""
    export let classList = ""
    export let onInput = () => {}
    export let changeType = () => {}

    const id = crypto.randomUUID()

    const typeAction = (node) => {
        node.type = type

        changeType = (newType) => {
            node.type = newType
        }
    }

    const internalOnInput = (event) => {
        onInput(event)
    }
</script>

<svelte:options accessors={true}/>

<div class="column">
    {#if label}
        <label for={id}>
            {label} 
            {#if required}
                <span class="labelRequired">(required)</span>
            {/if}
        </label>
    {/if}
    
    <div class="row"> 
        <input bind:value={value} on:input={internalOnInput} class={classList} {id} {name} {required} {autocomplete} use:typeAction/>

        <slot/>
    </div>

    {#if description}
        <span class="description">{description}</span>
    {/if}
</div>

<style lang="postcss">
    div.row {
        @apply flex flex-row;
    }

    div.column {
        @apply flex flex-col h-fit;
    }

    label {
        @apply mb-[1px] text-sm;
    }

    span.description {
        @apply text-xs mt-[1px] text-gray-600/80;
    }

    span.labelRequired {
        @apply text-xs text-gray-600;
    }

    input {
        @apply text-sm bg-gray-200 border border-black px-2 py-1 outline-none m-0 rounded-md w-fit duration-100;
    }

    input.valid {
        @apply bg-green-300/80;
    }

    input.invalid {
        @apply bg-red-300/80;
    }
</style>