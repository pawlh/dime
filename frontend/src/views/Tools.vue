<script setup>
import {useRoute} from "vue-router";

const route = useRoute();
</script>

<template>
  <div class="tools">
    <div class="options">
      <div class="nav-option" :class="{selected: route.name === 'import'}">
        <router-link :to="{name: 'import'}">Import Transactions</router-link>
      </div>
      <div class="nav-option" :class="{selected: route.name === 'accounts'}">
        <router-link :to="{name: 'accounts'}">Manage Accounts</router-link>
      </div>
    </div>
    <div class="content">
      <router-view></router-view>
    </div>
  </div>
</template>

<style lang="less" scoped>
@import "@/assets/mixins.less";

.tools {
  background-color: var(--color-background-soft);
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.3);
}

.options {
  background-color: var(--color-background-lightest);

  //align-items: center;


  @media @md-desktop {
    width: 100%;
    min-height: 50px;

    display: flex;

    flex-direction: row;
    justify-content: space-evenly;
    column-gap: 5px;

  }
}

.nav-option {
  width: 100%;
  //height: 100%;
  text-align: center;


  display: flex;
  align-items: center;

  & > a {
    color: var(--color-background-soft);
    flex-grow: 1;

    font-weight: bold;

    &:hover {
      background-color: unset;
    }
  }

  border-bottom: 5px solid transparent;

  transition: border-bottom-color 0.2s ease-in;
}

//.nav-option > a {
//  color: var(--color-background-soft);
//}

.selected {
  border-bottom: 5px solid var(--color-accent);
}

.content {
  padding: 1rem;
}

a:link, a:visited, a:hover, a:active {
  text-decoration: none;
}

</style>