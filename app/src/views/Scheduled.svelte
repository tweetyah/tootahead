<script lang="ts">
  import { onMount } from 'svelte';
  import { text } from 'svelte/internal';
  import PostCard from '../components/PostCard.svelte';
  import View from '../components/View.svelte';
  import type { Post } from '../models'
  import { api } from '../store';

  let posts: Post[] = [];

  onMount(async () => {
    posts = await $api.fetchPosts()
  })

  
</script>

<View title="Scheduled posts">
  <div>
    {#each posts as post}
      <PostCard post={post} />
    {/each}
  </div>
</View>