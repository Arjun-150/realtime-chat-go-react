import React, { Component } from "react";
import "./ChatInput.css";

class ChatInput extends Component {
  handleKeyDown = (event) => {
    // If Enter is pressed WITHOUT Shift, send the message
    if (event.key === "Enter" && !event.shiftKey) {
      event.preventDefault(); // Prevents a new line being added after sending
      this.props.send(event.target.value);
      event.target.value = ""; // Clear the box
    }
    // If Shift + Enter is pressed, it naturally goes to a new line
  };

  render() {
    return (
      <div className="ChatInput">
        <textarea 
          placeholder="Type a message... (Shift+Enter for new line)" 
          onKeyDown={this.handleKeyDown} 
          rows="1"
        />
      </div>
    );
  }
}

export default ChatInput;