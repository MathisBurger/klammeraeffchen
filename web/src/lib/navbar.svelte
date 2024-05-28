<script lang="ts">

    import {type DiscordUser, discordUser, websocket} from "../routes/authWithCode/stores";

    let ws = null;
    let dcUser: DiscordUser|null = null;
    let loginBaseUrl = process.env.NODE_ENV === "production" ? "/login" : "http://localhost:3000/login"
    websocket.subscribe((newWs) => ws = newWs);
    discordUser.subscribe((newUser) => dcUser = newUser);
</script>

<nav class="navbar" role="navigation" aria-label="main navigation">
    <div class="navbar-brand">
        <a class="navbar-item" href="#">
            <img src="/logo.png" alt="logo" />
            <h1 class="is-size-3 ml-2">Klammeraeffchen</h1>
        </a>
    </div>

    <div id="navbarBasicExample" class="navbar-menu">
        <div class="navbar-end">
            <div class="navbar-item">
                {#if dcUser === null}
                    <div class="buttons">
                        <a class="button is-link" href="{loginBaseUrl}">
                            <strong>Login</strong>
                        </a>
                    </div>
                {:else }
                    <p>{dcUser.global_name}</p>
                {/if}
            </div>
        </div>
    </div>
</nav>