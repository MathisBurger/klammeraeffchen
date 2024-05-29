<script lang="ts">
    import {websocket} from "../authWithCode/stores";

    let ws: WebSocket|null = new WebSocket("");
    websocket.subscribe((nws) => ws = nws);

    let guilds = [];
    ws.send(JSON.stringify({
        action: "GET_COMMON_GUILDS"
    }));

    ws.onmessage = (msg) => {
        console.log(msg.data);
    }

</script>

<h1>Dashboard</h1>