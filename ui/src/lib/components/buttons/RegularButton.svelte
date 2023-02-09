<script>
    import { goto } from '$app/navigation';
    
    let submit = false
    let href = ""
    let classList = ""
    let onClick = () => {}
    export { submit, href, classList as class, onClick }

    if (href !== "") {
        onClick = () => goto(href)
    }
</script>

<button class={classList} on:click={onClick} type={submit ? "submit": "button"}><slot/></button>

<style lang="postcss">
    button {
        @apply w-fit text-sm rounded-md px-2 py-1 text-center duration-100 relative border border-black cursor-pointer bg-gray-100 text-black shadow-none;
    }

    button:hover, button:focus {
        @apply duration-[0ms];
    }

    button:hover {
        @apply bg-gray-200/80;
    }
</style>