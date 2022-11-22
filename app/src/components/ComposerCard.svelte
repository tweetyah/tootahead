<script lang="ts">
  import type { Tweet } from "../models";
  import { name, handle, profileImgUrl } from "../store";
  import Button from "./Button.svelte";
  import ComposerTextarea from "./ComposerTextarea.svelte";

  // Props
  export let tweet: Tweet;
  export let index: number;
  export let total: number;

  // Functions
  function selectImage() {

  }
</script>

<div class="flex">
  <div>
    <img src={ $profileImgUrl } class="w-[50px] h-[50px] rounded-full m-2" alt="avatar" />
  </div>
  <div class="flex-1 m-2">
    <div>
      <span class="font-bold">{ $name }</span>
      <span class="italic text-slate-600 text-sm">@{ $handle }</span>
    </div>
    <ComposerTextarea bind:value={tweet.text} />
    <div class="flex text-sm align-center text-slate-600">
      <div class="flex-1">
        <span class="mr-2">{ tweet.text.length }/240</span>
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