<script setup>
</script>

<template>
    <div class="container">
        <h1>About</h1>
        <p>Dime is a self-hostable web application that allows users to track their spending. It is built with Go in the
            backend and VueJS in the frontend. Users can upload .csv files from their bank and view their spending in a
            variety of ways. Budgeting goals can be set and tracked. Users can also create and manage multiple
            accounts.</p>
        <p>This is in early development and is some time away from having many of the basic features implemented.</p>
        <div>
            <p>Development work can be found on <a href="https://github.com/pawlh/dime">GitHub</a></p>
            <img src="https://gh-stats.pawlh.dev/api/pin?username=pawlh&repo=dime&theme=transparent"
                 alt="dime stats on github">
        </div>
    </div>
</template>

<style scoped>
.container {
    max-width: 800px;
    margin: 0 auto;
    padding: 20px;
}

h1 {
    text-align: center;
    font-size: 3em;
}

p {
    margin-bottom: 20px;
}

a {
    color: var(--color-text);
    text-decoration: none;
    font-weight: bold;
}
</style>