<script>
    import TextInput from "$lib/components/inputs/TextInput.svelte";

    export let value = ""
    export let valid = null 
    export let name = ""
    export let label = ""
    export let description = ""
    export let required = true
    export let classList = ""
    export let verifiyValidity = true

    let validClass = ""

    const internalOnInput = () => {
        if (verifiyValidity) {
            const regex = /^(([^<>()\[\]\\.,;:\s@"]+(\.[^<>()\[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/
            valid = regex.test(value)

            if (value === "") {
                validClass = ""
            } else if (valid === true) {
                validClass = "valid"
            } else if (valid === false) {
                validClass = "invalid"
            }
        }
    }
</script>

<svelte:options accessors={true}/>

<TextInput bind:value={value} onInput={internalOnInput} {name} {label} {description} {required} classList={`${validClass} ${classList}`} autocomplete="email" type="email"/>