<script lang="ts">
	import { isLoggedIn, user } from '$lib/stores/auth';
	import cookies from '$lib/utils/cookies';
	import { navigateTo, getPage } from '$lib/utils/navigation';
	import { onMount } from 'svelte';

	$: if (!isLoggedIn) {
		if (getPage().url.pathname !== '/auth/register') navigateTo('/auth/login');
	}

	// TODO: Remove this when there is a proper login flow
	onMount(() => {
		cookies.set('loggedIn', 'true', { path: '/' });
		user.login({
			firstName: 'Joe',
			lastName: 'Tester'
		});
	});
</script>

<slot />
