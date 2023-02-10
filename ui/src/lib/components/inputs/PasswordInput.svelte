<script>
    import { Eye, EyeSlash } from "svelte-heros-v2"

    import TextInput from "$lib/components/inputs/TextInput.svelte"

    let value = ""
    let id = undefined
    let name = ""
    let label = ""
    let description = ""
    let required = true
    let classList = ""
    let invalidMsg = "Invalid password."
    let checkValid = true
    export { value, id, name, label, description, required, classList as class, invalidMsg, checkValid }

    let input

    let type = "password"
    let icon = Eye

    const checkValidity = () => {
        if (checkValid) {
            const valid = /^((?=.*([A-Z]){1,})(?=.*([a-z]){1,})(?=.*([0-9]){1,})(?=.*([!#$%^'"\`&*-=_+><?;:(){}\[\].,@]){1,})).{8,}$/.test(value)
            if (value === "" || valid)
                input.setCustomValidity("")
            else if (!valid)
                input.setCustomValidity(invalidMsg)
        }
    }

    const toggleHidden = () => {
        [type, icon] = type === "password" ? ["text", EyeSlash]: ["password", Eye]
    }
</script>

<TextInput bind:this={input} bind:value={value} bind:type={type} on:input={checkValidity} class={classList} autocomplete="password" {id} {name} {label} {description} {required}>
    <slot name="left" slot="left"/>
    <svelte:fragment slot="right">
        <button tabindex="-1" type="button" on:click={toggleHidden} class="font-mono">
            <svelte:component this={icon} size="20" strokeWidth="1"/>
        </button>  
        <slot name="right"/>
    </svelte:fragment>
</TextInput>

<style lang="postcss">
    button {
        @apply font-mono duration-100;
    }

    button:hover:not(:active) {
        @apply brightness-90;
    }

    :global(svg) {
        @apply outline-none;
    }
</style>