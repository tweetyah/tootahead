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

  let timeout
  let isMobileDrawerOpen = false
  let sendAt: Date = new Date()
  $: if(sendAt) calcIsSaveDisabled()
  let retweetAt: Date
  let shouldRetweet: boolean
  let posts: Post[] = [{
    text: ""
  }]

  onMount(() => {
    sendAt = addMinutes(new Date(), 15)
    timeout = setTimeout(() => {
      console.log("timeout hit")
      calcIsSaveDisabled()
    }, 30000)
  })

  onDestroy(() => {
    timeout = undefined
  })

  function addTweet() {
    posts = [...posts, {
      text: ""
    }]
  }

  async function saveTweets() {
    // TODO: app or comp state
    posts.forEach(t => {
      t.sendAt = sendAt
      if(shouldRetweet) {
        t.retweetAt = retweetAt
      }
    })
    await $api.savePosts(posts)
    reset()
    alert.set({
      title: "Post saved",
      body: "You're post was scheduled successfully!"
    })
    isMobileDrawerOpen = false
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
    posts.forEach(p => !p.text ? isDisabled = true : null )
    if(sendAt < new Date()) isDisabled = true
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
      <!-- <Button onClick={() => addTweet()} icon="bx-list-plus" title={textvars[$service]["add-post"]} /> -->
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
          <!-- <AccordionNode title="Retweet at" subtitle={shouldRetweet ? `${retweetAt.toLocaleDateString()} ${retweetAt.toLocaleTimeString([], {hour: '2-digit', minute:'2-digit'})}` : 'Off'}>
            <div class="px-3">
              <RetweetAtScheduler bind:value={retweetAt} bind:isEnabled={shouldRetweet} />
            </div>
          </AccordionNode> -->
          <!-- <AccordionNode title="Categories">
            Categories
          </AccordionNode>
          <AccordionNode title="Other">
            Save to library opt
          </AccordionNode> -->
        </Accordion>
      </div>
      <div class="flex">
        <Button
          onClick={() => saveTweets()}
          icon="bxs-save"
          title="Save"
          disabled={isSaveDisabled}
        />
      </div>
    </div>
  </div>

  <button
    on:click={() => (isMobileDrawerOpen = true)}
    class="visible sm:invisible shadow-sm hover:shadow-xl hover:bg-slate-800 absolute right-5 bottom-5 rounded-full round-full bg-slate-900 text-white grid justify-center content-center h-[48px] w-[48px] disabled:bg-slate-600 disabled:cursor-not-allowed"
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
      <div class="bg-slate-800 w-full p-4 shadow-sm text-slate-50 flex">
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
          <!-- <AccordionNode title="Retweet at" subtitle={shouldRetweet ? `${retweetAt.toLocaleDateString()} ${retweetAt.toLocaleTimeString([], {hour: '2-digit', minute:'2-digit'})}` : 'Off'}>
            <div class="px-3">
              <RetweetAtScheduler bind:value={retweetAt} bind:isEnabled={shouldRetweet} />
            </div>
          </AccordionNode> -->
          <!-- <AccordionNode title="Categories">
            Categories
          </AccordionNode>
          <AccordionNode title="Other">
            Save to library opt
          </AccordionNode> -->
        </Accordion>
        <div class="grid m-2">
          <Button onClick={() => saveTweets()} icon="bxs-save" title="Save" disabled={isSaveDisabled}/>
        </div>
      </div>
    </div>
  {/if}
</div>
