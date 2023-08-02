<script lang="ts">

    // onMount(() => {
    // 	if (isLoggedIn) {
    // 		navigateTo('/');
    // 	}
    // });
</script>

<div id="container">
    <div class="auth-section">
        <slot/>
    </div>
</div>

<style>
    #container {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        height: 100vh;
    }

    .auth-section {
        width: 400px;
        height: 500px;
        background-color: var(--surface--color--primary);
        border-radius: 10px;
        box-shadow: 0 0 10px 0 rgba(0, 0, 0, 0.2);
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
    }
</style>