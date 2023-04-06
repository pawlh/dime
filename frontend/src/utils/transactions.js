import {useStateStore} from "@/store/state";
import {SERVER_URL} from "@/store/app";

export const setupTransactionWS = () => {
    const ws = new WebSocket(`ws://${SERVER_URL}/api/transactions`);
    ws.onopen = () => {
        console.log("Connected to websocket");
    };
    ws.onmessage = (event) => {
        console.log("Received message from websocket");

        // if (event.data === "ping") {
        //     ws.send("pong")
        //     return
        // }

        const data = JSON.parse(event.data);
        if ('transactions' in data) {
            const stateStore = useStateStore();

            stateStore.transactions = data.transactions;
        }
    };
    ws.onclose = () => {
        console.log("Disconnected from websocket");
    };
    ws.onerror = (event) => {
        console.log("Error with websocket");
        console.log(event);
    };
}