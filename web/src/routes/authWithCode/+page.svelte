<script lang="ts">
    import { page } from '$app/stores';
    import {onMount} from "svelte";
    import {discordUser, websocket} from "./stores";
    import {goto} from "$app/navigation";
    onMount(() => {
        let hasCode = $page.url.searchParams.has('code');
        if (hasCode) {
            let code = $page.url.searchParams.get("code");
            let websocketUrl = (process.env.NODE_ENV === "production" ? "/ws" : "http://localhost:3000/ws") + "?code=" + code;
            let ws = new WebSocket(websocketUrl);
            ws.onmessage = (msg: MessageEvent) => {
                const json = JSON.parse(msg.data);
                if (json.message.indexOf("User fetched successfully") > -1) {
                    websocket.set(ws);
                    discordUser.set(json.content)
                    goto("/dashboard");
                }
            }
        }
    })
</script>

<h1>Loading...</h1>
