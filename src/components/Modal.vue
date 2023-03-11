<script setup>
import {ref} from 'vue';

const modal = ref('modal')//document.getElementById('modal');
const closeButton = document.getElementsByClassName('close')[0];

function openModal() {
  modal.value.style.display = 'block';
}

function closeModal() {
  modal.value.style.display = 'none';
}

// window.onclick = function(event) {
//   if (event.target === modal) {
//     closeModal();
//   }
// }

</script>

<template>
  <div id="modal" class="modal" ref="modal">
    <div class="modal-content">
      <span class="close" @click="closeModal">&times;</span>
      <p>This is the modal content!</p>
    </div>
  </div>

</template>

<style scoped>
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

.modal-content {
  background-color: var(--color-background-soft);
  margin: 15% auto;
  padding: 20px;
  /*border: 1px solid #888;*/
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
  color: black;
  text-decoration: none;
  cursor: pointer;
}

</style>