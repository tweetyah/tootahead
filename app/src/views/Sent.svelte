<script lang="ts">
  import { onMount } from 'svelte';
  import { text } from 'svelte/internal';
  import PostCard from '../components/PostCard.svelte';
  import View from '../components/View.svelte';
  import type { Post } from '../models'
  import { api } from '../store';
  import { SortPosts } from '../utils';

  let postmap: {[key: string]: Post[]} = {};

  onMount(async () => {
    let pm: {[key: string]: Post[]} = {};

    let posts = await $api.fetchPosts("sent")
    SortPosts(posts, true)
    posts.forEach(p => {
      let dateStr = p.sendAt.toLocaleDateString()
      if(!pm[dateStr]) pm[dateStr] = []
      pm[dateStr].push(p)
    })

    Object.keys(pm).forEach(date => {
      SortPosts(pm[date])
    })

    postmap = pm
  })
</script>

<View title="Sent posts">
  <div>
    {#each Object.keys(postmap) as date}
      <div>
        <h2 class="flex items-center"><i class="bx bx-calendar mr-2" /> { date }</h2>
        <div class="grid gap-2 grid-cols-1 lg:grid-cols-3">
          {#each postmap[date] as post}
            <PostCard post={post}/>
          {/each}
        </div>
      </div>
    {/each}
  </div>
</View>