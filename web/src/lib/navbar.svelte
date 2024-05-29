<script lang="ts">

    import {type DiscordUser, discordUser, websocket} from "../routes/authWithCode/stores";
    import {tryGetWebSocketSession} from "$lib/session";
    import {onMount} from "svelte";
    import {page} from "$app/stores";
    import {goto} from "$app/navigation";

    let ws: WebSocket|null = null;
    let dcUser: DiscordUser|null = null;
    let loginBaseUrl = process.env.NODE_ENV === "production" ? "/login" : "http://localhost:3000/login"
    onMount(() => {
        websocket.subscribe((newWs) => {
            if (newWs === null) {
                tryGetWebSocketSession();
            } else {
                ws = newWs
            }
        });
        setTimeout(() => {
            if (ws === null && $page.url.pathname.indexOf("/dashboard") > -1) {
                goto("/")
            }
        }, 1000);
    })
    discordUser.subscribe((newUser) => {
        console.log(newUser);
        dcUser = newUser;
    });
    let userAvatar: string|null = null;
    if (dcUser !== null) {
        userAvatar = "https://cdn.discordapp.com/avatars/" + dcUser?.id + "/" + dcUser?.avatar;
    }
    console.log(userAvatar)


</script>

<nav class="navbar navbarShadow" role="navigation" aria-label="main navigation">
    <div class="navbar-brand">
        <a class="navbar-item" href="#">
            <img src="/logo.png" alt="logo" />
            <h1 class="is-size-3 ml-2 has-text-weight-bold">Klammeraeffchen</h1>
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
                    <p class="has-text-weight-bold mr-1">{dcUser.global_name}</p>
                    <!--<img src={userAvatar} alt="avatar" class="profilePic" />-->
                {/if}
            </div>
        </div>
    </div>
</nav>

<style>
    .profilePic {
        height: 28px;
        width: 28px;
        border-radius: 50%;
    }
    .navbarShadow {
        box-shadow: 5px 5px 5px #00000030;
    }
</style>