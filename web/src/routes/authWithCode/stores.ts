import {type Writable, writable} from "svelte/store";

export interface DiscordUser {
    username: string;
    global_name: string;
    id: string;
    avatar: string;
}

export const websocket: Writable<WebSocket|null> = writable<WebSocket|null>(null);
export const discordUser: Writable<DiscordUser|null> = writable<DiscordUser|null>(null);