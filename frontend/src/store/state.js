import {defineStore} from "pinia";
import {ref} from "vue";

export const useStateStore = defineStore('user', () => {
    const loggedIn = ref(false)
    const loggedInUser = ref({
        name: '',
    })

    const transactions = ref([])

    return {loggedIn, loggedInUser, transactions}
})

// load store values from localStorage
// useStateStore().loggedIn.value = localStorage.getItem('loggedIn') === 'true'
// useStateStore().loggedInUser.value = JSON.parse(localStorage.getItem('loggedInUser'))
