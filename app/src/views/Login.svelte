<script lang="ts">
  import { fade } from "svelte/transition";

  let isCheckingDomain = false;
  let instanceUrl = "";
  let isLoginDisabled = true;
  let isMastodonInstanceInvalid = false;
  let timer
  $: if (isMastodonInstanceInvalid === true) {
    timer = setTimeout(() => {
      isMastodonInstanceInvalid = false
    }, 4000)
  }

  $: if (instanceUrl != undefined) {
    isLoginDisabled = !/^[a-zA-Z0-9][a-zA-Z0-9-]{1,61}[a-zA-Z0-9](?:\.[a-zA-Z]{2,})+$/.test(instanceUrl)
  }

  async function validateMastodonInstance() {
    try {
      isCheckingDomain = true
      let url = `https://${instanceUrl}/api/v2/instance`
      let res = await fetch(url)
      let json = await res.json()
      if(json && json.version) {
        await loginToMastodon()
      } else {
        isMastodonInstanceInvalid = true
      }
    } catch (err) {
      isMastodonInstanceInvalid = true
    } finally {
      isCheckingDomain = false
    }
  }

  async function loginToMastodon() {
    let res = await fetch(`/.netlify/functions/mastodon_app?domain=${instanceUrl}`)
    let json = await res.json()
    let url = `https://${instanceUrl}/oauth/authorize?`
    url += `&client_id=${json.clientId}`
    url += `&redirect_uri=${import.meta.env.VITE_REDIRECT_URI}`
    url += `&scope=read:accounts write:statuses`
    url += '&grant_type=authorization_code'
    url += '&response_type=code'
    url += '&state=' + instanceUrl
    location.href = url
  }

  function loginWithTwitter() {
    let loginUrl = "https://twitter.com/i/oauth2/authorize"
    loginUrl += `?response_type=code`
    loginUrl += `&client_id=${import.meta.env.VITE_TWITTER_CLIENT_ID}`
    loginUrl += `&redirect_uri=${import.meta.env.VITE_REDIRECT_URI}`
    loginUrl += `&scope=tweet.read tweet.write users.read offline.access`
    loginUrl += `&code_challenge=challenge`
    loginUrl += `&code_challenge_method=plain`
    loginUrl += `&state=twitter`
    location.href = loginUrl
  }
</script>

<div class="grid justify-center content-center">
  <div class="mt-20">
    <!-- <h1>Login</h1> -->
  </div>
  <div class="bg-white border-1 rounded w-[300px] mb-2">
    <div class="w-full p-3 border-b-slate-50 border-b-2">
      <input bind:value={instanceUrl} class="w-full p-2 mb-2 rounded bg-slate-50" type="text" placeholder="Mastodon domain (ie; fosstodon.org)">
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