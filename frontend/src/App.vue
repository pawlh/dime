<script setup lang="ts">
import {RouterView} from 'vue-router'

// Default application to dark mode
document.documentElement.setAttribute("theme", "dark");

</script>

<template>
  <router-view></router-view>
</template>

<style scoped>
</style>
