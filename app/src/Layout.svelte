<script lang="ts">
  import { Route } from 'svelte-router-spa'
  import Sidebar from './components/Sidebar.svelte';
  export let currentRoute
  const params = {}
</script>

<div class="h-full w-full flex">
  <Sidebar />
  <main class="flex-1 p-2">
    <Route {currentRoute} {params} />
  </main>
</div>