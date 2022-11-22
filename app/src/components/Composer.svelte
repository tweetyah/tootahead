<script lang="ts">
  import type { Tweet } from "../models";
  import { api } from "../store";
  import Accordion from "./Accordion.svelte";
  import AccordionNode from "./AccordionNode.svelte";
  import Button from "./Button.svelte";
  import ComposerCard from "./ComposerCard.svelte";
  import RetweetAtScheduler from "./RetweetAtScheduler.svelte";
  import SendAtScheduler from "./SendAtScheduler.svelte";

  let sendAt: Date = new Date()
  let retweetAt: Date
  let shouldRetweet: boolean
  let tweets: Tweet[] = [{
    text: ""
  }]

  function addTweet() {
    tweets = [...tweets, {
      text: ""
    }]
  }

  async function saveTweets() {
    // TODO: app or comp state
    tweets.forEach(t => {
      t.sendAt = sendAt
      if(shouldRetweet) {
        t.retweetAt = retweetAt
      }
    })
    await $api.saveTweets(tweets)
    reset()
  }

  function reset() {
    tweets = [{
      text: ""
    }]
  }
</script>

<div>
  <div class="grid grid-cols-2 gap-2">
    <div id="composer-wrapper">
      <div class="bg-white shadow-sm rounded mb-2">
        {#each tweets as t, idx}
          <ComposerCard
            bind:tweet={t}
            index={idx}
            total={tweets.length} />
        {/each}
      </div>
      <Button onClick={() => addTweet()} icon="bx-list-plus" title="Add tweet" />
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