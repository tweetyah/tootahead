<script lang="ts">
  import { fly } from "svelte/transition";
  import { alert } from "../store"

  let timeout

  function dismiss() {
    alert.set(undefined)
    timeout = undefined
  }

  alert.subscribe((val) => {
    if(val && !timeout) {
      timeout = 1
      timeout = setTimeout(() => {
        dismiss()
        timeout = undefined
      }, 5000)
    }
  })
</script>

<div>
  {#if $alert}
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <div transition:fly={{ x: 200 }} on:click={() => dismiss()}
      class="rounded absolute bottom-0 right-0 w-full sm:w-[300px]" style="z-index: 10000">
      <div class="m-2 bg-slate-300 shadow-lg">
        <div class="bg-slate-900 rounded-t p-2 text-slate-100">
          <i class="bx bx-alarm" /> { $alert.title }
        </div>
        <div class="p-2">
          { $alert.body }
        </div>
      </div>
    </div>
  {/if}
</div>