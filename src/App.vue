<script setup lang="ts">
import {RouterView} from 'vue-router'
import Navigation from "@/views/Navigation.vue";
import ProfilePreview from "@/components/ProfilePreview.vue";

// Default application to dark mode
document.documentElement.setAttribute("theme", "dark");

</script>

<template>
  <div id="app-container">
    <nav>
      <h1 class="app-title">Unero</h1>
      <hr/>
      <Navigation class="navigation"/>
      <hr/>
      <ProfilePreview class="profile"/>
    </nav>
    <main>
      <router-view></router-view>
    </main>
  </div>
</template>

<style lang="less" scoped>
@import "@/assets/mixins.less";

#app-container {
  height: 100vh;
  width: 100%;

  display: flex;
  flex-direction: column;

  @media @md-desktop {
    flex-direction: row;
  }
}

.app-title {
  font-size: 2rem;
  font-weight: 700;

  text-align: center;

  @media @md-desktop {
    text-align: start;
  }
}

nav {
  display: flex;
  flex-direction: column;

  background-color: var(--color-background-mute);
  padding: 1rem;
  flex-grow: 1;

  box-shadow: 0 5px 10px rgb(0 0 0 / 0.2);

  @media @md-desktop {
    max-width: 20rem;

    box-shadow: 5px 0 10px rgb(0 0 0 / 0.2);
  }

  &.navigation-open {
    transform: translateX(-100%);
    flex-shrink: 0;

  }

  transition: transform 0.5s ease-in-out, flex-grow 0.5s ease-in-out;

  .navigation {
    flex-grow: 1;
  }

  .profile {
    align-self: flex-end;
  }
}

main {
  padding: 1rem;
  flex-grow: 4;
}

hr {
  margin-bottom: 1rem;
}
</style>
