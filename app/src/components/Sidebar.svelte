<script lang="ts">
  import { alert, auth, service } from '../store'
  import NavLink from "./NavLink.svelte";
  import NavHeader from './NavHeader.svelte';
  import Button from './Button.svelte';
  import { navigateTo } from 'svelte-router-spa';
  import { Service } from '../models';
  import colors from '../colors';
  import UsernameLabel from './UsernameLabel.svelte';

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

<div id="sidebar" class="shadow-sm flex flex-col justify-left h-100 w-full m-2 p-2 rounded bg-dark1 text-slate-100">
  <div class="flex">
    <div class="p-3 flex-1 flex text-2xl">
      <img src="/img/logo.png" alt="logo" class="w-[30px] h-auto mr-2"/>
      <span>TootAhead!</span>
    </div>
    {#if closeButton}
      <button on:click={() => onClose()}>
        <i class="bx bx-x" style="font-size: 24px;" />
      </button>
    {/if}
  </div>
  <hr class="border-dark2 my-2" />
  <div class="flex-1 marker:flex flex-col">
    <NavHeader>Posts</NavHeader>
    <NavLink title="Home" icon="bx-home" to="/" />
    <NavLink title="Scheduled" icon="bx-time" to="/scheduled" />
  </div>
  {#if isLoggedIn}
    <div class="bg-dark2 flex rounded shadow-sm hover:shadow-md p-1 text-slate-50">
      <div class="relative mr-1">
        <img src={profilePicUrl} class="rounded-full m-0.5 w-[54px] border-white" style="border-width: 1px;" alt="avatar" />
        <div class="grid content-center justify-center w-[20px] h-[20px] rounded-full absolute top-0 border-white" style={`background-color: ${$service === Service.Twitter ? colors.twitter : colors.mastodon }; border-width: 1px; right: -5px;`}>
          <i style="font-size: 12px;" class={`bx bxl-${$service === Service.Twitter ? "twitter" : "mastodon"}`}></i>
        </div>
      </div>
      <div class="ml-1 flex flex-col">
        <UsernameLabel class="font-bold" name={name} />
        <!-- <span class="font-bold">{ name }</span> -->
        <span class="italic text-sm">{ handle }</span>
      </div>
    </div>
  {/if}
  <Button title="Log out" onClick={() => logout()} />

  <!-- <NavLink title="Library" icon="bx-collection" to="/library" /> -->
</div>
