<script setup>
</script>

<template>
  <div class="tools">
    <div class="options">
      <router-link to="/tools/import">Import Transactions</router-link>
      <router-link to="/tools/accounts">Manage Accounts</router-link>
    </div>
    <router-view></router-view>
  </div>
</template>

<style lang="less" scoped>
.options {
  display: flex;
  flex-direction: column;
  align-items: center;
}

</style>