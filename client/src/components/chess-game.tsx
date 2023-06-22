import useWebSocket, { ReadyState } from "react-use-websocket";
import "../styles/chess-game.css"
import ChessBoard from "./chess-board"
import { useCallback, useEffect, useState } from "react";
import { ChessMessage, ChessMove, ChessState, DefaultChessState } from "../classes/chess-data";

type ChessGameProps = {
    gameID: string;
};

const ChessGame = ({ gameID }: ChessGameProps) => {

    const [gameState, setGameState] = useState<ChessState>(DefaultChessState);
    const [gameEnd, setGameEnd] = useState("Continuing")
    const [statusMessage, setStatusMessage] = useState("Normal")

    const url = "ws://localhost:3000/game/" + gameID;
    const { sendJsonMessage, lastJsonMessage, readyState } = useWebSocket<ChessMessage>(url);

    // TODO: actual implementation
    const handleSendMove = useCallback((move: ChessMove) => sendJsonMessage(move), []);

    // TODO: actual implementation
    useEffect(() => {
        if (lastJsonMessage != null) {
            if (lastJsonMessage.messageType == 0) {
                console.log("Recieving gameState");
                setGameState(lastJsonMessage.gameState);
            } else if (lastJsonMessage.messageType == 1) {
                console.log("Recieving gameInfo");
                setGameEnd(lastJsonMessage.messageContent);
                console.log(gameEnd);
            } else if (lastJsonMessage.messageType == 2) {
                console.log("Recieving miscMessage");
                setStatusMessage(lastJsonMessage.messageContent)
                console.log(statusMessage);
            }
        }
    }, [lastJsonMessage]);

    const connectionStatus = {
        [ReadyState.CONNECTING]: "Connecting",
        [ReadyState.OPEN]: "Open",
        [ReadyState.CLOSING]: "Closing",
        [ReadyState.CLOSED]: "Closed",
        [ReadyState.UNINSTANTIATED]: "Uninstantiated",
    }[readyState];

    return (
        <div>
            <ChessBoard boardState={gameState} moveHandler={handleSendMove}/>
            <h2>My Turn: {gameState.turn ? "Yes" : "No"}</h2>
            <h2>Game End State: {gameEnd}</h2>
            <h2>Status: {statusMessage}</h2>
            <span>The WebSocket is currently {connectionStatus}</span>
        </div>
    );
};

export default ChessGame;