<script setup>
import {FontAwesomeIcon} from "@fortawesome/vue-fontawesome";
import {useStateStore} from "@/store/state";

const stateStore = useStateStore();

const logout = () => {
  stateStore.loggedInUser.firstName = "";
  stateStore.loggedInUser.lastName = "";
  stateStore.loggedIn = false;
};

</script>

<template>
  <div>
    <font-awesome-icon class="profileImage" icon="fa-user"/>
    <div class="info">
      <h2>{{ stateStore.loggedInUser.firstName }} {{ stateStore.loggedInUser.lastName }}</h2>
      <ul>
        <!--        <li>-->
        <!--          <a href="#">Settings</a>-->
        <!--        </li>-->
        <li>
          <a @click="logout">Logout</a>
        </li>
      </ul>
    </div>
  </div>
</template>


<style scoped>
div {
  width: 100%;
  height: 6rem;

  display: flex;
  align-items: center;
  gap: 1rem;
}

.info {
  display: flex;
  flex-direction: column;
  align-items: start;
  gap: .25rem;
}

ul {
  list-style-type: none;
  padding: 0 0 0 1rem;
}

a {
  text-decoration: none;
  color: var(--color-text);
  cursor: pointer;
}

.profileImage {
  height: 80%;
}
</style>