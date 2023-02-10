<script>
    import TextInput from "$lib/components/inputs/TextInput.svelte"

    let value = `#${(Math.random() * 0xfffff * 1000000).toString(16).slice(0, 6)}` 
    let id = undefined
    let name = ""
    let label = ""
    let description = ""
    let required = true
    let classList = ""
    let invalidMsg = "Invalid colour."
    let checkValid = true
    export { value, id, name, label, description, required, classList as class, invalidMsg, checkValid }

    let input

    const checkValidity = () => {
        if (checkValid) {
            const valid = /^[#]{1,1}[0123456789abcdef]{6,6}$/.test(value)
            if (value === "" || valid)
                input.setCustomValidity("")
            else if (!valid)
                input.setCustomValidity(invalidMsg)
        }
    }
</script>

<TextInput bind:this={input} bind:value={value} on:input={checkValidity} class={classList} {id} {name} {label} {description} {required}>
    <svelte:fragment slot="left">
        <input type="color" bind:value={value}/>
        <slot name="left"/>
    </svelte:fragment>
    <slot name="right" slot="right"/>
</TextInput>

<style>
    input {
        @apply p-0 m-0 h-auto border border-black rounded-l-md cursor-pointer duration-100;
    }

    input:hover {
        @apply brightness-90;
    }

    input::-moz-color-swatch {
        @apply rounded-l-[5px] border-0 p-0 m-0 w-full;
    }

    input::-moz-color-swatch-wrapper {
        @apply m-0 p-0;
    }

    input::-webkit-color-swatch {
        @apply rounded-l-[5px] border-0 p-0 m-0 w-full;
    }

    input::-webkit-color-swatch-wrapper {
        @apply p-0 m-0;
    }
</style>
