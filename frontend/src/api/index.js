var socket = new WebSocket("ws://localhost:8080/ws")

let connect = callback => {
    console.log("Attempting onnection ...");

    socket.onopen = () => {
        console.log("Sucessfully connected!")
    };

    socket.onmessage = msg => {
        callback(msg)
    };

    socket.onclose = event => {
        console.log("Socket closed connection: ", event)
    };

    socket.onerror = error => {
        console.log("Socket Error: ", error)
    };

};

let sendMsg = msg => {
    console.log("sending msg: ", msg);
    socket.send(msg);
}

export { connect, sendMsg }