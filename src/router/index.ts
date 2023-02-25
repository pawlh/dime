import {createRouter, createWebHistory} from 'vue-router'
import Transactions from '@/views/Transactions.vue'
import Budgets from "@/views/Budgets.vue";
import Reports from "@/views/Reports.vue";
//
const router = createRouter({
    history: createWebHistory(import.meta.env.BASE_URL),
    routes: [
        {
            path: '/',
            name: 'home',
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
        }
    ]
})

export default router
