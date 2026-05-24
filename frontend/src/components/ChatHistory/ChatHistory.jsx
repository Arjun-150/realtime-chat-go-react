import React, { Component } from "react";
import "./ChatHistory.css";

class ChatHistory extends Component {
  render() {
    return (
      <div className="ChatHistory">
        <h2>Chat History</h2>

        {this.props.chatHistory.map((msg, index) => {
          // SYSTEM MESSAGE
          if (msg.type === "system") {
            return (
              <div className="systemMessage" key={index}>
                {msg.username} {msg.body}
              </div>
            );
          }

          // CHAT MESSAGE
          return (
            <div className="messageBubble" key={index}>
              <div className="messageTop">
                <span className="username">
                  {msg.username || "Anonymous"}: 
                </span>
                <span className="body">
                  {msg.body}
                </span>
              </div>
              <div className="time">{msg.time}</div>
            </div>
          );
        })}
      </div>
    );
  }
}

export default ChatHistory;