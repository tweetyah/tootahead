<script lang="ts">
  import type { Post } from "../models";
  import { alert, api, service } from "../store";
  import textvars from "../textvars";
  import Accordion from "./Accordion.svelte";
  import AccordionNode from "./AccordionNode.svelte";
  import Button from "./Button.svelte";
  import ComposerCard from "./ComposerCard.svelte";
  import RetweetAtScheduler from "./RetweetAtScheduler.svelte";
  import SendAtScheduler from "./SendAtScheduler.svelte";
  import { fly } from "svelte/transition";
  import { cubicInOut } from "svelte/easing"
  import { addMinutes } from 'date-fns'
  import { onDestroy, onMount } from "svelte";

  let isMobileDrawerOpen = false

  export let sendAt: Date
  $: if(sendAt) calcIsSaveDisabled()
  export let posts: Post[]
  export let onUpdated: Function

  let timeout
  onMount(() => {
    timeout = setTimeout(() => {
      calcIsSaveDisabled()
    }, 30000)
  })

  onDestroy(() => {
    timeout = undefined
  })

  async function update() {
    posts.forEach(p => {
      p.sendAt = sendAt
    })
    await $api.updatePosts(posts)
    reset()
    alert.set({
      title: "Post updated",
      body: "You're post was updated successfully!"
    })
    if (onUpdated) {
      onUpdated()
    }
  }

  async function deletePosts() {
    await $api.deletePosts(posts)
    reset()
    alert.set({
      title: "Post deleted",
      body: "You're post was deleted successfully!"
    })
    if (onUpdated) {
      onUpdated()
    }
  }

  function reset() {
    posts = [{
      text: ""
    }]
    isSaveDisabled = true
    isMobileDrawerDisabled = true
  }

  function calculateValidation() {
    calcIsSaveDisabled()
    calcIsMobileDrawerDisabled()
  }

  let isSaveDisabled = true
  function calcIsSaveDisabled() {
    let isDisabled = false
    posts.forEach(p => {
      if(!p.text) {
        isDisabled = true
      }
    })
    if(sendAt < new Date()) {
      isDisabled = true
    }
    isSaveDisabled = isDisabled
  }

  let isMobileDrawerDisabled = true
  function calcIsMobileDrawerDisabled() {
    let isDisabled = false
    posts.forEach(p => !p.text ? isDisabled = true : null )
    isMobileDrawerDisabled = isDisabled
  }
</script>

<div>
  <div class="grid grid-cols-2 gap-2">
    <div id="composer-wrapper" class="col-span-2 sm:col-span-1">
      <div class="bg-white shadow-sm rounded mb-2">
        {#each posts as p, idx}
          <ComposerCard bind:post={p} index={idx} total={posts.length} onUpdate={() => calculateValidation()} />
        {/each}
      </div>
    </div>
    <div id="composer-settings" class="invisible w-0 sm:visible sm:w-auto">
      <div class="mb-2">
        <Accordion>
          <AccordionNode
            title="Send at"
            subtitle="{sendAt.toLocaleDateString()} {sendAt.toLocaleTimeString(
              [],
              { hour: '2-digit', minute: '2-digit' }
            )}"
          >
            <div class="px-3">
              <SendAtScheduler bind:value={sendAt} />
            </div>
          </AccordionNode>
        </Accordion>
      </div>
      <div class="flex gap-2">
        <Button
          onClick={() => update()}
          icon="bxs-save"
          title="Update"
          disabled={isSaveDisabled}
        />
        <Button
          onClick={() => deletePosts()}
          icon="bxs-trash"
          title="Delete"
        />
      </div>
    </div>
  </div>

  <button
    on:click={() => (isMobileDrawerOpen = true)}
    class="visible sm:invisible shadow-sm hover:shadow-xl hover:bg-mastodon-hover absolute right-5 bottom-5 rounded-full round-full bg-mastodon text-white grid justify-center content-center h-[48px] w-[48px] disabled:bg-mastodon-disabled disabled:cursor-not-allowed"
    style="font-size: 32px;"
    disabled={isMobileDrawerDisabled}
  >
    <i class="bx bx-right-arrow-alt" />
  </button>

  {#if isMobileDrawerOpen}
    <div
      transition:fly={{ y: 500, easing: cubicInOut, duration: 500 }}
      class="absolute top-0 left-0 bottom-0 right-0 bg-slate-100 flex flex-col"
      style="z-index: 1000;"
    >
      <div class="bg-dark1 w-full p-4 shadow-sm text-slate-50 flex text-lg">
        <span class="flex-1">Send post</span>
        <button
          style="font-size: 18px;"
          on:click={() => (isMobileDrawerOpen = false)}
          ><i class="bx bx-x" /></button
        >
      </div>
      <div>
        <Accordion>
          <AccordionNode
            title="Send at"
            subtitle="{sendAt.toLocaleDateString()} {sendAt.toLocaleTimeString(
              [],
              { hour: '2-digit', minute: '2-digit' }
            )}"
          >
            <div class="px-3">
              <SendAtScheduler bind:value={sendAt} />
            </div>
          </AccordionNode>
        </Accordion>
        <div class="grid m-2">
          <Button onClick={() => update()} icon="bxs-save" title="Update" disabled={isSaveDisabled}/>
        </div>
      </div>
    </div>
  {/if}
</div>
