import React, { Component, createRef } from "react";
import "./ChatHistory.css";

class ChatHistory extends Component {
  messagesEndRef = createRef();
  buttonPressTimer = null;
  state = { longPressId: null };

  componentDidUpdate() { this.scrollToBottom(); }
  componentDidMount() { this.scrollToBottom(); }
  componentWillUnmount() { document.removeEventListener("mousedown", this.handleClickOutside); }

  scrollToBottom = () => {
    this.messagesEndRef.current?.scrollIntoView({ behavior: "smooth" });
  };

  handleClickOutside = (event) => {
    if (this.state.longPressId && !event.target.closest(".delete-menu")) {
      this.setState({ longPressId: null });
      document.removeEventListener("mousedown", this.handleClickOutside);
    }
  };

  handleStart = (id) => {
    if (!id) return;
    this.buttonPressTimer = setTimeout(() => {
      this.setState({ longPressId: id });
      document.addEventListener("mousedown", this.handleClickOutside);
    }, 700);
  };

  handleStop = () => clearTimeout(this.buttonPressTimer);

  handleDelete = (id) => {
    this.props.sendMsg(JSON.stringify({ type: "delete", body: id }));
    this.setState({ longPressId: null });
    document.removeEventListener("mousedown", this.handleClickOutside);
  };

  render() {
    const { chatHistory, currentUsername } = this.props;

    return (
      <div className="ChatHistoryContainer">
        <div className="chatRoomBadge">
          <span className="pulseIndicator"></span> general
        </div>

        <div className="ChatHistoryStream">
          {chatHistory.map((msg, index) => {
            if (msg.type === "system") {
              if (!msg.username || msg.username === "undefined") return null;
              return <div className="systemMessage" key={index}>{msg.username} {msg.body}</div>;
            }

            const isSelf = msg.username === currentUsername;
            const isBeingDeleted = this.state.longPressId === msg.id && msg.id;

            return (
              <div className={`messageRow ${isSelf ? "selfRow" : "otherRow"}`} key={index}>
                <div
                  className={`messageBubble ${isSelf ? "self" : "other"}`}
                  onMouseDown={() => isSelf && this.handleStart(msg.id)}
                  onMouseUp={this.handleStop}
                  onMouseLeave={this.handleStop}
                  onTouchStart={() => isSelf && this.handleStart(msg.id)}
                  onTouchEnd={this.handleStop}
                >
                  {isBeingDeleted && (
                    <div className="delete-menu" onClick={(e) => { e.stopPropagation(); this.handleDelete(msg.id); }}>
                      Delete for Everyone
                    </div>
                  )}

                  <div className="messageTop">
                    {!isSelf && <span className="username">{msg.username || "Anonymous"}</span>}
                    {/* 🚀 PURE CSS MULTILINE: No more ReactMarkdown */}
                    <div className="body">
                      {msg.body}
                    </div>
                  </div>
                  <div className="time">{msg.time}</div>
                </div>
              </div>
            );
          })}
          <div ref={this.messagesEndRef} />
        </div>
      </div>
    );
  }
}

export default ChatHistory;