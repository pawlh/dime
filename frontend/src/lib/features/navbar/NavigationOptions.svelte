<script lang="ts">
    enum NavItem {
        Dashboard = 'dashboard',
        Budget = 'budget',
        Transactions = 'transactions'
    }

    let selectedItem: NavItem = NavItem.Dashboard;
</script>

<ul class="outer-list">
    <li class:selected={selectedItem === NavItem.Dashboard}>
        <a on:click={() => selectedItem = NavItem.Dashboard} href="/#" class="nav-link" role="tab">Dashboard</a>
    </li>
    <li class:selected={selectedItem === NavItem.Budget}>
        <a on:click={() => selectedItem = NavItem.Budget} href="/#" class="nav-link" role="tab">Budget</a>
    </li>
    <li class:selected={selectedItem === NavItem.Transactions}
        class="inner-list">
        <a on:click={() => selectedItem = NavItem.Transactions} href="/#" class="nav-link"
           role="tab">Transactions</a>
        <ul>
            <li><a href="/#">Account 1</a></li>
            <li><a href="/#">Account 2</a></li>
            <li><a href="/#">Account 3</a></li>
        </ul>
    </li>
</ul>

<style lang="scss">
  ul {
    list-style-type: none;
    padding: 0;
  }

  .nav-link {
    padding-left: 30px;
    border-top-right-radius: 25px;
    border-bottom-right-radius: 25px;

    &:hover {
      background-color: var(--surface--color--secondary);
    }
  }

  li {
    padding: 5px 0;
  }

  .selected > .nav-link {
    border-left: 2px solid var(--border--color);

    background-color: var(--surface--color--secondary);
  }
</style>