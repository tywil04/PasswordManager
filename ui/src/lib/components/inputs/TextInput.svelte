<script>
    import { createEventDispatcher } from 'svelte';

    const dispatch = createEventDispatcher()

    let value = ""
    let id = crypto.randomUUID()
    let name = ""
    let label = ""
    let description = ""
    let required = true
    let autocomplete = ""
    let type = ""
    let classList = ""
    export { value, id, name, label, description, required, autocomplete, type, classList as class } 

    let input

    const setType = (node, type) => {
        node.type = type
        return {
            update(type) {
                node.type = type
            }
        }
    }

    export const setCustomValidity = (value) => input.setCustomValidity(value)
</script>

<div class="outerContainer">
    {#if label}
        <label for={id}>
            {label} 
            {#if required}
                <span class="labelRequired">(required)</span>
            {/if}
        </label>
    {/if}
    
    <div class="inputContainer"> 
        <slot name="left"/>
        <input bind:this={input} bind:value={value} on:input={(e) => dispatch("input", e)} use:setType={type} class={classList} {id} {name} {required} {autocomplete}/>
        <slot name="right"/>
    </div>

    {#if description}
        <span class="description">{description}</span>
    {/if}
</div>

<style lang="postcss">
    div.inputContainer {
        @apply flex flex-row;
    }

    div.inputContainer :global(*:first-child:not(:last-child):not(input)) {
        @apply bg-gray-300/80 border border-black px-2 py-1 rounded-md rounded-r-none w-fit duration-100 text-sm;
    }

    div.inputContainer :global(*:last-child:not(:first-child):not(input)) {
        @apply bg-gray-300/80 border border-black px-2 py-1 rounded-md rounded-l-none w-fit duration-100 text-sm;
    }

    div.inputContainer > input:first-child:not(:last-child) {
        @apply border-r-0 rounded-r-none;
    }

    div.inputContainer > input:last-child:not(:first-child) {
        @apply border-l-0 rounded-l-none;
    }

    div.outerContainer {
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
</style>