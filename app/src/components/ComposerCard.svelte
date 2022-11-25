<script lang="ts">
  import { text } from "svelte/internal";
import type { Post } from "../models";
  import { name, handle, profileImgUrl } from "../store";
  import Button from "./Button.svelte";
  import ComposerTextarea from "./ComposerTextarea.svelte";

  // Props
  export let post: Post;
  export let index: number;
  export let total: number;
  export let onUpdate: Function = undefined

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
      <span class="font-bold">{ $name }</span>
      <span class="italic text-slate-600 text-sm">@{ $handle }</span>
    </div>
    <div class="w-full">
      <ComposerTextarea bind:value={post.text} onUpdate={onUpdate} />
    </div>
    <div class="flex text-sm align-center text-slate-600">
      <div class="flex-1">
        <span class="mr-2">{ post.text ? post.text.length : 0 }/240</span>
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