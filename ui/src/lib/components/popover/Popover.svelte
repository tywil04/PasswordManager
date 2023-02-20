<script>
    import { fade } from "svelte/transition"
    import { writable } from "svelte/store";

    let open = writable(false) 
    let anchor
    export { open, anchor }

    let bottom = writable(0)
    let left = writable(0)

    const anchorPosition = () => {
        const bounding = anchor?.getBoundingClientRect() ?? { bottom: 0, left: 0 }
        bottom.set(bounding.bottom)
        left.set(bounding.left)
    }

    open.subscribe(anchorPosition)
</script>

<svelte:window on:resize={anchorPosition} />
  
{#if $open}
    <div class="popover">
        <div style="top: {$bottom}px; left: {$left}px;" class="wrapper">
            <slot/>
        </div>
    </div>
{/if}
  
<style>
    .popover {
        @apply fixed;
    }

    .wrapper {
        @apply fixed w-fit h-fit;
    }
</style>