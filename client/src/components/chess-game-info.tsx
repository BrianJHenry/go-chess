import { ChessInfo } from "../classes/chess-data";
import "../styles/chess-game-info.css";

export type ChessGameInfoProps = {
    gameInfo: ChessInfo;
};

const ChessGameInfo = ({ gameInfo }: ChessGameInfoProps) => {

    const turn: string = gameInfo.turn ? "White" : "Black";

    return (
        <div className="game-info-container">
            <div className="info-container">
                <div>
                    <h1 className="info-header">Game Info</h1>
                    {gameInfo.gameEnd == "Continuing" ? <p>Turn: {turn}</p> : <p>{gameInfo.gameEnd}</p>}
                </div>
                <div>
                    <p>Status: {gameInfo.statusMessage}</p>
                    <p>Connection: {gameInfo.connectionStatus}</p>
                    <p>GameID: {gameInfo.gameID}</p>
                </div> 
            </div>
        </div>
    );
};

export default ChessGameInfo;