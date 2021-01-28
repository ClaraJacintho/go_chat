import React, { Component } from 'react';
import Message from "../Message/Message"
import "./ChatHistory.scss"

class ChatHistory extends Component {
    render() {
        console.log(this.props.chatHistory);
        const messages = this.props.chatHistory.map(msg => <Message message={msg.data} />);
      
        return (
          <div className='ChatHistory'>
            <h2>Chat</h2>
            {messages}
          </div>
        );
      };
}

export default ChatHistory;