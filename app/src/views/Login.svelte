<script lang="ts">
  let instanceUrl;

  function redirectToLogin() {
    let url = `https://${instanceUrl}/oauth/authorize?`
    url += `&client_id=${import.meta.env.VITE_MASTODON_CLIENT_ID}`
    url += `&redirect_uri=${import.meta.env.VITE_REDIRECT_URI}`
    url += `&scope=read write follow`
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

<div class="grid h-full justify-center content-center">
  <div>
    <!-- logo here -->
  </div>
  <div class="bg-white border-1 rounded w-[300px]">
    <div class="w-full p-3 border-b-slate-50 border-b-2">
      <input bind:value={instanceUrl} class="w-full p-2 mb-2 rounded bg-slate-50" type="text" placeholder="Mastodon domain (ie; fosstodon.org)">
      <button on:click={e => redirectToLogin(e)} class="w-full p-2 rounded text-white" style="background-color: #6364FF">Login with Mastodon</button>
    </div>
    <div class="p-3">
      <button on:click={loginWithTwitter} class="w-full p-2 rounded text-white" style="background-color: #1DA1F2">Login with Twitter</button>
    </div>
  </div>
</div>