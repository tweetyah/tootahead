<script lang="ts">
  import { onMount } from 'svelte';
  import { text } from 'svelte/internal';
  import Loading from '../components/Loading.svelte';
  import PostCard from '../components/PostCard.svelte';
  import View from '../components/View.svelte';
  import { ViewState, type Post } from '../models'
  import { api } from '../store';
  import { SortPosts } from '../utils';

  let postmap: {[key: string]: Post[]} = {};
  let state: ViewState = ViewState.None

  onMount(async () => {
    state = ViewState.Loading
    let pm: {[key: string]: Post[]} = {};

    let posts = await $api.fetchPosts("sent")
    if (posts && posts.length) {
      SortPosts(posts, true)
      posts.forEach(p => {
        let dateStr = p.sendAt.toLocaleDateString()
        if(!pm[dateStr]) pm[dateStr] = []
        pm[dateStr].push(p)
      })

      Object.keys(pm).forEach(date => {
        SortPosts(pm[date])
      })
      state = ViewState.Done
    } else {
      state = ViewState.NoData
    }
    postmap = pm
  })
</script>

<View title="Sent posts">
  {#if state == ViewState.NoData}
    <div class="text-gray-600 italic">
      You have no sent posts.
    </div>
  {/if}
  {#if state == ViewState.Loading}
    <Loading />
  {/if}
  {#if state == ViewState.Done}
    <div>
      {#each Object.keys(postmap) as date}
      <div>
        <h2>{ date }</h2>
        <div class="grid gap-2 grid-cols-1 lg:grid-cols-3">
          {#each postmap[date] as post}
            <PostCard post={post}/>
          {/each}
        </div>
      </div>
      {/each}
    </div>
  {/if}
</View>