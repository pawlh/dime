<script setup>
import DataTable from "@/components/DataTable.vue";
import {useStateStore} from "@/store/state";
import {SERVER_URL} from "@/store/app";

const headers = [
    {label: "id", sortable: false, hidden: true},
    {label: "date", sortable: true, formatter: value => (new Date(value)).toLocaleDateString()},
    {label: "description", sortable: false},
    {label: "amount", sortable: true, formatter: value => `$${value}`},
    {label: "category", sortable: true},
    {label: "account", sortable: true},
];

const stateStore = useStateStore()

fetchTransactions()
async function fetchTransactions() {
    const res = await fetch(SERVER_URL + '/api/transactions', {
        method: 'GET',
        credentials: 'include'
    })
    const data = await res.json()

    if (!res.ok) {
        console.error('Failed to fetch transactions: ' + res.status + ' ' + res.error)
        return
    }

    stateStore.transactions = data
}

</script>

<template>
    <DataTable :headers="headers"
               :rows="stateStore.transactions"
               :paginate="true"
               :rows-per-page="10"/>
</template>

<style scoped>

</style>