<script lang="ts">
  import { navigateTo } from 'svelte-router-spa'

  // Props
  export let to: string
  export let icon: string | undefined
  export let title: string
  export let onNavigate: Function = undefined

  function navigate(to: string) {
    navigateTo(to)
    if(onNavigate) {
      onNavigate()
    }
  }
</script>

<button on:click={() => navigate(to)}
  class="rounded hover:bg-dark3 p-3 text-lg text-left w-full">
  {#if icon}
    <i class="bx {icon}"></i>
  {/if}{ title }
</button>