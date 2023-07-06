<script>
    import {theme} from "$lib/stores/theme";
</script>

<div class="switch-container">
    <input type="checkbox" id="theme-switch" on:change={theme.toggle} checked={$theme === 'dark'}>
    <label for="theme-switch" class="switch"></label>
</div>

<style>

    #theme-switch {
        background-color: var(--background--color);
    }

    .switch-container {
        display: flex;
        justify-content: center;
        align-items: center;
    }

    .switch {
        display: inline-block;
        width: 60px;
        height: 34px;
        position: relative;
        background-color: #ccc;
        border-radius: 17px;
        transition: background-color 0.3s ease;
    }

    .switch:before,
    .switch:after {
        content: "";
        position: absolute;
        top: 2px;
        width: 30px;
        height: 30px;
        border-radius: 50%;
        transition: transform 0.3s ease;
    }

    .switch:before {
        left: 2px;
        background-color: #fff;
    }

    .switch:after {
        right: 2px;
        background-color: #000;
    }

    input[type="checkbox"] {
        visibility: hidden;
    }

    input[type="checkbox"]:checked + .switch:before {
        transform: translateX(26px);
    }

    input[type="checkbox"]:checked + .switch:after {
        transform: translateX(-26px);
    }

    input[type="checkbox"]:checked + .switch {
        background-color: #6e6e6e;
    }

</style>