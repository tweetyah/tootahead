<script lang="ts">

  export let title: string = ""
  export let subtitle: string = ""

  function toggleNode(el: HTMLElement) {
    const parentNode = el.parentElement as HTMLElement;
    let accordionId = parentNode.dataset['accordionId']
    let accordionNodeId = parentNode.dataset['accordionNodeId']

    const accordionNodes = document.querySelectorAll(`[data-accordion-id="${accordionId}"]`)
    accordionNodes.forEach((node: HTMLElement) => {
      const nodeContentEl = node.querySelector(".accordion-content") as HTMLElement
      if(node.dataset['accordionNodeId'] === accordionNodeId) {
        let accordionMaxHeight = nodeContentEl.style.maxHeight;

        // Check if the element is already collapsed
        if (accordionMaxHeight == "0px" || accordionMaxHeight.length == 0) {
          nodeContentEl.style.maxHeight = `${nodeContentEl.scrollHeight + 32}px`;
          nodeContentEl.classList.add('border-b-[1px]')
        } else {
          nodeContentEl.style.maxHeight = `0px`;
          nodeContentEl.classList.remove('border-b-[1px]')
        }
      } else {
        nodeContentEl.style.maxHeight = `0px`;
        nodeContentEl.classList.remove('border-b-[1px]')
      }
    })
  }
</script>

<div class="accordion-node transition">
  <!-- svelte-ignore a11y-click-events-have-key-events -->
  <div
    on:click={e => toggleNode(e.currentTarget)} 
    class="accordion-header cursor-pointer transition flex space-x-5 px-5 items-center h-14 border-b-[1px]">
    <div>{ title.toUpperCase() }</div> 
    <div>{ subtitle }</div>
  </div>
  <div class="accordion-content overflow-hidden max-h-0">
    <slot></slot>
  </div>
</div>

<style lang="postcss">
  .accordion-content {
    transition: max-height 0.3s ease-out, padding 0.3s ease;
  }
</style>