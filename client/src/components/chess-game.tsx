import useWebSocket, { ReadyState } from "react-use-websocket";
import "../styles/chess-game.css"
import ChessBoard from "./chess-board"
import { useCallback, useEffect, useState } from "react";

type ChessGameProps = {
    gameID: string;
};

const ChessGame = ({ gameID }: ChessGameProps) => {

    const [gameState, setGameState] = useState("");

    const url = "ws://localhost:3000/game/" + gameID;
    const { sendJsonMessage, lastJsonMessage, readyState } = useWebSocket(url);

    // TODO: actual implementation
    const handleSendMove = useCallback((move: any) => sendJsonMessage(move), []);

    // TODO: actual implementation
    useEffect(() => {
        if (lastJsonMessage !== null) {
            // TODO: parse lastJsonMessage to get state
            setGameState("")
        }
    }, [lastJsonMessage, setGameState]);

    const connectionStatus = {
        [ReadyState.CONNECTING]: "Connecting",
        [ReadyState.OPEN]: "Open",
        [ReadyState.CLOSING]: "Closing",
        [ReadyState.CLOSED]: "Closed",
        [ReadyState.UNINSTANTIATED]: "Uninstantiated",
    }[readyState];

    return (
        <div>
            <ChessBoard/>
            <span>The WebSocket is currently {connectionStatus}</span>
        </div>
    );
};

export default ChessGame;