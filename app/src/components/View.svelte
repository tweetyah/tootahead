<script lang="ts">

  export let title: string = undefined;
</script>

<div class="overflow-y-scroll p-2 h-full">
  {#if title}
    <h1>{ title }</h1>
  {/if}
  <slot />
</div>
