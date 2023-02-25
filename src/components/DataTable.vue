<script lang="ts" setup>

import Pagination from "@/components/Pagination.vue"
import {ref, computed} from 'vue'

export interface ColumnDefinition {
  label: string,
  sortable: boolean,
  prefix?: string,
  hidden?: boolean
}

export interface RowData {
  [key: string]: any
}

const props = defineProps<{
  headers: ColumnDefinition[],
  rows: RowData[],
  paginate: boolean,
  rowsPerPage: number
}>()

const currentPage = ref(1);

const paginatedRows = computed(() => {
  const startIndex = (currentPage.value - 1) * props.rowsPerPage;
  const endIndex = startIndex + props.rowsPerPage;
  return props.rows.slice(startIndex, endIndex);
});

</script>

<template>
  <div class="datatable-container">
    <table class="datatable">
      <thead>
      <tr>
        <template v-for="header in headers" :key="header.label">
          <th v-if="!header.hidden">
            {{ header.label }}
          </th>
        </template>
      </tr>
      </thead>
      <tbody>
      <tr v-for="row in paginatedRows" :key="row.id">
        <template v-for="header in headers" :key="header.label">
          <td v-if="!header.hidden">
            {{ header.prefix }}{{ row[header.label] }}
          </td>
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