var socket = new WebSocket("ws://localhost:8080/ws");

let connect = (cb) => {
  console.log("connecting");

  socket.onopen = () => {
    console.log("connected");
  };

  socket.onmessage = (msg) => {
    cb(msg);
  };

  socket.onclose = () => {
    console.log("closed");
  };

  socket.onerror = (err) => {
    console.log("error", err);
  };
};

let sendMsg = (msg) => {
  socket.send(msg);
};

export { connect, sendMsg };