import { Component } from "react";
import "./App.css";
import Header from "./components/Header/Header";
import ChatHistory from "./components/ChatHistory/ChatHistory";
import ChatInput from "./components/ChatInput/ChatInput";
import { connect, sendMsg } from "./api";

class App extends Component {
  constructor(props) {
    super(props);

    this.state = {
      chatHistory: [],
      username: ""
    };
  }

  componentDidMount() {
    const name = prompt("Enter your name");

    this.setState({ username: name || "Anonymous" }, () => {
      connect((msg) => {
        let data = JSON.parse(msg.data);

        // Fix the nested stringified JSON body from Go backend
        if (data.body && typeof data.body === "string" && data.body.trim().startsWith("{")) {
          try {
            const nestedData = JSON.parse(data.body);
            data = {
              type: nestedData.type || data.type,
              username: nestedData.username || data.username,
              body: nestedData.body,
              time: nestedData.time || data.time
            };
          } catch (e) {
            console.error("Failed to parse nested message body JSON:", e);
          }
        }

        this.setState((prev) => ({
          chatHistory: [...prev.chatHistory, data]
        }));
      });
    });
  }

  send = (text) => {
    if (!text.trim()) return;

    const payload = {
      type: "chat",
      username: this.state.username,
      body: text,
      time: new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit', second: '2-digit' })
    };

    sendMsg(JSON.stringify(payload));
  };

  render() {
    return (
      <div className="App">
        <Header />
        <ChatHistory chatHistory={this.state.chatHistory} />
        <ChatInput send={this.send} />
      </div>
    );
  }
}

export default App;