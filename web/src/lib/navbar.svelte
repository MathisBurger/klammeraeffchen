<script lang="ts">

    import {type DiscordUser, discordUser, websocket} from "../routes/authWithCode/stores";
    import {page} from "$app/stores";
    import {goto} from "$app/navigation";
    import {browser} from "$app/environment";
    import {getCookie, setCookie} from "$lib/cookie";

    let ws: WebSocket|null;
    let dcUser: DiscordUser|null = null;
    let loginBaseUrl = process.env.NODE_ENV === "production" ? "/login" : "http://localhost:3000/login"

    const onClose = () => {
        console.log("close");
        setCookie("session", "", "");
        goto("/");
    }

    const onMessage = (msg: MessageEvent) => {
        const json = JSON.parse(msg.data);
        if (json.action === "AUTH_REFRESH_TOKEN") {
            setCookie("session", json.content.refresh_token, json.content.expires_in);
            return;
        }
        if (json.action === "AUTH_USER_ID") {
            discordUser.set(json.content)
            websocket.set(ws);
            return;
        }
    }

    if (browser) {
        if ($page.url.pathname.indexOf("/authWithCode") === -1) {
            websocket.subscribe((newWs) => {
                if (newWs !== null) {
                    ws = newWs
                    ws.onmessage = onMessage;
                    ws.onclose = () => onClose;

                } else if (newWs === null && getCookie("session") !== ""){
                    let websocketUrl = (process.env.NODE_ENV === "production" ? "/ws" : "http://localhost:3000/ws") + "?refreshToken=" + getCookie("session");
                    ws = new WebSocket(websocketUrl);
                    ws.onmessage = onMessage;
                    ws.onclose = () => onClose;
                } else {
                    goto("/");
                }
            })
        }
    }
    discordUser.subscribe((newUser) => {
        dcUser = newUser;
    });
    let userAvatar: string|null = null;
    if (dcUser !== null) {
        userAvatar = "https://cdn.discordapp.com/avatars/" + dcUser?.id + "/" + dcUser?.avatar;
    }


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