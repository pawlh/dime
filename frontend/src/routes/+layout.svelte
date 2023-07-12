<script lang="ts">
    import NavBar from '$lib/features/navbar/NavBar.svelte';
    import {isLoggedIn, user} from '$lib/stores/auth';
    import cookies from '$lib/utils/cookies';
    import {getPage, navigateTo} from '$lib/utils/navigation';
    import {onDestroy, onMount} from 'svelte';
    import ThemeChooser from "$lib/features/navbar/ThemeChooser.svelte";
    import {theme} from "$lib/stores/theme";


    $: if (!isLoggedIn) {
        if (getPage().url.pathname !== '/auth/register') navigateTo('/auth/login');
    }

    const themeUnsubscribe = theme.subscribe(value => {
        document.documentElement.setAttribute('data-theme', value);
    });

    onDestroy(themeUnsubscribe);

    // TODO: Remove this when there is a proper login flow
    onMount(() => {
        cookies.set('loggedIn', 'true', {path: '/'});
        user.login({
            firstName: 'Joe',
            lastName: 'Tester'
        });
    });
</script>

<div id="container">
    <div class="nav-bar">
        <NavBar/>
    </div>
    <div class="main">
        <div class="status-bar">
            <ThemeChooser/>
        </div>
        <slot/>
    </div>
</div>

<style lang="scss">
  @import './global.css';

  #container {
    background-color: var(--background--color);
    color: var(--text--color--primary);

    display: flex;
    flex-direction: column;
    height: 100vh;

  }

  .nav-bar {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    z-index: 100;
  }

  @media (min-width: 768px) {
    #container {
      flex-direction: row;
    }

    .nav-bar {
      position: static;
      width: 300px;
    }
  }
</style>
