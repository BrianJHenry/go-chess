import { ChessMove, ChessState } from "../classes/chess-data";
import ChessBoard from "./chess-board";
import ChessGameMoves from "./chess-game-moves";
import '../styles/chess-game.css'

export type ChessGameProps = {
    gameState: ChessState; 
    moveHandler: (move: ChessMove) => void;
};

const ChessGame = ({ gameState, moveHandler }: ChessGameProps) => {

    return (
        <div className="chess-game-container">
            <ChessBoard sideColor={0} boardState={gameState} moveHandler={moveHandler}/>
            <ChessGameMoves moves={gameState.previousMoves}/>
        </div>
    )
};

export default ChessGame;