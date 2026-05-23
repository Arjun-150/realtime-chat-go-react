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
      chatHistory: []
    };
  }

  componentDidMount() {
    connect((msg) => {
  const data = JSON.parse(msg.data);

  this.setState((prev) => ({
    chatHistory: [...prev.chatHistory, data.body]
  }));
});
  }

 send = (msg) => {
  sendMsg(msg);
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