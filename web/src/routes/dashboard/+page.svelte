<script lang="ts">
    import {websocket} from "../authWithCode/stores";
    import type {CommonGuild} from "$lib/guild";
    import CommonGuildSelect from "$lib/commomGuildSelect.svelte"

    let ws: any|WebSocket = null;
    let commonGuilds: CommonGuild[] = [];
    let selectedGuild: CommonGuild | null = null;
    let displayMessage: string|null = null;

    const setDisplayMessage = (msg: string) => {
        displayMessage = msg;
        setTimeout(() => displayMessage = null, 1000);
    }

    const connect = () => {
        if (ws !== null) {
            ws.send(JSON.stringify({
                action: "CONNECT"
            }));
        }
    }

    websocket.subscribe((nws) => {
        ws = nws;
        if (ws !== null) {
            ws.onmessage = (msg: MessageEvent) => {
                const json = JSON.parse(msg.data);
                if (json.action === "GET_COMMON_GUILDS") {
                    commonGuilds = json.content.guilds;
                }
                if (json.action === "CONNECT") {
                    setDisplayMessage(json.message);
                }
            }
            ws.send(JSON.stringify({
                action: "GET_COMMON_GUILDS"
            }));
        }
    });


</script>

<div class="fixed-grid has-10-cols">
    <div class="grid">
        <div class="cell is-col-span-2">
            <CommonGuildSelect commonGuilds={commonGuilds} bind:selectedGuild={selectedGuild} />
        </div>
        <div class="cell is-col-span-7"></div>
        <div class="cell is-col-end-1">
            {#if selectedGuild}
                <button class="button is-primary" on:click={connect}>Connect</button>
            {/if}
        </div>
    </div>
</div>