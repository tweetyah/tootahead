<script lang="ts">
  import { onMount } from "svelte/internal";
  import type { Post } from "../models";
  import { name, handle, profileImgUrl, api } from "../store";
  import ComposerTextarea from "./ComposerTextarea.svelte";
  import { instance } from "../store"
  import UsernameLabel from "./UsernameLabel.svelte";
  import Button from "./Button.svelte";
  import Loading from "./Loading.svelte";
  import Modal from "./Modal.svelte";

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
    fileInput.click()
  }

  let fileInput;
	let files;
  $: if (files) {
    uploadMedia(files[0])
  }

  let images = [];

  async function uploadMedia(file) {
    images = [...images, { isLoading: true }]
    let encoded = await getBase64(file)
    let res = await $api.uploadMedia(encoded)
    images = [...images.filter(i => i.isLoading !== true), res]
    console.log('images', images)
  }

  async function getBase64(file) {
    return new Promise((resolve, reject) => {
      let reader = new FileReader();
      reader.readAsDataURL(file);
      reader.onload = function () {
        resolve(reader.result);
      };
      reader.onerror = function (error) {
        reject(error);
      };
    })
  }

  let isLightBoxOpen = false
  let lightBoxImageUrl;
  function openLightbox(url) {
    lightBoxImageUrl = url
    isLightBoxOpen = true
  }

  function closeLightbox() {
    isLightBoxOpen = false
    lightBoxImageUrl = undefined
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
    <div class="flex text-sm items-center align-center text-slate-600">
      <div>
        <span class="mr-2">{ post.text ? post.text.length : 0 }/{maxChars}</span>
        {#if total > 1}
          <span class="mr-2">#{ index + 1 }/{total}</span>
        {/if}
      </div>
      <div class="flex items-center gap-2">
        <Button onClick={() => selectImage()} icon="bx-image-add" class="rounded-full flex items-center justify-center" />
        <input class="invisible w-0 h-0" type="file" bind:this={fileInput} bind:files>
      </div>
    </div>

    {#if images}
      <div class="flex gap-2 mt-2">
        {#each images as image}
          <!-- svelte-ignore a11y-click-events-have-key-events -->
          <div on:click={() => openLightbox(image.preview_url)}
            class="bg-gray-50/80 h-[100px] w-[100px] rounded hover:cursor-pointer hover:border hover:border-mastodon hover:shadow-lg flex items-center justify-center p-1">
            {#if image.isLoading}
              <Loading />
            {:else}
                <img
                src={image.preview_url}
                alt={image.id}
              class="" />
            {/if}
          </div>
        {/each}
      </div>
    {/if}
  </div>

  <Modal title="Image preview" onClose={closeLightbox} open={isLightBoxOpen}>
    <div class="flex items-center justify-center">
      <img src={lightBoxImageUrl} alt="lightbox preview" />
    </div>
  </Modal>
</div>