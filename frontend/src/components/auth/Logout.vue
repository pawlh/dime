<script setup>
import {useStateStore} from "@/store/state";
import {useCookies} from "vue3-cookies";
import {ref, watch} from "vue";
import router from "@/router";

const stateStore = useStateStore();

const { cookies } = useCookies()

// There is currently no backend support for deactivating JWTs, so temporarily just clear user from store
stateStore.loggedIn = false;
stateStore.loggedInUser.name = '';
stateStore.transactions = [];
cookies.remove('token')

const countdown = ref(3);

//decrement countdown every second
setInterval(() => {
  countdown.value--;
}, 1000);

watch(countdown, (newCountdown) => {
  if (newCountdown === 0)
    router.push({name: 'login'})
})


</script>

<template>
  <h2>You have been logged out. Redirecting in {{countdown}}</h2>
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
