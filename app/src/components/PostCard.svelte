<script lang="ts">
  import type { Post } from "../models";
  import ComposerModal from "./ComposerModal.svelte";

  export let post: Post;
  export let editable: boolean = false
  export let onUpdated: Function = undefined
  let isModalOpen: boolean = false;

  function openModal() {
    if(editable) {
      isModalOpen = true
    }
  }

  function onPostUpdated() {
    if(onUpdated) {
      onUpdated()
    }
  }
</script>

<!-- svelte-ignore a11y-click-events-have-key-events -->
<div>
  <div on:click={() => openModal()}
    class="p-2 bg-white border rounded shadow-sm flex flex-col {editable ? "hover:cursor-pointer hover:shadow-md" : null}">
    <div class="mb-2">
      { @html post.html() }
    </div>
    <div class="flex">
      <div class="flex-1 flex">
        <div class="bg-gray-200 flex content-center items-center px-2 rounded text-sm">
          <i class="bx bx-calendar mr-1" />
          { post.sendAt ? `${post.sendAt.toLocaleDateString()} ${post.sendAt.toLocaleTimeString(
            [],
            { hour: '2-digit', minute: '2-digit' })}`: null
          }
        </div>
      </div>
    </div>
  </div>
  <ComposerModal
    open={isModalOpen}
    onClose={() => isModalOpen = false}
    onUpdated={() => onPostUpdated()}
    posts={[post]} />
</div>