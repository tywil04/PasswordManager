<script>
    import { page } from "$app/stores"

    console.log($page)

    import { ArrowLongRight } from "svelte-heros-v2"
</script>

<div class="flex flex-col min-w-[225px] w-fit border-r border-black h-full">
    <div class:active={$page.route.id === "/settings/account"} class="subHeaderBlockContainer">
        <p class="headerBlock"><b>Account</b></p>
        <a class="subHeaderBlock flex flex-row" href="/settings/account#updateemail">
            <ArrowLongRight size="29" strokeWidth="1" class="mr-2"/> Update Email
        </a>
        <a class="subHeaderBlock flex flex-row" href="/settings/account#updatepassword">
            <ArrowLongRight size="29" strokeWidth="1" class="mr-2"/> Update Password
        </a>
        <a class="subHeaderBlock flex flex-row" href="/settings/account#deleteaccount">
            <ArrowLongRight size="29" strokeWidth="1" class="mr-2"/> Delete Account
        </a>
    </div>

    <div class:active={$page.route.id === "/settings/2fa"} class="subHeaderBlockContainer">
        <p class="headerBlock"><b>2FA</b></p>
        <a class="subHeaderBlock flex flex-row" href="/settings/2fa#webauthn">
            <ArrowLongRight size="29" strokeWidth="1" class="mr-2"/> WebAuthn
        </a>
        <a class="subHeaderBlock flex flex-row" href="/settings/2fa#totp">
            <ArrowLongRight size="29" strokeWidth="1" class="mr-2"/> TOTP
        </a>
        <a class="subHeaderBlock flex flex-row" href="/settings/2fa#email">
            <ArrowLongRight size="29" strokeWidth="1" class="mr-2"/> Email
        </a>
    </div>
</div>

<slot/>

<style lang="postcss">
    .headerBlock {
        @apply h-fit text-black px-3 pb-3 text-sm leading-loose;
    }

    .subHeaderBlock {
        @apply h-fit bg-white text-black px-6 py-1 text-sm leading-loose;
    }

    .subHeaderBlock:hover:not(:active) {
        @apply brightness-95;
    }

    .subHeaderBlockContainer {
        @apply py-3 border-b border-black;
    }

    .subHeaderBlockContainer.active, .subHeaderBlockContainer.active > .subHeaderBlock  {
        @apply bg-gray-100;
    }

    .subHeaderBlockContainer.active:first-child {
        @apply border-t-0;
    }
</style>