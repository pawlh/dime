import {useStateStore} from "@/store/state";

export const setupTransactionWS = () => {
    const ws = new WebSocket("ws://localhost:1323/api/transactions");
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