import {defineStore} from "pinia";
import {ref} from "vue";

export const useStateStore = defineStore('counter', () => {
    const loggedIn = ref(false)
    const loggedInUser = ref({
        firstName: '',
        lastName: '',
    })

    return {loggedIn, loggedInUser}
})

// load store values from localStorage
// useStateStore().loggedIn.value = localStorage.getItem('loggedIn') === 'true'
// useStateStore().loggedInUser.value = JSON.parse(localStorage.getItem('loggedInUser'))
