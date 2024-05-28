<script lang="ts">
    import { page } from '$app/stores';
    import {onMount} from "svelte";
    onMount(() => {
        let hasCode = $page.url.searchParams.has('code');
        if (hasCode) {
            let code = $page.url.searchParams.get("code");
            let websocketUrl = (process.env.NODE_ENV === "production" ? "/ws" : "http://localhost:3000/ws") + "?code=" + code;
            let ws = new WebSocket(websocketUrl);
            ws.onmessage = (msg) => console.log(msg.data);
        }
    })
</script>

<h1>Loading...</h1>
