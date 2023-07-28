<script lang="ts">
	import { isLoggedIn } from '$lib/stores/auth';
	import { navigateTo } from '$lib/utils/navigation';
	import { onMount } from 'svelte';

	onMount(() => {
		if (isLoggedIn) {
			navigateTo('/');
		}
	});
</script>

<slot />
