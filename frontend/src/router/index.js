import {createRouter, createWebHistory} from 'vue-router'
import Transactions from '@/views/Transactions.vue'
import Budgets from "@/views/Budgets.vue";
import Reports from "@/views/Reports.vue";
import Tools from "@/views/Tools.vue";
import Import from "@/views/tools/Import.vue";
import Accounts from "@/views/tools/Accounts.vue";
import Home from "@/views/Home.vue";
import Register from "@/components/auth/Register.vue";
import Login from "@/components/auth/Login.vue";
import Authenticate from "@/views/Authenticate.vue";
import Logout from "@/components/auth/Logout.vue";
import {useCookies} from "vue3-cookies";

const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'home',
            component: Home,
            redirect: '/transactions',
            children: [
                {
                    path: '/transactions',
                    name: 'transactions',
                    component: Transactions
                },
                {
                    path: '/budgets',
                    name: 'budgets',
                    component: Budgets
                },
                {
                    path: '/reports',
                    name: 'reports',
                    component: Reports
                },
                {
                    path: '/tools',
                    component: Tools,
                    name: 'tools',
                    redirect: '/tools/import',
                    children: [
                        {
                            path: 'import',
                            name: 'import',
                            component: Import
                        },
                        {
                            path: 'accounts',
                            name: 'accounts',
                            component: Accounts
                        }
                    ]
                }
            ]
        },
        {
            path: '/auth',
            name: 'auth',
            component: Authenticate,
            redirect: '/login',
            children: [
                {
                    path: '/login',
                    name: 'login',
                    component: Login
                },
                {
                    path: '/register',
                    name: 'register',
                    component: Register
                },
                {
                    path: '/logout',
                    name: 'logout',
                    component: Logout
                }
            ]
        },
    ]
})

router.beforeEach((to) => {

const { cookies } = useCookies();

    const loggedIn = cookies.get('token') !== null

    // If the user is not logged in and is not headed to the login, logout, or register pages, redirect to login
    if (!loggedIn && !['login', 'register', 'logout'].includes(to.name)) {
        return {name: 'login'}
    }

    // The user can't be logged in and be on the login or register pages
    if (loggedIn && ['login', 'register'].includes(to.name)) {
        return {name: 'home'}
    }
})

export default router;