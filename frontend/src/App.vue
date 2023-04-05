<script setup lang="ts">
import {RouterView} from 'vue-router'
import {setupTransactionWS} from "@/utils/transactions";

// Default application to dark mode
document.documentElement.setAttribute("theme", "dark");


//connect to websocket at /api/transactions
setupTransactionWS()


</script>

<template>
  <router-view></router-view>
</template>

<style scoped>
</style>
