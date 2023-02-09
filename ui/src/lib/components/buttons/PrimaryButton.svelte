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
        @apply w-fit text-sm rounded-md px-2 py-1 text-center duration-100 relative border border-black shadow-[2px_2px_0px_0px_black] cursor-pointer bg-blue-500 text-white;
    }

    button:hover, button:focus {
        @apply bg-blue-600 duration-[0ms];
    }

    button:active {
        @apply shadow-none top-[2px] left-[2px];
    }
</style>