import React, { Component } from 'react';
import "./ChatHistory.scss"

class ChatHistory extends Component {
    render() {
        const messages = this.props.chatHistory.map(
            (msg, index) => ( <p key={index}>{msg.data}</p>)
        );
    

        return (
            <div clasName="ChatHistory">
                <h2>ChatHistory</h2>
                {messages}
            </div>

        )
    }
}

export default ChatHistory;