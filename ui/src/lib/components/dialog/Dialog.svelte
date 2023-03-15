<script>
    import { enhance } from "$app/forms";

    export let open
    export let title

    let dialog

    if (open) {
        dialog?.showModal()
    }
</script>

{#if open}
    <div class="absolute flex flex-row h-full w-full justify-center z-50">
        <div class="flex flex-col justify-center">
            <dialog class="bg-gray-200 border border-gray-400 flex flex-col shadow-xl rounded-[4px]" bind:this={dialog}>
                <h1 class="font-bold text-xl">
                    {title}
                </h1>
                <slot/>
            </dialog>
        </div>
    </div>
{/if}