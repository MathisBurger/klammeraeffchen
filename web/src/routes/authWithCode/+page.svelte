<script lang="ts">
    import { page } from '$app/stores';
    import {onMount} from "svelte";
    import {discordUser, websocket} from "./stores";
    import {goto} from "$app/navigation";
    import {setCookie} from "$lib/cookie";

    onMount(() => {
        let hasCode = $page.url.searchParams.has('code');
        if (hasCode) {
            let code = $page.url.searchParams.get("code");
            let websocketUrl = (process.env.NODE_ENV === "production" ? "/ws" : "http://localhost:3000/ws") + "?code=" + code;
            let ws = new WebSocket(websocketUrl);
            ws.onmessage = (msg) => {
                const json = JSON.parse(msg.data);
                if (json.action === "AUTH_REFRESH_TOKEN") {
                    setCookie("session", json.content.refresh_token, json.content.expires_in);
                    return;
                }
                if (json.action === "AUTH_USER_ID") {
                    websocket.set(ws);
                    discordUser.set(json.content)
                    goto("/dashboard");
                    return;
                }
                goto("/")
            }
        }
    })
</script>

<h1>Loading...</h1>
