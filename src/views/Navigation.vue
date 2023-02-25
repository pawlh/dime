<script lang="ts" setup>
import {useRoute} from "vue-router";
import NavItem from "@/components/NavItem.vue";

const route = useRoute();

</script>

<template>
  <ul>
    <li>
      <NavItem
          :selected="route.name === 'home'"
          icon-type="fa-piggy-bank"
          destination="home">Accounts
      </NavItem>
    </li>
    <li>
      <NavItem
          :selected="route.name === 'budgets'"
          icon-type="fa-calculator"
          destination="budgets">Budgets
      </NavItem>
    </li>
    <li>
      <NavItem
          :selected="route.name === 'reports'"
          icon-type="fa-chart-line"
          destination="reports">Reports
      </NavItem>
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