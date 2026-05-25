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
    const finalizedName = name || "Anonymous";

    this.setState({ username: finalizedName }, () => {
      connect((msg) => {
        console.log("RAW MESSAGE FROM GO:", msg.data);
        let data = JSON.parse(msg.data);

        // Fix nested stringified JSON body if present
        if (data.body && typeof data.body === "string" && data.body.trim().startsWith("{")) {
          try {
            const nestedData = JSON.parse(data.body);
            data = {
              id: nestedData.id || data.id,
              type: nestedData.type || data.type,
              username: nestedData.username || data.username,
              body: nestedData.body,
              time: nestedData.time || data.time
            };
          } catch (e) {
            console.error("Failed to parse nested message body JSON:", e);
          }
        }

        // 🚀 LIVE DELETE RECEPTION: Catch delete signals from the pool
        if (data.type === "delete") {
          this.setState(prev => ({
            chatHistory: prev.chatHistory.filter(m => m.id !== data.body)
          }));
          return;
        }

        this.setState((prev) => ({
          chatHistory: [...prev.chatHistory, data]
        }));
      });

      // Handshake to let Go connect usernames to pools
      setTimeout(() => {
        const joinPayload = {
          type: "system",
          username: finalizedName,
          body: "joined"
        };
        sendMsg(JSON.stringify(joinPayload));
      }, 300);
    });
  }

  send = (text) => {
    if (!text.trim()) return;

    const payload = {
      type: "chat",
      username: this.state.username,
      body: text,
      time: new Date().toLocaleTimeString([], { hour: '2-digit', minute: '2-digit' })
    };

    sendMsg(JSON.stringify(payload));
  };
  
  clearChatLocally = () => {
    this.setState({ chatHistory: [] });
  };

  render() {
    return (
      <div className="App">
        <Header onClear={this.clearChatLocally} />
        
        <ChatHistory 
          chatHistory={this.state.chatHistory} 
          currentUsername={this.state.username} 
          sendMsg={sendMsg} /* 🚀 Passing WebSocket transmission function down */
        />
        <ChatInput send={this.send} />
      </div>
    );
  }
}

export default App;