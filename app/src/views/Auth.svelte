<script lang="ts">
  import { onMount } from "svelte";
  import { navigateTo } from "svelte-router-spa";
  import { Service } from "../models";
  import { ApiService } from "../services/ApiService";
  import { init, service } from '../store'

  onMount(async () => {
    let query = location.search
    query = query.replace("?", "")
    let spl = query.split("&")
    let map = {}
    spl.forEach(kvp => {
      map[kvp.split("=")[0]] = kvp.split("=")[1]
    })

    let res = await fetch("/.netlify/functions/auth", {
      method: "post",
      headers: {
        "Content-Type": "application/json"
      },
      body: JSON.stringify({
        code: map["code"],
        state: map["state"]
      })
    })
    let json = await res.json()

    localStorage.setItem("auth", JSON.stringify(json))
    await init()

    if(map["state"] === "twitter") {
      service.set(Service.Twitter)
    } else {
      service.set(Service.Mastodon)
    }

    navigateTo("/")
  })
</script>

<div class="grid justify-center">
  <i style="font-size: 48px;" class='bx bx-loader-alt bx-spin bx-flip-horizontal mt-4' ></i>
</div>