<script lang="ts">
  export let title: string = ""
  export let onClick:Function
  export let icon: string = ""
  export let variant: string = ""
  export let disabled: boolean = false

  let classes = variant === "outlined" ?
    "bg-inherit hover:bg-mastodon-hover hover:text-white border-[1px] border-mastodon" :
    "bg-mastodon text-white hover:bg-mastodon-hover disabled:bg-mastodon-disabled disabled:cursor-not-allowed";
</script>


<button class="{$$props.class} {classes} shadow-sm rounded p-1 text-sm" on:click={() => onClick()} disabled={disabled}>
  {#if icon !== ""}
    <i class="bx {icon}"></i>
  {/if}
  {#if title}
    { title }
  {/if}
  <slot />
</button>