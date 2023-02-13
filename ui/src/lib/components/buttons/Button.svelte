<script>
    import { createEventDispatcher } from 'svelte';

    const dispatch = createEventDispatcher()

    let type = "button"
    let href = ""
    let classList = ""
    let variant = "default"
    let tabIndex = 0
    export { type, href, classList as class, variant, tabIndex }

    const setType = (node, type) => {
        node.type = type
        return {
            update(type) {
                node.type = type
            }
        }
    }

    let v = {}
    switch (variant) {
        case "accent": v.accent = true; break
        default: v.default = true; break
    }
</script>

{#if href === ""}
    <button on:click={(e) => dispatch("click", e)} use:setType={type} class:default={v.default} class:accent={v.accent} tabindex={tabIndex} class={classList}><slot/></button>
{:else}
    <a use:setType={type} class:default={v.default} class:accent={v.accent} tabindex={tabIndex} class={classList} href={href}><slot/></a>
{/if}

<style lang="postcss">
    .accent {
        @apply w-fit text-sm rounded-md px-2 py-1 text-center duration-100 relative border border-black shadow-[2px_2px_0px_0px_black] cursor-pointer bg-blue-500 text-white;
    }

    .accent:hover, .accent:focus {
        @apply bg-blue-600 duration-[0ms] outline-none ring-0;
    }

    .accent:active {
        @apply shadow-none top-[2px] left-[2px];
    }

    .default {
        @apply w-fit text-sm rounded-md px-2 py-1 text-center duration-100 relative border border-black cursor-pointer bg-gray-200 text-black shadow-none outline-none ring-0;
    }

    .default:hover:not(:active), .default:focus:not(:active) {
        @apply brightness-95 outline-none ring-0;
    }
</style>