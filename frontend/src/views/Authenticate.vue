<script setup>
import {useRouter} from "vue-router";

const router = useRouter()


</script>

<template>
    <div class="auth-container">
        <h1 class="login-heading">
            {{ router.currentRoute.value.name === 'login' ? 'Login' : 'Register' }}
        </h1>
        <router-view></router-view>
    </div>
</template>

<style scoped>
.auth-container {
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    height: 100vh;
    background-color: var(--color-background);
    color: var(--color-text);
}

.login-heading {
    font-size: 3rem;
    font-weight: bold;
    margin-bottom: 2rem;
}
</style>