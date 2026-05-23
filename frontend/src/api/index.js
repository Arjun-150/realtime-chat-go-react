let socket = null;

export const connect = (cb) => {
  console.log("connecting");

  socket = new WebSocket("ws://localhost:8080/ws");

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

export const sendMsg = (msg) => {
  if (!socket || socket.readyState !== WebSocket.OPEN) {
    console.log("socket not ready");
    return;
  }
  socket.send(msg);
};