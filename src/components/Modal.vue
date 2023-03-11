<script setup>
import {ref} from 'vue';

defineProps({
  visible: Boolean
})

const modal = ref('modal')

</script>

<template>
  <div id="modal"
       class="modal"
       :class="{hidden: !visible}"
       ref="modal"
       @click="$emit('close-modal')">
    <div class="modal-content" @click.stop>
      <span class="close" @click="$emit('close-modal')">&times;</span>
      <p>This is the modal content!</p>
    </div>
  </div>

</template>

<style lang="less" scoped>
.modal {
  display: block;
  position: fixed;
  z-index: 1;
  left: 0;
  top: 0;
  width: 100%;
  height: 100%;
  overflow: auto;
  background-color: rgba(0, 0, 0, 0.4);
}

.hidden {
  display: none;
}

.modal-content {
  background-color: var(--color-background-soft);
  margin: 15% auto;
  padding: 20px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.3);
  width: 80%;
}

.close {
  color: #aaa;
  float: right;
  font-size: 28px;
  font-weight: bold;
}

.close:hover,
.close:focus {
  //color: darken(#aaa);
  text-decoration: none;
  cursor: pointer;
}

</style>