<script lang="ts">
    import {isLoggedIn} from '$lib/stores/auth';
    import {getPage, navigateTo} from '$lib/utils/navigation';
    import {onDestroy} from 'svelte';
    import {theme} from "$lib/stores/theme";


    $: if (!$isLoggedIn) {
        console.log('is not logged in')
        if (getPage().url.pathname !== '/register') navigateTo('/login');
    }
    const themeUnsubscribe = theme.subscribe(value => {
        document.documentElement.setAttribute('data-theme', value);
    });

    onDestroy(themeUnsubscribe);

    // TODO: Remove this when there is a proper login flow
    // onMount(() => {
    //     cookies.set('loggedIn', 'true', {path: '/'});
    //     user.login({
    //         firstName: 'Joe',
    //         lastName: 'Tester'
    //     });
    // });

</script>

<div id="container">
    {$isLoggedIn}
    <slot/>
</div>


<style lang="scss">
  @import 'global.css';

  #container {
    background-color: var(--background--color);
    color: var(--text--color--primary);
  }
</style>