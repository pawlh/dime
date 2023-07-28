<script>
    import UserInfo from './UserInfo.svelte';
    import NavigationOptions from "$lib/features/navbar/NavigationOptions.svelte";
</script>

<div class="container">
    <UserInfo/>
    <NavigationOptions/>
</div>

<style lang="scss">
  .container {
    height: 100%;
    box-sizing: border-box;
    padding: 15px 15px 15px 0;

    background-color: var(--surface--color--primary);
  }

  .container {
    display: flex;
    flex-direction: row;
  }

  @media (min-width: 768px) {
    .container {
      flex-direction: column;
    }
  }
</style>
