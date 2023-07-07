<script lang="ts">
    import NavBar from '$lib/features/navbar/NavBar.svelte';
    import {isLoggedIn, user} from '$lib/stores/auth';
    import {theme} from '$lib/stores/theme';
    import cookies from '$lib/utils/cookies';
    import {getPage, navigateTo} from '$lib/utils/navigation';
    import {onMount} from 'svelte';


    $: if (!isLoggedIn) {
        if (getPage().url.pathname !== '/auth/register') navigateTo('/auth/login');
    }

    // TODO: Remove this when there is a proper login flow
    onMount(() => {
        cookies.set('loggedIn', 'true', {path: '/'});
        user.login({
            firstName: 'Joe',
            lastName: 'Tester'
        });
    });
</script>

<div id="container" class={$theme === 'dark' ? 'dark-theme' : 'light-theme'}>
    <NavBar/>
    <h2>hi</h2>
    <slot/>
</div>

<style lang="scss" global>
  @import './global.css';

  #container {
    background-color: var(--background--color);
    color: var(--text--color--primary);
  }

</style>
