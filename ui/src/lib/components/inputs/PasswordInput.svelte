<script>
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
    let buttonText = "Show"

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
        [type, buttonText] = type === "password" ? ["text", "Hide"]: ["password", "Show"]
    }
</script>

<TextInput bind:this={input} bind:value={value} bind:type={type} on:input={checkValidity} class={classList} autocomplete="password" {id} {name} {label} {description} {required}>
    <svelte:fragment slot="right">
        <button tabindex="-1" type="button" on:click={toggleHidden} class="font-mono">{buttonText}</button>  
    </svelte:fragment>
</TextInput>