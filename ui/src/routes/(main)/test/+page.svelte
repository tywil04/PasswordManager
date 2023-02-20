<script>
    import Popover from "$lib/components/popover/Popover.svelte"

    let popoverAnchor
    let popoverOpen
</script>



<main>
    <button class="w-fit bg-gray-100 p-5 m-5">Click Me!</button>

    <button bind:this={popoverAnchor} on:click={() => popoverOpen.set(!$popoverOpen)} class="w-fit bg-gray-100 p-5 m-5">Click Me!</button>

    <Popover bind:anchor={popoverAnchor} bind:open={popoverOpen}>
        Test
    </Popover>
</main>