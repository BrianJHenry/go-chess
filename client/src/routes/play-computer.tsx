import "../styles/play-computer.css"
import { useEffect, useState } from "react";
import ChessConnection from "../components/chess-connection";

const findGameEndpoint = "http://localhost:3000/findGame/1"

const findGame = async () => {
    const response = await fetch(findGameEndpoint);
    const jsonResponse = await response.json();
    return JSON.stringify(jsonResponse);
}

const PlayComputer = () => {
    const [gameID, setGameID] = useState("...")

    useEffect(() => {
        findGame().then(
            result => setGameID(result)
        );
    }, []);

    

    return (
        <div className="play-computer-container">
            {gameID !== "..." && <ChessConnection gameID={gameID}/>}
        </div>
    );
};

export default PlayComputer;