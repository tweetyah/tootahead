<script lang="ts">
  import { elasticInOut, sineIn } from "svelte/easing";
  import { fly } from "svelte/transition";
  import Sidebar from "./Sidebar.svelte";


  let isSidebarShown = false
</script>

<div id="navbar" class="shadow-sm flex justify-left h-100 w-100 m-2 p-2 rounded bg-dark1 text-slate-100 sm:collapse sm:w-0 sm:p-0 sm:m-0">
  <button on:click={() => isSidebarShown = true}><i class="bx bx-menu mr-4" style="font-size: 24px" /></button>
  <!-- TootAhead -->
  <div class="flex-1 flex text-lg">
    <img src="/img/logo.png" alt="logo" class="w-[28px] h-auto mr-2"/>
    <span>TootAhead!</span>
  </div>
  {#if isSidebarShown}
    <div
      class="w-full sm:w-0 h-full flex sm:h-0 absolute top-0 left-0 right-0 bottom-0 bg-dark1"
      style="z-index: 10000"
      transition:fly={{ x:-1000, duration: 300, easing: sineIn }} >
      <Sidebar closeButton onClose={() => isSidebarShown = false} onNavigate={() => isSidebarShown = false} />
    </div>
  {/if}
</div>