import { ChessMove } from "../classes/chess-data";
import "../styles/chess-game-moves.css"

export type ChessGameMovesProps = {
    moves: ChessMove[];
    handleMoveClick: any;
};

const ChessGameMoves = ({ moves, handleMoveClick }: ChessGameMovesProps) => {
    
    const movesOutput: string[] = moves.map<string>((move) => {
       // translate moves to string coded move
       const stringMove = move.oldSquare.toString() + "-" + move.newSquare.toString()
       return stringMove; 
    });

    return (
        <div className="moves-container">
            <ul className="moves-list">
                {movesOutput.map((move, index) => {
                    if (index % 2 === 0) {
                        return (
                            <li key={index / 2}>
                                <div className="move-container">
                                    <button className="move-btn" onClick={() => handleMoveClick(index)}>{move}</button>
                                    <button className="move-btn" onClick={() => handleMoveClick(index+1)}>{movesOutput[index + 1]}</button>
                                </div>
                            </li>
                        );
                    }
                })}
            </ul>
        </div>
    ); 
};

export default ChessGameMoves;