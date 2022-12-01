<script lang="ts">
  import { onMount } from 'svelte';
  import { text } from 'svelte/internal';
  import Loading from '../components/Loading.svelte';
  import PostCard from '../components/PostCard.svelte';
  import View from '../components/View.svelte';
  import { ViewState, type Post } from '../models'
  import { api } from '../store';

  let postmap: {[key: string]: Post[]} = {};
  let state: ViewState = ViewState.None

  onMount(async () => {
    await load()
  })

  async function load() {
    state = ViewState.Loading
    let pm: {[key: string]: Post[]} = {};

    let posts = await $api.fetchPosts("scheduled")
    if(posts && posts.length > 0) {
      posts.sort((a, b) => a.sendAt && b.sendAt ? a.sendAt.getTime() - b.sendAt.getTime() : 0)
      posts.forEach(p => {
        let dateStr = p.sendAt.toLocaleDateString()
        if(!pm[dateStr]) pm[dateStr] = []
        pm[dateStr].push(p)
      })
      state = ViewState.Done
    } else {
      state = ViewState.NoData
    }
    postmap = pm
  }

  async function onPostUpdated() {
    await load()
  }

</script>

<View title="Scheduled posts">
  {#if state == ViewState.NoData}
    <div class="text-gray-600 italic">
      You have no scheduled posts.
    </div>
  {/if}
  {#if state == ViewState.Loading}
    <Loading />
  {/if}
  {#if state == ViewState.Done}
    <div class="grid grid-cols-[min-content_1fr] items-center justify-center">
      {#each Object.keys(postmap) as date}
        <div class="rounded-full border-mastodon border-2 text-white bg-mastodon grid items-center justify-center w-[30px] h-[30px] mx-2">
          <i class="bx bx-calendar" />
        </div>
        <h2>{ date }</h2>
        <div class="grid items-center justify-center h-full">
           <div class="bg-mastodon w-[2px] h-full">

           </div>
        </div>
        <div class="grid gap-2 grid-cols-1 lg:grid-cols-3 mb-4">
          {#each postmap[date] as post}
            <PostCard post={post} onUpdated={() => onPostUpdated()} editable/>
          {/each}
        </div>
      {/each}
    </div>
  {/if}
</View>