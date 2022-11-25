<script lang="ts">
  import { autoresize } from 'svelte-textarea-autoresize'
	import { service } from "../store";
  import textvars from "../textvars";

	export let onUpdate: Function = undefined
	export let value: string
</script>

<div>
	<textarea use:autoresize
		bind:value={value}
		placeholder={textvars[$service]["post-box-placeholder"]}
		class="focus:outline-none w-full"
		on:input={onUpdate ? onUpdate() : null} />
</div>