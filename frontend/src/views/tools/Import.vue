<script setup>
import {FontAwesomeIcon} from "@fortawesome/vue-fontawesome";
import { Toaster, toast} from "vue-sonner";
import {SERVER_URL} from "@/store/app";

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

// temporary, this will ultimately be replaced with a configurable modal
const sampleMapping = {
    "column_mapping": {
        "id": "id",
        "date": "date",
        "description": "description",
        "amount": "amount",
        "category": "category",
        "account": "account"
    },
    "date_format": "2006-01-02"
}
async function handleFiles(files) {
    for (const file of files) {
        const formData = new FormData();
        formData.append("file", file);
        formData.append("meta", JSON.stringify(sampleMapping));

        const res = await fetch(SERVER_URL + "/api/upload", {
            method: "POST",
            credentials: "include",
            body: formData
        })


        if (res.ok) {
            toast.success("Transaction import successful")
        } else {
            toast.error("Transaction import failed", {
                description: `There was an error importing your transactions. Check the console for more details.`
            })

            const data = await res.json()
            console.log(`Error uploading CSV: ${data.error}`)
        }
    }
}

</script>

<template>
    <Toaster/>
    <div class="import">
        <h1>Import transactions from your bank</h1>
        <span class="demo-csv"><a href="/demo-transactions.csv">Download a demo csv</a></span>
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

.demo-csv {
    display: block;
    /*margin: 0 auto;*/
    text-align: center;
}
.demo-csv > a {
    color: var(--color-text);
    text-decoration: underline;
}

</style>