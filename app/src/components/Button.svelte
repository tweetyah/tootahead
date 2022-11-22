<script lang="ts">
  export let title: string
  export let onClick:Function
  export let icon: string = ""
  export let variant: string = ""

  let classes = variant === "outlined" ?
    "bg-inherit hover:bg-slate-100 border-[1px] border-slate-800" :
    "bg-slate-800 text-slate-50 hover:bg-slate-600";
</script>
  

<button class="{classes} shadow-sm rounded p-1 text-sm" on:click={() => onClick()}>
  {#if icon !== ""}
    <i class="bx {icon}"></i>
  {/if}
  { title }
</button>