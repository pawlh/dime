<script>
	import {user} from '$lib/stores/auth';
</script>

<div class="container">
    <!-- <img /> -->
    <div>
        <h3>{$user.firstName} {$user.lastName}</h3>
    </div>
</div>

<style>
    .container {
        background-color: var(--surface--color--primary);
        display: flex;
        flex-direction: column;
    }
</style>
