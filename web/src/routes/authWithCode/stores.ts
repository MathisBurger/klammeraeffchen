import {writable} from "svelte/store";


export const websocket = writable(new WebSocket(""));