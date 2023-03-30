<script setup lang="ts">
import {RouterView} from 'vue-router'
import Navigation from "@/views/Navigation.vue";
import ProfilePreview from "@/components/ProfilePreview.vue";
import FloatingActionButton from "@/components/FloatingActionButton.vue";

import {useRouter} from "vue-router";
import {useStateStore} from "@/store/state";
import Login from "@/views/Login.vue";

// Default application to dark mode
document.documentElement.setAttribute("theme", "dark");

</script>

<template>
  <router-view></router-view>
</template>

<style scoped>
</style>
