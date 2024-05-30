import {getCookie, setCookie} from "$lib/cookie";
import {discordUser, websocket} from "../routes/authWithCode/stores";


export function tryGetWebSocketSession() {
    const cookie = getCookie("session");
    if (cookie === "") {
        return;
    }
    let websocketUrl = (process.env.NODE_ENV === "production" ? "/ws" : "http://localhost:3000/ws") + "?refreshToken=" + cookie;
    const ws = new WebSocket(websocketUrl);
    console.log("open socket")
    ws.onmessage = (msg: MessageEvent) => {
        console.log(msg.data);
        const json = JSON.parse(msg.data);
        if (json.action === "AUTH_REFRESH_TOKEN") {
            setCookie("session", json.content.refresh_token, json.content.expires_in);
            return;
        }
        if (json.action === "AUTH_USER_ID") {
            discordUser.update(() => json.content)
            return;
        }
    }
}

export function initializeWebsocketSession() {

}