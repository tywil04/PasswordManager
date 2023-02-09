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
    let hidden = true
    let changeType

    const internalOnInput = () => {
        if (verifiyValidity) {
            const regex = /^((?=.*([A-Z]){2,})(?=.*([!#$%^'"`&*-=_+><?;:(){}\[\].,@]){2,})(?=.*([0-9]){2,})(?=.*([a-z]){2,})).{12,}$/
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

    const toggleType = () => {
        if (hidden)
            changeType("text")
        else
            changeType("password")
        hidden = !hidden
    }
</script>

<svelte:options accessors={true}/>

<TextInput bind:value={value} bind:changeType={changeType} onInput={internalOnInput} {name} {label} {description} {required} classList={`${validClass} ${classList}`} autocomplete="password" type="password">
    <button tabindex="-1" type="button" on:click={toggleType}>{hidden ? "Show": "Hide"}</button>
</TextInput>

<style lang="postcss">
    button {
        @apply text-sm bg-black text-white font-mono border border-black h-fit px-2 py-1 rounded-md ml-2;
    }
</style>