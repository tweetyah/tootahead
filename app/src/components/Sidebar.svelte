<script lang="ts">
  import { alert, auth, service } from '../store'
  import NavLink from "./NavLink.svelte";
  import Button from './Button.svelte';
  import { navigateTo } from 'svelte-router-spa';
  import { Service } from '../models';
  import colors from '../colors';

  export let closeButton: boolean
  export let onClose: Function | undefined

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
    if (value.service === "twitter") {
      service.set(Service.Twitter)
    } else {
      service.set(Service.Mastodon)
    }
    isLoggedIn = true
  })

  function logout() {
    localStorage.removeItem("auth")
    navigateTo("/login")
  }

</script>

<div id="sidebar" class="shadow-sm flex flex-col justify-left h-100 w-full m-2 p-2 rounded bg-slate-800 text-slate-100">
  <div class="flex">
    <div class="p-3 text-lg flex-1">Tweetyah!</div>
    {#if closeButton}
      <button on:click={() => onClose()}>
        <i class="bx bx-x" style="font-size: 24px;" />
      </button>
    {/if}
  </div>
  <hr class="border-slate-700 my-2" />
  <div class="flex-1 marker:flex flex-col">
    <NavLink title="Home" icon="bx-home" to="/" />
  </div>
  {#if isLoggedIn}
    <div class="bg-slate-600 flex rounded shadow-sm hover:shadow-md p-1 text-slate-50">
      <div class="relative mr-1">
        <img src={profilePicUrl} class="rounded-full m-0.5 w-[54px] border-white" style="border-width: 1px;" alt="avatar" />
        <div class="grid content-center justify-center w-[20px] h-[20px] rounded-full absolute top-0 border-white" style={`background-color: ${$service === Service.Twitter ? colors.twitter : colors.mastodon }; border-width: 1px; right: -5px;`}>
          <i style="font-size: 12px;" class={`bx bxl-${$service === Service.Twitter ? "twitter" : "mastodon"}`}></i>
        </div>
      </div>
      <div class="ml-1 flex flex-col">
        <span class="font-bold">{ name }</span>
        <span class="italic text-sm">{ handle }</span>
      </div>
    </div>
  {/if}
  <Button title="Log out" onClick={() => logout()} />

  <!-- <NavLink title="Library" icon="bx-collection" to="/library" /> -->
</div>
