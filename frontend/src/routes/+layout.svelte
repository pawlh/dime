<script lang="ts">
	import NavBar from '$lib/features/navbar/NavBar.svelte';
	import { isLoggedIn, user } from '$lib/stores/auth';
	import { theme } from '$lib/stores/theme';
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

<svelte:head>
	<link rel="stylesheet" href={`/${$theme}Theme.css`} />
</svelte:head>

<div class="container">
	<NavBar />
	<slot />
</div>

<style lang="scss">
	@import '/global.css';
</style>
