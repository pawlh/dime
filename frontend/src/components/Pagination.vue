<script lang="ts" setup>
import {ref, computed, watch} from 'vue';
import {FontAwesomeIcon} from "@fortawesome/vue-fontawesome";

const props = defineProps({
  totalRows: {
    type: Number,
    required: true
  },
  rowsPerPage: {
    type: Number,
    default: 10
  }
})

const emit = defineEmits(['pageChanged'])

const currentPage = ref(1);

const totalPages = computed(() => Math.ceil(props.totalRows / props.rowsPerPage));

watch(currentPage, () => {
  emit('pageChanged', currentPage.value);
})

function goToPreviousPage() {
  if (currentPage.value > 1)
    currentPage.value--;
}

function goToNextPage() {
  if (currentPage.value < totalPages.value)
    currentPage.value++;
}
</script>

<template>
  <div class="pagination">
    <span class="arrow" @click="goToPreviousPage">
      <font-awesome-icon icon="fa-chevron-left"/>
    </span>
    <span class="page-status">{{ currentPage }} of {{ totalPages }}</span>
    <span class="arrow" @click="goToNextPage">
      <font-awesome-icon icon="fa-chevron-right"/>
    </span>
  </div>
</template>

<style scoped>

.pagination {
  padding: .4rem 0;
  display: flex;
  align-items: center;
  justify-content: center;
}

.arrow {
  cursor: pointer;
}

.page-status {
  margin: 0 .5rem;

  font-size: 1rem;
}
</style>