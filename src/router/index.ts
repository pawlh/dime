import {createRouter, createWebHistory} from 'vue-router'
import Transactions from '@/views/Transactions.vue'
import Budgets from "@/views/Budgets.vue";
import Reports from "@/views/Reports.vue";
import Tools from "@/views/Tools.vue";
import Import from "@/views/tools/Import.vue";
import Accounts from "@/views/tools/Accounts.vue";

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
})

export default router
