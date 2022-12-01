<script lang="ts">
  import { onMount } from "svelte";
  import { fade } from "svelte/transition";

  let isCheckingDomain = false;
  let instanceUrl = "";
  let isLoginDisabled = true;
  let isMastodonInstanceInvalid = false;
  let isInstanceSet = false;
  let setInstanceName = "";
  let setInstanceUrl = "";
  let timer
  $: if (isMastodonInstanceInvalid === true) {
    timer = setTimeout(() => {
      isMastodonInstanceInvalid = false
    }, 4000)
  }

  $: if (instanceUrl != undefined) {
    isLoginDisabled = !/^[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9](?:\.[a-zA-Z]{2,})+$/.test(instanceUrl)
  }

  onMount(() => {
    let instanceItem = localStorage.getItem("instance")
    console.log(instanceItem)
    if(instanceItem) {
      let instance = JSON.parse(instanceItem)
      setInstanceName = instance.title
      setInstanceUrl = instance.domain
      isInstanceSet = true
    }
  })

  async function validateMastodonInstance() {
    try {
      isCheckingDomain = true
      let url = `https://${instanceUrl}/api/v2/instance`
      let res = await fetch(url)
      let json = await res.json()
      if(json && json.version) {
        localStorage.setItem('instance', JSON.stringify(json))
        await loginToMastodon(instanceUrl)
      } else {
        isMastodonInstanceInvalid = true
      }
    } catch (err) {
      isMastodonInstanceInvalid = true
    } finally {
      isCheckingDomain = false
    }
  }

  async function loginToMastodon(validatedUrl: string) {
    let res = await fetch(`/.netlify/functions/mastodon_app?domain=${validatedUrl}`)
    let json = await res.json()
    let url = `https://${validatedUrl}/oauth/authorize?`
    url += `&client_id=${json.clientId}`
    url += `&redirect_uri=${import.meta.env.VITE_REDIRECT_URI}`
    url += `&scope=read:accounts write:statuses`
    url += '&grant_type=authorization_code'
    url += '&response_type=code'
    url += '&state=' + validatedUrl
    location.href = url
  }
</script>

<div class="grid justify-center content-center">
  <div class="p-3 flex-1 flex text-2xl mt-20">
    <img src="/img/logo.png" alt="logo" class="w-[30px] h-auto mr-2"/>
    <span>TootAhead!</span>
  </div>
  <div class="bg-white border-1 rounded w-[300px] mb-2 shadow-sm">
    <div class="w-full p-3 border-b-slate-50 border-b-2">
      {#if isInstanceSet}
        <div class="border-b pb-2 mb-2">
          <button on:click={() => loginToMastodon(setInstanceUrl)}
            class="w-full p-2 rounded text-white disabled:cursor-not-allowed login-btn bg-mastodon disabled:bg-mastodon-disabled">
            <i class='bx bxl-mastodon' ></i>
            Log into {setInstanceName}
          </button>
        </div>
      {/if}
      <label class="font-bold text-sm" for="insanceUrl">
        {#if isInstanceSet}
          <span>Or log into a different instance</span>
        {:else}
          <span>Mastodon URL <span class="text-red-600">*</span></span>
        {/if}
        <!-- Mastodon URL  -->
      </label>
      <input name="instanceUrl" bind:value={instanceUrl} on:keypress={e => e.key === "Enter" ? validateMastodonInstance() : null} class="w-full p-2 mb-2 rounded bg-slate-50 border border-gray-200 shadow-sm" type="text" placeholder="ex: fosstodon.org">
      <button on:click={() => validateMastodonInstance()}
        class="w-full p-2 rounded text-white disabled:cursor-not-allowed login-btn bg-mastodon disabled:bg-mastodon-disabled"
        disabled={isLoginDisabled}>
        {#if isCheckingDomain}
          <i class='bx bx-loader-alt bx-spin' ></i>
        {:else}
          <i class='bx bxl-mastodon' ></i>
        {/if}
        Login with Mastodon
      </button>
    </div>
    <!-- <div class="p-3">
      <button on:click={loginWithTwitter} class="w-full p-2 rounded text-white" style="background-color: #1DA1F2">Login with Twitter</button>
    </div> -->
  </div>
  {#if isMastodonInstanceInvalid}
    <div transition:fade class="bg-white border-1 rounded w-[300px] p-2">
      <div class="w-full">
        <i class="bx bx-error-circle text-red-400" /> This domain does not appear to be to a valid Mastodon instance.
      </div>
    </div>
  {/if}
</div>