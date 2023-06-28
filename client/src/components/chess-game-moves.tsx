import { ChessMove } from "../classes/chess-data";

export type ChessGameMovesProps = {
    moves: ChessMove[];
};

const ChessGameMoves = ({ moves }: ChessGameMovesProps) => {
    
    const movesOutput: string[] = moves.map<string>((move) => {
       // translate moves to string coded move
       const stringMove = move.oldSquare.toString() + "-" + move.newSquare.toString()
       return stringMove; 
    });

    return (
        <div>
            <ul>
                {movesOutput.map((move, index) => {
                    if (index % 2 === 0) {
                        return (
                            <li>
                                <p>{index / 2 + 1} {move} {movesOutput[index + 1]}</p>
                            </li>
                        );
                    }
                })}
            </ul>
        </div>
    ); 
};

export default ChessGameMoves;