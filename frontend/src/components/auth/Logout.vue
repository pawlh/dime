<script setup>
import {useStateStore} from "@/store/state";

const stateStore = useStateStore();

// There is currently no backend support for deactivating JWTs, so temporarily just clear user from store
stateStore.loggedIn = false;
stateStore.loggedInUser.name = '';
stateStore.transactions = [];


</script>

<template>
  <h2>You have been logged out.</h2>
  <p>
    <router-link class="login-link" to="/login">Go back to the login page</router-link>
  </p>
</template>

<style>

.login-link {
  color: var(--color-text);
  text-decoration: none;
}
</style>
