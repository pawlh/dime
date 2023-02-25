<script setup>
import DataTable from "@/components/DataTable.vue";

const headers = [
  {label: "id", sortable: false, hidden: true},
  {label: "Date", sortable: true},
  {label: "Description", sortable: false},
  {label: "Category", sortable: true},
  {label: "Amount", sortable: true, prefix: "$"},
];

const rows = [
  {id: 1, Date: "2021-01-01", Description: "Groceries", Category: "Food", Amount: 100},
  {id: 2, Date: "2021-01-02", Description: "Rent", Category: "Rent", Amount: 1000},
  {id: 3, Date: "2021-01-03", Description: "Gas", Category: "Transportation", Amount: 50},
  {id: 4, Date: "2021-01-04", Description: "Dinner", Category: "Food", Amount: 50},
  {id: 5, Date: "2021-01-05", Description: "Coffee", Category: "Food", Amount: 5},
  {id: 6, Date: "2021-01-06", Description: "Netflix", Category: "Entertainment", Amount: 10},
  {id: 7, Date: "2021-01-07", Description: "Gym", Category: "Health", Amount: 20},
  {id: 8, Date: "2021-01-08", Description: "Car Insurance", Category: "Insurance", Amount: 100},
  {id: 9, Date: "2021-01-09", Description: "Car Payment", Category: "Transportation", Amount: 300},
  {id: 10, Date: "2021-01-10", Description: "Electricity", Category: "Utilities", Amount: 50},
  {id: 11, Date: "2021-01-11", Description: "Water", Category: "Utilities", Amount: 50},
  {id: 12, Date: "2021-01-12", Description: "Internet", Category: "Utilities", Amount: 50},
  {id: 13, Date: "2021-01-13", Description: "Phone", Category: "Utilities", Amount: 50},
  {id: 14, Date: "2021-01-14", Description: "Gas", Category: "Transportation", Amount: 50},
  {id: 15, Date: "2021-01-15", Description: "Dinner", Category: "Food", Amount: 50}
]

</script>

<template>
  <DataTable :headers="headers" :rows="rows" :paginate="true" :rows-per-page="10"/>
</template>

<style scoped>

</style>