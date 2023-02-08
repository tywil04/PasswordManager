<nav>
    <div class="navBlock justify-start">
        <p class="text-blue-700 font-semibold text-xl p-5 w-fit">Password Manager</p>
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
</style>