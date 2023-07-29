<script>
    import NavBar from "$lib/features/navbar/NavBar.svelte";
    import ThemeChooser from "$lib/features/navbar/ThemeChooser.svelte";
</script>

<div id="container">
    <div class="nav-bar">
        <NavBar/>
    </div>
    <div class="main">
        <div class="status-bar">
            <ThemeChooser/>
        </div>
        <slot/>
    </div>
</div>

<style lang="scss">
  #container {
    display: flex;
    flex-direction: column;
    height: 100vh;
  }

  .nav-bar {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    z-index: 100;
  }

  @media (min-width: 768px) {
    #container {
      flex-direction: row;
    }

    .nav-bar {
      position: static;
      width: 250px;

      height: 100vh;
    }
  }
</style>
