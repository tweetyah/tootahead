<script lang="ts">
  import { sineIn } from "svelte/easing";
  import { fade } from "svelte/transition";

  export let open: boolean
  export let onClose: Function
  export let title: string
</script>

{#if open}
  <div transition:fade={{ duration: 150, easing: sineIn }}
    class="absolute top-0 bottom-0 left-0 right-0 grid justify-center items-center">
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <div on:click={() => onClose()} class="absolute top-0 bottom-0 left-0 right-0 bg-gray-900/90" style="z-index: 100;" />
    <div class={`relative rounded shadow-sm bg-bglight ${$$props.class}`} style="z-index: 1000;">
      <div class="bg-dark1 text-white p-2 rounded-t-sm flex items-center">
        <span class="flex-1"> { title }</span>
        <button on:click={() => onClose()}><i class="bx bx-x text-4xl md:text-xl" /></button>
      </div>
      <div class="p-2">
        <slot />
      </div>
    </div>
  </div>
{/if}