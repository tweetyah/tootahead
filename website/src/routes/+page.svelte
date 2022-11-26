<script>
  function redirectToApp() {
    location.href = "https://app.tootahead.com"
  }
</script>

<div class="h-full bg-slate-100 grid justify-center content-start md:content-center">
  <div class="grid grid-cols-1 md:grid-cols-2 bg-white rounded p-4">
    <div class="mb-6">
      <div class="italic text-gray-700">Introducing:</div>
      <div class="font-bold text-4xl flex my-2">
        <img src="/img/logo.png" alt="logo" class="w-[50px] h-auto"/> 
        <span>TootAhead!</span>
      </div>
      <div class="text-lg">
        <div class="mb-2">A new post scheduler for Mastodon users.</div>
        <ul class="list-disc ml-10 mb-4">
          <li>Log into any instance.</li>
          <li>Easily schedule posts.</li>
          <li>Fully <a href="https://github.com/tweetyah/tootahead" target="_blank" rel="noreferrer">open source</a>!</li>
        </ul>
        <button 
          class="p-2 px-6 bg-slate-900 text-slate-100 rounded hover:bg-slate-800 w-full md:w-auto" 
          on:click={() => redirectToApp()}>
          Try it now
        </button>
      </div>
    </div>
    <div class="grid grid-cols-2 gap-2 justify-center place-content-center">
      <img src="/img/home-ui.png" alt="home ui" class="max-h-[400px] shadow-md border-2 rounded-lg"/>
      <img src="/img/scheduler-ui.png" alt="home ui" class="max-h-[400px] shadow-md border-2 rounded-lg"/>
    </div>
  </div>
</div> 