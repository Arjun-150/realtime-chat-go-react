let socket = null;

export const connect = (cb, username) => {
  console.log("connecting");

  // 🚀 DYNAMIC IP DETECTION
  // window.location.hostname will be 'localhost' or '192.168.x.x' automatically
  const host = window.location.hostname;
  socket = new WebSocket(`ws://${host}:8080/ws?username=${username}`);

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