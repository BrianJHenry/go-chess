import useWebSocket, { ReadyState } from "react-use-websocket";
import "../styles/chess-connection.css"
import { useCallback, useEffect, useState } from "react";
import { ChessInfo, ChessMessage, ChessMove, ChessState, DefaultChessState, MoveSearch } from "../classes/chess-data";
import ChessGame from "./chess-game";

type ChessConnectionProps = {
    gameID: string;
};

const ChessConnection = ({ gameID }: ChessConnectionProps) => {

    const [gameState, setGameState] = useState<ChessState>(DefaultChessState);
    const [gameEnd, setGameEnd] = useState("Continuing")
    const [statusMessage, setStatusMessage] = useState("Normal")

    const url = "ws://localhost:3000/game/" + gameID;
    const { sendJsonMessage, lastJsonMessage, readyState } = useWebSocket<ChessMessage>(url);

    const handleSendMove = useCallback((move: ChessMove) => sendJsonMessage(move), []);
    const handleSearchMove = useCallback((index: number) => {
        console.log("Searching for move.");
        const msg: MoveSearch = {
            messageType: 0,
            index: index,
        };
        sendJsonMessage(msg);
    }, [])

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

    const gameInfo: ChessInfo = {
        gameID: gameID,
        connectionStatus: connectionStatus,
        statusMessage: statusMessage,
        turn: gameState.turn,
        gameEnd: gameEnd,
    };

    return (
        <div className="chess-connection-container">
            <ChessGame gameState={gameState} moveHandler={handleSendMove} gameInfo={gameInfo} moveClickHandler={handleSearchMove}/>
        </div>
    );
};

export default ChessConnection;