<script setup>
import {FontAwesomeIcon} from "@fortawesome/vue-fontawesome";

function handleDrop(event) {
  event.preventDefault();
  handleFiles(event.dataTransfer.files);
}

function handleClick() {
  const input = document.createElement('input');
  input.type = 'file';
  input.accept = '.csv';
  input.onchange = (event) => {
    console.log(event)
    handleFiles(event.target.files)
  };
  input.click();
}

function handleFiles(files) {
  for (const file of files) {
    alert('uploading ' + file.name);
  }
}

</script>

<template>
  <div class="import">
    <h1>Import transactions from your bank</h1>
    <div class="dropzone-container">
      <div class="dropzone" @dragover.prevent @drop="handleDrop" @click="handleClick">
        <font-awesome-icon class="icon" icon="fa-file-import"/>
        <div class="dropzone-text">Drag and drop files here or click to upload <br/> (.csv only)</div>
      </div>
    </div>
  </div>
</template>

<style scoped>
h1 {
  text-align: center;
  font-weight: bold;
}

.dropzone-container {
  margin-top: 3rem;

  display: flex;
  justify-content: center;
  align-items: center;
  height: 100%;
}

.dropzone {
  width: 300px;
  height: 300px;
  border: 2px dashed gray;
  border-radius: 10px;
  background-color: var(--color-background-mute);

  display: flex;
  flex-direction: column;
  row-gap: 1rem;

  justify-content: center;
  align-items: center;

  cursor: pointer;
}

.dropzone-text {
  font-size: 1.2rem;
  text-align: center;
  color: var(--color-text);
}

.icon {
  font-size: 3rem;
  color: var(--color-text);
}

</style>