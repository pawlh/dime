<script lang="ts" setup>
const props = defineProps({
  iconType: {type: String, required: true}
})
</script>

<template>
  <div class="space"></div>
  <font-awesome-icon class="icon" :icon="['fa-solid', props.iconType]"/>
  <span><slot></slot> </span>
</template>

<style lang="less" scoped>
.icon {
  font-size: 1.5rem;
  width: 2rem;
}

span {
  margin-left: 0.5rem;
}
</style>
