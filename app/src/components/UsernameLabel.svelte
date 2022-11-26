<script lang="ts">
  import { onMount } from "svelte";
  import { custom_emoji } from '../store'

  export let name: string
  let display
  let isReplaced = false

  onMount(() => {
    replaceEmoji()
  });

  custom_emoji.subscribe(() => {
    replaceEmoji()
  })

  function replaceEmoji() {
    if($custom_emoji && !isReplaced) {
      let emojiInName = name.match(/(?<=:)[a-zA-Z]+(?=:)/g)
      emojiInName.forEach(el => {
        let emoji = $custom_emoji.find(ce => ce.shortcode === el)
        if (emoji) {
          name = name.replace(`:${el}:`, `<img src="${emoji.static_url}" class="w-auto max-w-[16px] h-auto max-h-[16px] ml-1" />`)
        }
      })
      display = name
      isReplaced = true
    }
  }
</script>

<div class={`flex items-center ${$$props.class}`}>
  {@html name}
</div>