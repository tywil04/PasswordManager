<script>
    export let data

    console.log(data)
</script>

Test

{#each data.passwords as password}
    {password}
{/each}