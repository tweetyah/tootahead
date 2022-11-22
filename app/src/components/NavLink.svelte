<script lang="ts">
  import { navigateTo } from 'svelte-router-spa'

  // Props
  export let to: string
  export let icon: string | undefined
  export let title: string
</script>

<button on:click={() => navigateTo(to)} class="rounded hover:bg-slate-700 p-3 text-lg text-left hover:text-slate-200">
  {#if icon}
    <i class="bx {icon}"></i>
  {/if}{ title }
</button>