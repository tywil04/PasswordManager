<script>
    import TextInput from "$lib/components/inputs/TextInput.svelte"

    let value = ""
    let id = undefined 
    let name = ""
    let label = ""
    let description = ""
    let required = true
    let classList = ""
    let invalidMsg = "Invalid email."
    let checkValid = true
    export { value, id, name, label, description, required, classList as class, invalidMsg, checkValid }

    let input

    const checkValidity = () => {
        if (checkValid) {
            const valid = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/.test(value)
            if (value === "" || valid)
                input.setCustomValidity("")
            else if (!valid)
                input.setCustomValidity(invalidMsg)
        }
    }
</script>

<TextInput bind:this={input} bind:value={value} on:input={checkValidity} class={classList} autocomplete="email" type="email" {id} {name} {label} {description} {required}>
    <slot name="left" slot="left"/>
    <slot name="right" slot="right"/>
</TextInput>