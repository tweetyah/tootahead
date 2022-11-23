<script lang="ts">
  import type { Post } from "../models";
  import { api, service } from "../store";
  import textvars from "../textvars";
  import Accordion from "./Accordion.svelte";
  import AccordionNode from "./AccordionNode.svelte";
  import Button from "./Button.svelte";
  import ComposerCard from "./ComposerCard.svelte";
  import RetweetAtScheduler from "./RetweetAtScheduler.svelte";
  import SendAtScheduler from "./SendAtScheduler.svelte";

  let sendAt: Date = new Date()
  let retweetAt: Date
  let shouldRetweet: boolean
  let posts: Post[] = [{
    text: ""
  }]

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
  }

  function reset() {
    posts = [{
      text: ""
    }]
  }
</script>

<div>
  <div class="grid grid-cols-2 gap-2">
    <div id="composer-wrapper">
      <div class="bg-white shadow-sm rounded mb-2">
        {#each posts as p, idx}
          <ComposerCard
            bind:post={p}
            index={idx}
            total={posts.length} />
        {/each}
      </div>
      <Button onClick={() => addTweet()} icon="bx-list-plus" title={textvars[$service]["add-post"]} />
      <Button onClick={() => saveTweets()} icon="bxs-save" title="Save" />
    </div>
    <div id="composer-preview">
      <Accordion>
        <AccordionNode title="Send at"
          subtitle="{sendAt.toLocaleDateString()} {sendAt.toLocaleTimeString()}">
          <div class="px-3">
            <SendAtScheduler bind:value={sendAt} />
          </div>
        </AccordionNode>
        <AccordionNode title="Retweet at" subtitle={shouldRetweet ? `${retweetAt.toLocaleDateString()} ${retweetAt.toLocaleTimeString()}` : 'Off'}>
          <div class="px-3">
            <RetweetAtScheduler bind:value={retweetAt} bind:isEnabled={shouldRetweet} />
          </div>
        </AccordionNode>
        <!-- <AccordionNode title="Categories">
          Categories
        </AccordionNode>
        <AccordionNode title="Other">
          Save to library opt
        </AccordionNode> -->
      </Accordion>
    </div>
  </div>
</div>