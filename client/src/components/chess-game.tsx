import { ChessInfo, ChessMove, ChessState } from "../classes/chess-data";
import ChessBoard from "./chess-board";
import ChessGameMoves from "./chess-game-moves";
import '../styles/chess-game.css'
import ChessGameInfo from "./chess-game-info";

export type ChessGameProps = {
    gameState: ChessState; 
    moveHandler: (move: ChessMove) => void;
    gameInfo: ChessInfo;
    moveClickHandler: (index: number) => void;
};

const ChessGame = ({ gameState, moveHandler, gameInfo, moveClickHandler }: ChessGameProps) => {

    return (
        <div className="chess-game-container">
            <ChessGameInfo gameInfo={gameInfo}/>
            <ChessBoard sideColor={0} boardState={gameState} moveHandler={moveHandler}/>
            <ChessGameMoves moves={gameState.previousMoves} handleMoveClick={moveClickHandler}/>
        </div>
    )
};

export default ChessGame;