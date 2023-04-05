<script setup>
import DataTable from "@/components/DataTable.vue";
import {useStateStore} from "@/store/state";

const headers = [
    {label: "id", sortable: false, hidden: true},
    {label: "date", sortable: true, formatter: value => (new Date(value)).toLocaleDateString()},
    {label: "description", sortable: false},
    {label: "amount", sortable: true, formatter: value => `$${value}`},
    {label: "category", sortable: true},
    {label: "account", sortable: true},
];

const stateStore = useStateStore()

</script>

<template>
    <DataTable :headers="headers"
               :rows="stateStore.transactions"
               :paginate="true"
               :rows-per-page="10"/>
</template>

<style scoped>

</style>