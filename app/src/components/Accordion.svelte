
<script lang="ts">
  import { afterUpdate, onMount } from "svelte";

  let accordionId

  onMount(() => {
    accordionId = Math.floor(Math.random() * 10000);
  });

  afterUpdate(() => {
    const nodes = document.querySelectorAll(".accordion-node");
    nodes.forEach((node: HTMLElement, idx: number) => {
      if(!node.dataset['accordionId']) {
        node.dataset['accordionId'] = accordionId
        node.dataset['accordionNodeId'] = `${accordionId}-${idx}`
      }
    });
  })
</script>

<div class="bg-white rounded shadow-sm">
  <slot></slot>
</div>