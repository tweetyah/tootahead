<script lang="ts">
  import { sineIn } from "svelte/easing";
  import { fade } from "svelte/transition";
  import type { Post } from "../models";
  import PostEditor from "./PostEditor.svelte";

  export let open: boolean
  export let onClose: Function
  export let onUpdated: Function
  export let posts: Post[]

  function onPostUpdated() {
    onUpdated()
    onClose()
  }
</script>

{#if open}
  <div transition:fade={{ duration: 150, easing: sineIn }}
    class="absolute top-0 bottom-0 left-0 right-0 grid justify-center items-center">
    <!-- svelte-ignore a11y-click-events-have-key-events -->
    <div on:click={() => onClose()} class="absolute top-0 bottom-0 left-0 right-0 bg-gray-900/90" style="z-index: 100;" />
    <div class="relative rounded shadow-sm w-screen sm:w-[80vw] sm:max-w-[800px] h-screen sm:h-auto bg-bglight" style="z-index: 1000;">
      <div class="bg-dark1 text-white p-2 rounded-t-sm flex items-center">
        <span class="flex-1">Edit post</span>
        <button on:click={() => onClose()}><i class="bx bx-x text-4xl md:text-xl" /></button>
      </div>
      <div class="p-2">
        <PostEditor posts={posts} sendAt={posts[0].sendAt} onUpdated={() => onPostUpdated()} />
      </div>
    </div>
  </div>
{/if}