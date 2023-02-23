<script lang="ts" setup>

import NavItem from "@/components/NavItem.vue";
</script>

<template>
  <ul>
    <li>
      <NavItem icon-type="fa-piggy-bank">Accounts</NavItem>
    </li>
    <li>
      <NavItem icon-type="fa-calculator">Budgets</NavItem>
    </li>
    <li>
      <NavItem icon-type="fa-chart-line">Reports</NavItem>
    </li>
  </ul>
</template>

<style lang="less" scoped>
ul {
  list-style: none;
  padding: 0;
}

li {
  display: flex;
  align-items: center;

  &:not(:last-child) {
    margin-bottom: 1rem;
  }
}
</style>