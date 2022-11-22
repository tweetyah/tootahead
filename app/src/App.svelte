<script lang="ts">
  import { api, setAuth } from './store'
  import { Router } from 'svelte-router-spa'
  import { routes } from './router'
  import { onMount } from 'svelte';
  import { ApiService } from './services/ApiService';
  import type { Auth } from './models';

  onMount(() => {
    let authItem = localStorage.getItem("auth")
    if(authItem) {
      let auth: Auth = JSON.parse(authItem)
      setAuth(auth)

      api.set(new ApiService(auth.access_token))
    }
  })
</script>

<Router {routes} />