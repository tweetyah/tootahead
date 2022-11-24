<script>
  import { autoresize } from 'svelte-textarea-autoresize'
	import { service } from "../store";
  import textvars from "../textvars";
</script>

<div>
	<textarea use:autoresize
		placeholder={textvars[$service]["post-box-placeholder"]}
		class="focus:outline-none w-full" />
</div>