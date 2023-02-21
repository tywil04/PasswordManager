<script>
    import { enhance } from "$app/forms"
    import * as server from "$lib/js/server.js"

    const signoutForm = async ({ cancel }) => {
        await server.signout()
        cancel()
    }
</script>

<nav>
    <div class="navBlock justify-start">
        <p class="text-blue-700 font-semibold text-xl p-5 w-fit">Password Manager</p>
    </div>

    <div class="navBlock justify-center">
        <a href="/" class="navLink" class:active={window.location == "/"}>Passwords</a>
    </div>

    <div class="navBlock justify-end">
        <form method="POST" use:enhance={signoutForm}>
            <button type="submit" class="navLink">Sign out</button>
        </form>
    </div>
</nav>

<slot/>

<style lang="postcss">
    nav {
        @apply w-full h-fit bg-white border-b border-black flex flex-row;
    }

    .navBlock {
        @apply w-1/3 flex flex-row h-fit;
    }

    .navLink {
        @apply h-fit bg-gray-200 border-l border-black text-black p-5 text-sm leading-loose duration-100;
    }

    .navLink:hover:not(:active) {
        @apply brightness-90;
    }

    .navLink.active {
        @apply bg-gray-300;
    }

    .navLink:last-of-type {
        @apply border-r;
    }
</style>