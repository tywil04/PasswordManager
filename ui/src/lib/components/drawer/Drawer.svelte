<script>
  import { blur } from "svelte/transition"
  import { sineIn } from "svelte/easing"

  import { classNames } from "$lib/components/utils/classNames.js"

  export let visible = false
  export let width = "w-80"
  export let position = "left"

  const transition = {
    x: -320,
    duration: 200,
    easing: sineIn
  }

  function toggleVisibility() {
    visible = !visible
  }

  const classes = {
    container: "fixed overflow-y-auto z-50 p-4 bg-white dark:bg-gray-800",
    backdrop: "fixed top-0 left-0 z-50 w-full h-full bg-gray-900 bg-opacity-75",
    position: {
      left: "inset-y-0 left-0",
      right: "inset-y-0 right-0",
      top: "inset-x-0 top-0",
      bottom: "inset-x-0 bottom-0"
    }[position]
  }
</script>

{#if visible}
  <div role="presentation" class={classes.backdrop} on:click={toggleVisibility}/>
  <div tabindex="-1" class={classNames(width, classes.container, classes.position, $$props.class)} transition:blur={transition} {...$$restProps}>
    <slot/>
  </div>
{/if}