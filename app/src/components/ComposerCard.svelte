<script lang="ts">
  import { onMount } from "svelte/internal";
  import type { Post } from "../models";
  import { name, handle, profileImgUrl } from "../store";
  import ComposerTextarea from "./ComposerTextarea.svelte";
  import { instance } from "../store"
  import UsernameLabel from "./UsernameLabel.svelte";

  // Private fields
  let maxChars = 500

  // Public fields
  export let post: Post;
  export let index: number;
  export let total: number;
  export let onUpdate: Function = undefined

  // Hooks
  onMount(() => {
    if($instance?.configuration?.statuses?.max_characters) {
      maxChars = $instance.configuration.statuses.max_characters
    }
  })

  // Functions
  function selectImage() {

  }
</script>

<div class="flex">
  <div>
    <img src={ $profileImgUrl } class="w-[50px] h-[50px] rounded-full m-2" alt="avatar" />
  </div>
  <div class="flex-1 m-2 min-w-0">
    <div>
      <UsernameLabel name={$name} class="font-bold" />
      <span class="italic text-slate-600 text-sm">@{ $handle }</span>
    </div>
    <div class="w-full">
      <ComposerTextarea bind:value={post.text} onUpdate={onUpdate} />
    </div>
    <div class="flex text-sm align-center text-slate-600">
      <div class="flex-1">
        <span class="mr-2">{ post.text ? post.text.length : 0 }/{maxChars}</span>
        {#if total > 1}
          <span class="mr-2">#{ index + 1 }/{total}</span>
        {/if}
      </div>
      <div>
        <!-- <Button onClick={() => selectImage()} icon="bx-image-add" title="Add image" /> -->
      </div>
    </div>
  </div>
</div>