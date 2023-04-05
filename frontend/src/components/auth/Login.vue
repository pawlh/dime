<script setup>
import {useRouter} from "vue-router";
import {useStateStore} from "@/store/state";
import {SERVER_URL} from "@/store/app";
import {ref} from "vue";
import {useCookies} from "vue3-cookies";

const stateStore = useStateStore();
const { cookies } = useCookies();

const router = useRouter();

const username = ref('')
const password = ref('')

const error = ref('')

const login = async () => {
    try {
        const res = await fetch(SERVER_URL + "/api/login", {
            method: 'POST',
            headers: {
                'Content-Type': 'application/json'
            },
            body: JSON.stringify({
                username: username.value,
                password: password.value
            })
        })

        if (res.status === 200) {
            const data = await res.json()
            error.value = ''
            stateStore.loggedInUser.name = data.name;
            stateStore.loggedIn = true;

            // TEMPORARY: ultimately the token will be stored as an httpOnly cookie
            cookies.set('token', data.token, {
                expires: 4 * 60 * 60
            })

            await router.push({name: 'home'});
        } else if (res.status === 401) {
            error.value = "Invalid username or password"
        } else {
            error.value = "Unknown error occurred"
            console.error(await res.text())
        }
    } catch (e) {
        console.log(`Something went wrong: ${e}`)
    }
}
</script>

<template>
    <form class="login-form">
        <label for="username" class="login-label">Username</label>
        <input type="text"
               v-model="username"
               id="username"
               class="login-input"
               placeholder="Enter your username"
               :class="{ 'border-red': username === '' }"/>

        <label for="password" class="login-label">Password</label>
        <input type="password"
               v-model="password"
               id="password"
               class="login-input"
               placeholder="Enter your password"
               :class="{ 'border-red': password === '' }"/>

        <button
                class="login-button"
                @click.prevent="login"
                :disabled="username === '' || password === ''">Login
        </button>
        <p v-if="error" class="error-message">{{ error }}</p>
        <span>
            Need an account?
            <br/>
            <router-link to="/register">Click here to register</router-link>
        </span>
    </form>
    <!--Add register link-->


</template>

<style>


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

.login-button[disabled] {
    background-color: #ccc;
    cursor: not-allowed;
}

.error-message {
    margin-top: 1rem;
    color: red;
}

.border-red {
    border: 1px solid red;
}

span {
    padding-top: 1rem;
    font-size: 1.5rem;
    text-align: center;
}

a {
    color: var(--color-text);
    font-weight: bold;
    text-decoration: none;
}
</style>
