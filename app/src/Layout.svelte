<script lang="ts">
  import { Route } from 'svelte-router-spa'
  import AlertLayer from './components/AlertLayer.svelte';
  import Navbar from './components/Navbar.svelte';
  import Sidebar from './components/Sidebar.svelte';
  export let currentRoute
  const params = {}
</script>

<div class="h-full w-full flex sm:flex-row flex-col">
  <Navbar />
  <div class="h-0 sm:h-full flex sm:w-[300px] w-0 invisible sm:visible">
    <Sidebar />
  </div>
  <main class="flex-1">
    <Route {currentRoute} {params} />
  </main>
  <AlertLayer />
</div>