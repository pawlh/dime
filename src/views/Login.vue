<script setup>
import {useRouter} from "vue-router";
import {useStateStore} from "@/store/state";

const stateStore = useStateStore();

const router = useRouter();
const login = () => {
  stateStore.loggedInUser.firstName = "John";
  stateStore.loggedInUser.lastName = "Smith";
  stateStore.loggedIn = true;
  // router.push({name: 'home'});
};
</script>

<template>
  <div class="login-container">
    <h1 class="login-heading">Login</h1>
    <form class="login-form">
      <label for="username" class="login-label">Username</label>
      <input type="text" id="username" class="login-input" placeholder="Enter your username"/>

      <label for="password" class="login-label">Password</label>
      <input type="password" id="password" class="login-input" placeholder="Enter your password"/>

      <button class="login-button" @click="login">Login</button>
    </form>
  </div>
</template>

<style>
.login-container {
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

.login-form {
  display: flex;
  flex-direction: column;
  align-items: center;
  background-color: var(--color-background-soft);
  padding: 3rem;
  border-radius: 0.5rem;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.2);
}

.login-label {
  font-size: 1.2rem;
  margin-bottom: 0.5rem;
}

.login-input {
  padding: 0.5rem;
  font-size: 1rem;
  border: none;
  border-radius: 0.2rem;
  margin-bottom: 1rem;
  box-shadow: 0 0 5px rgba(0, 0, 0, 0.1);
}

.login-button {
  padding: 0.5rem 1rem;
  font-size: 1.2rem;
  border: none;
  border-radius: 0.2rem;
  background-color: #4CAF50;
  color: white;
  cursor: pointer;
  box-shadow: 0 0 5px rgba(0, 0, 0, 0.1);
}

.login-button:hover {
  background-color: #3e8e41;
}
</style>
