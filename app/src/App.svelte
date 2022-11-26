<script lang="ts">
  import { init } from './store'
  import { Router } from 'svelte-router-spa'
  import { routes } from './router'
  import { onMount } from 'svelte';
  import { ApiService } from './services/ApiService';
  import type { Auth } from './models';

  onMount(async () => {
    await init()
  })
</script>

<Router {routes} />