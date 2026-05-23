import "./ChatInput.css";

function ChatInput({ send }) {
  const handleKeyDown = (event) => {
    if (event.key === "Enter") {
      send(event.target.value);
      event.target.value = "";
    }
  };

  return (
    <div className="ChatInput">
      <input placeholder="Type a message..." onKeyDown={handleKeyDown} />
    </div>
  );
}

export default ChatInput;