import React, { Component, createRef } from "react"; // 🚀 Import createRef
import "./ChatHistory.css";

class ChatHistory extends Component {
  // Create a ref anchor
  messagesEndRef = createRef();

  // Scroll whenever the component updates (new messages arrive)
  componentDidUpdate() {
    this.scrollToBottom();
  }

  // Scroll whenever the component first loads (history fetch)
  componentDidMount() {
    this.scrollToBottom();
  }

  scrollToBottom = () => {
    // This scrolls the invisible div at the bottom into view
    this.messagesEndRef.current?.scrollIntoView({ behavior: "smooth" });
  };

  render() {
    const { chatHistory, currentUsername } = this.props;

    return (
      <div className="ChatHistoryContainer">
        <div className="chatRoomBadge">
          <span className="pulseIndicator"></span> online
        </div>

        <div className="ChatHistoryStream">
          {chatHistory.map((msg, index) => {
            // ... Your existing mapping logic (keep it exactly as you have it)
            if (msg.type === "system") {
                if (!msg.username || msg.username === "undefined") return null;
                const systemText = msg.body.includes(msg.username) ? msg.body : `${msg.username} ${msg.body}`;
                return <div className="systemMessage" key={index}>{systemText}</div>;
            }

            const isSelf = msg.username === currentUsername;
            return (
              <div className={`messageRow ${isSelf ? "selfRow" : "otherRow"}`} key={index}>
                <div className={`messageBubble ${isSelf ? "self" : "other"}`}>
                  <div className="messageTop">
                    {!isSelf && <span className="username">{msg.username || "Anonymous"}</span>}
                    <span className="body">{msg.body}</span>
                  </div>
                  <div className="time">{msg.time}</div>
                </div>
              </div>
            );
          })}
          
          {/* 🚀 THE ANCHOR: This div stays at the very bottom of the stream */}
          <div ref={this.messagesEndRef} />
        </div>
      </div>
    );
  }
}

export default ChatHistory;