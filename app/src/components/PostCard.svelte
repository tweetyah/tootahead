<script lang="ts">
  import type { Post } from "../models";
  import ComposerModal from "./ComposerModal.svelte";

  export let post: Post;
  let isModalOpen: boolean = false;
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<div>
  <div on:click={() => isModalOpen = true} class="my-2 p-2 bg-white border-border border rounded-sm shadow-sm flex flex-col hover:bg-gray-100 hover:cursor-pointer">
    <div class="mb-2">
      { @html post.html() }
    </div>
    <div class="flex">
      <div class="flex-1 flex">
        <!-- { post.sendAt } -->
        <div class="bg-gray-200 flex content-center items-center px-2 rounded text-sm">
          <i class="bx bx-calendar mr-1" />
          { post.sendAt ? `${post.sendAt.toLocaleDateString()} ${post.sendAt.toLocaleTimeString(
            [],
            { hour: '2-digit', minute: '2-digit' })}`: null
          }
        </div>
      </div>
      <!-- <i class="bx bx-edit text-2xl" /> -->
    </div>
  </div>
  <ComposerModal open={isModalOpen} onClose={() => isModalOpen = false} posts={[post]} />
</div>