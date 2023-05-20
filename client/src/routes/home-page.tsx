import { ChangeEvent, useState } from "react";
import WebSocket from "../api/web-socket";
import "../styles/home-page.css"

const Home = () => {
    const [wsConnect, setWSConnect] = useState(false);
    const [inputText, setInputText] = useState("");

    const buttonClick = () => {
        if (wsConnect) {
            setWSConnect(false);
        } else {
            setWSConnect(true);
        }
    }

    const handleChange = (e: ChangeEvent<HTMLInputElement>) => {
        setInputText(e.target.value);
    };

    return (
        <div>
            <h1>Home Page</h1>
            <input type="text" onChange={handleChange} value={inputText} />
            <button onClick={buttonClick}>Connect to Websocket</button>
            {wsConnect && <WebSocket url={inputText} />}
        </div>
    );
};

export default Home;