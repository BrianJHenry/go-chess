import "../styles/play-online.css"
import { useEffect, useState } from "react";
import ChessConnection from "../components/chess-connection";

const findGameEndpoint = "http://localhost:3000/findGame/2"

const findGame = async () => {
    const response = await fetch(findGameEndpoint);
    const jsonResponse = await response.json();
    return JSON.stringify(jsonResponse);
}

const PlayOnline = () => {
    const [gameID, setGameID] = useState("...")

    useEffect(() => {
        findGame().then(
            result => setGameID(result)
        );
    }, []);

    return (
        <div className="play-online-container">
            {gameID !== "..." && <ChessConnection gameID={gameID}/>}
        </div>
    );
};

export default PlayOnline;