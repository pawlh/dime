<script setup>

import Pagination from "@/components/Pagination.vue"
import {ref, computed} from 'vue'

const props = defineProps({
  headers: Array,
  rows: Array,
  paginate: Boolean,
  rowsPerPage: Number
})

const sortedBy = ref("");

const sortedRows = computed(() => {
  if (sortedBy.value === "") {
    return props.rows;
  }

  return props.rows.sort((a, b) => {
    if (a[sortedBy.value] < b[sortedBy.value]) {
      console.log("a is less than b")
      return -1;
    }
    if (a[sortedBy.value] > b[sortedBy.value]) {
      console.log("a is greater than b")
      return 1;
    }
    return 0;
  });
})

const currentPage = ref(1);

const paginatedRows = computed(() => {
  const startIndex = (currentPage.value - 1) * props.rowsPerPage;
  const endIndex = startIndex + props.rowsPerPage;
  return sortedRows.value.slice(startIndex, endIndex);
});

</script>

<template>
  <div class="datatable-container">
    <table class="datatable">
      <thead>
      <tr>
        <template v-for="header in headers" :key="header.label">
          <th v-if="!header.hidden && header.sortable"
              @click="() => sortedBy = header.label">
            {{ header.label }}
          </th>
          <th v-else-if="!header.hidden">
            {{ header.label }}
          </th>
        </template>
      </tr>
      </thead>
      <tbody>
      <tr v-for="row in paginatedRows" :key="row.id">
        <template v-for="header in headers" :key="header.label">
          <template v-if="!header.hidden">
            <td v-if="header.formatter">
              {{ header.formatter(row[header.label]) }}
            </td>
            <td v-else>
              {{ row[header.label] }}
            </td>
          </template>
        </template>
      </tr>
      </tbody>
    </table>
    <pagination :total-rows="rows.length"
                :rows-per-page="rowsPerPage"
                v-if="paginate"
                @page-changed="currentPage = $event"
    />
  </div>
</template>

<style scoped>
.datatable-container {
  background-color: #252525;
  border-radius: 8px;
  padding: 20px;
}

.datatable {
  width: 100%;
  border-collapse: collapse;
}

.datatable th {
  background-color: #333333;
  color: #ffffff;
  font-weight: 500;
  padding: 12px;
  text-align: left;
}

.datatable td {
  background-color: #444444;
  color: #ffffff;
  padding: 12px;
}

.datatable tbody tr:nth-of-type(odd) td {
  background-color: #555555;
}

.datatable tbody tr:hover td {
  background-color: #666666;
}
</style>