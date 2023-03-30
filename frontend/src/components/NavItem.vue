<script lang="ts" setup>
import {useRouter} from "vue-router";

defineProps({
  iconType: {type: String, required: true},
  destination: {type: String, required: true},
  selected: {type: Boolean, default: false}
})

const router = useRouter();


</script>

<template>
  <div class="container"
       :class="{selected}"
       @click="() => {router.push({name: destination})}">
    <div class="space">&nbsp;</div>
    <font-awesome-icon class="icon" :icon="['fa-solid', iconType]"/>
    <span><slot></slot></span>
  </div>
</template>

<style lang="less" scoped>
.container {
  display: flex;
  flex-direction: row;

  cursor: pointer;

  &:hover {
    color: var(--color-accent);
  }
}

.space {
  width: .35rem;
  margin-right: 1rem;
}

.selected {
  color: var(--color-accent);

  > .space {
    background-color: var(--color-accent);
  }
}

.icon {
  font-size: 1.5rem;
  width: 2rem;
}

span {
  margin-left: 0.5rem;
}
</style>
