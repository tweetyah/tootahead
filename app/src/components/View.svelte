<script lang="ts">

  export let title: string;
</script>

<div class="overflow-y-scroll">
  <h1>{ title }</h1>
  <slot />
</div>
