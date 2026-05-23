import { Component } from "react";
import "./App.css";
import Header from "./components/Header/Header";
import ChatHistory from "./components/ChatHistory/ChatHistory";
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
      this.setState((prevState) => ({
        chatHistory: [...prevState.chatHistory, msg]
      }));
    });
  }

  send = () => {
    sendMsg("hello");
  };

  render() {
    return (
      <div className="App">
        <Header />
        <ChatHistory chatHistory={this.state.chatHistory} />
        <button onClick={this.send}>Hit</button>
      </div>
    );
  }
}

export default App;