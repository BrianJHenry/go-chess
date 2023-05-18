import "./ChessBoard.css";
import ChessSquare from "../ChessSquare/ChessSquare";

import {ReactComponent as BlackKing} from "../../assets/black-pieces/BlackKing.svg"
import {ReactComponent as BlackQueen} from "../../assets/black-pieces/BlackQueen.svg"
import {ReactComponent as BlackRook} from "../../assets/black-pieces/BlackRook.svg"
import {ReactComponent as BlackBishop} from "../../assets/black-pieces/BlackBishop.svg"
import {ReactComponent as BlackKnight} from "../../assets/black-pieces/BlackKnight.svg"
import {ReactComponent as BlackPawn} from "../../assets/black-pieces/BlackPawn.svg"
import {ReactComponent as WhiteKing} from "../../assets/white-pieces/WhiteKing.svg"
import {ReactComponent as WhiteQueen} from "../../assets/white-pieces/WhiteQueen.svg"
import {ReactComponent as WhiteRook} from "../../assets/white-pieces/WhiteRook.svg"
import {ReactComponent as WhiteBishop} from "../../assets/white-pieces/WhiteBishop.svg"
import {ReactComponent as WhiteKnight} from "../../assets/white-pieces/WhiteKnight.svg"
import {ReactComponent as WhitePawn} from "../../assets/white-pieces/WhitePawn.svg"

type ChessBoardProps = {
    boardState: number[];
}

const ChessBoard = ({boardState}: ChessBoardProps) => {

    const pieces: any = [
        <BlackKing />,
        <BlackQueen />,
        <BlackRook />,
        <BlackBishop />,
        <BlackKnight />,
        <BlackPawn />,
        <></>,
        <WhitePawn />,
        <WhiteKnight />,
        <WhiteBishop />,
        <WhiteRook />,
        <WhiteQueen />,
        <WhiteKing />,
    ];

    return (
        <div className="chess-board">
            {boardState.map((squareState, index) => {
                const color: string = ((index % 2 + Math.floor(index / 8)) % 2 == 1) ? "#B7C0D8" : "#E8EDF9";
                return (
                    <ChessSquare color={color}>{pieces[squareState+6]}</ChessSquare>
                )
            })}
        </div>
    );
};

export default ChessBoard;