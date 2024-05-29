<script lang="ts">
    import {websocket} from "../authWithCode/stores";
    import type {CommonGuild} from "$lib/guild";
    import CommonGuildSelect from "$lib/commomGuildSelect.svelte"

    let ws: WebSocket|null = new WebSocket("");
    let commonGuilds: CommonGuild[] = [];
    let selectedGuild: CommonGuild|null = null;
    websocket.subscribe((nws) => ws = nws);
    ws.send(JSON.stringify({
        action: "GET_COMMON_GUILDS"
    }));

    ws.onmessage = (msg) => {
        const json = JSON.parse(msg.data);
        if (json.action === "GET_COMMON_GUILDS") {
            commonGuilds = json.content.guilds;
        }
    }

</script>

<div class="fixed-grid has-10-cols">
    <div class="grid">
        <div class="cell is-col-span-2">
            <CommonGuildSelect commonGuilds={commonGuilds} bind:selectedGuild={selectedGuild} />
        </div>
    </div>
</div>