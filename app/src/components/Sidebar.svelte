<script lang="ts">
  import { auth } from '../store'
  import NavLink from "./NavLink.svelte";
  import Button from './Button.svelte';
  import { navigateTo } from 'svelte-router-spa';

  let name: string
  let handle : string
  let profilePicUrl: string
  let isLoggedIn: boolean

  // TODO: type this
  auth.subscribe((value: any) => {
    if(!value || !value.name) {
      navigateTo("/login")
      return
    }
    name = value.name
    handle = value.username
    profilePicUrl = value.profile_image_url
    isLoggedIn = true
  })

  function logout() {
    localStorage.removeItem("auth")
    navigateTo("/login")
  }

</script>

<div id="sidebar" class="shadow-sm flex flex-col justify-left h-100 w-[250px] m-2 p-2 rounded bg-slate-800 text-slate-100">
  <div class="p-3 text-lg">Tweetyah!</div>
  <hr class="border-slate-700 my-2" />
  <div class="flex-1 marker:flex flex-col">
    <NavLink title="Home" icon="bx-home" to="/" />
  </div>
  {#if !isLoggedIn}
    <a>Login</a>
  {:else}
    <div class="bg-slate-600 flex rounded shadow-sm hover:shadow-md p-1 text-slate-50">
      <img src={profilePicUrl} class="rounded-full m-0.5 w-[50px] h-[50px]" />
      <div class="ml-1 flex flex-col">
        <span class="font-bold">{ name }</span>
        <span class="italic text-sm">{ handle }</span>
      </div>
    </div>
    {/if}
    <Button title="Log out" onClick={() => logout()} />

  <!-- <NavLink title="Library" icon="bx-collection" to="/library" /> -->
</div>