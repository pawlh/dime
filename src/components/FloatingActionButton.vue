<script setup>
</script>

<template>
  <div class="floating-circle">
    <font-awesome-icon class="icon" icon="fa-upload"/>
  </div>
</template>

<style lang="less" scoped>
.floating-circle {
  position: fixed;
  bottom: 20px;
  right: 20px;
  width: 60px;
  height: 60px;
  border-radius: 50%;
  background-color: var(--color-accent);
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.3);
  display: flex;
  justify-content: center;
  align-items: center;
  z-index: 9999;

  cursor: pointer;

  &:hover {
    // darken the color var(--color-accent)
    filter: brightness(90%);
  }

  transition: filter 0.2s ease-in-out;
}

.icon {
  height: 25px;
  /*color: #FF5722;*/
}
</style>