<script lang="ts">
  import { onMount } from 'svelte';
  import { text } from 'svelte/internal';
  import PostCard from '../components/PostCard.svelte';
  import View from '../components/View.svelte';
  import type { Post } from '../models'
  import { api } from '../store';

  let postmap: {[key: string]: Post[]} = {};
  let isNoContent: boolean

  onMount(async () => {
    await load()
  })

  async function load() {
    let pm: {[key: string]: Post[]} = {};

    let posts = await $api.fetchPosts("scheduled")
    console.log('posts', posts)
    if(posts && posts.length > 0) {
      posts.sort((a, b) => a.sendAt && b.sendAt ? a.sendAt.getTime() - b.sendAt.getTime() : 0)
      posts.forEach(p => {
        let dateStr = p.sendAt.toLocaleDateString()
        if(!pm[dateStr]) pm[dateStr] = []
        pm[dateStr].push(p)
      })
    } else {
      isNoContent = true
    }
    postmap = pm
  }

  async function onPostUpdated() {
    console.log('onpostupdated')
    await load()
  }

</script>

<View title="Scheduled posts">
  {#if isNoContent}
    <div>
      You have no scheduled posts...
    </div>
  {/if}
  <div>
    {#each Object.keys(postmap) as date}
      <div>
        <h2>{ date }</h2>
        <div class="grid gap-2 grid-cols-1 lg:grid-cols-3">
          {#each postmap[date] as post}
            <PostCard post={post} onUpdated={() => onPostUpdated()} editable/>
          {/each}
        </div>
      </div>
    {/each}
  </div>
</View>