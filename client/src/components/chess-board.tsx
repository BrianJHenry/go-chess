import "../styles/chess-board.css";
import ChessSquare from "./chess-square";
import { useState } from "react";

import {ReactComponent as BlackKing} from "../assets/black-pieces/BlackKing.svg"
import {ReactComponent as BlackQueen} from "../assets/black-pieces/BlackQueen.svg"
import {ReactComponent as BlackRook} from "../assets/black-pieces/BlackRook.svg"
import {ReactComponent as BlackBishop} from "../assets/black-pieces/BlackBishop.svg"
import {ReactComponent as BlackKnight} from "../assets/black-pieces/BlackKnight.svg"
import {ReactComponent as BlackPawn} from "../assets/black-pieces/BlackPawn.svg"
import {ReactComponent as WhiteKing} from "../assets/white-pieces/WhiteKing.svg"
import {ReactComponent as WhiteQueen} from "../assets/white-pieces/WhiteQueen.svg"
import {ReactComponent as WhiteRook} from "../assets/white-pieces/WhiteRook.svg"
import {ReactComponent as WhiteBishop} from "../assets/white-pieces/WhiteBishop.svg"
import {ReactComponent as WhiteKnight} from "../assets/white-pieces/WhiteKnight.svg"
import {ReactComponent as WhitePawn} from "../assets/white-pieces/WhitePawn.svg"
import { ChessMove, ChessState } from "../classes/chess-data";

export type ChessBoardProps = {
    sideColor: number;
    boardState: ChessState;
    moveHandler: (move: ChessMove) => void;
}

const ChessBoard = ({sideColor, boardState, moveHandler}: ChessBoardProps) => {

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

    const [activeIndex, setActiveIndex] = useState(-1);

    const handleClick = (index: number) => {
        // case where no active square
        if (activeIndex === -1) {
            // check if the clicked square has a piece
            if (boardState.board[index] !== 0) {
                setActiveIndex(index);
                return;
            } else {
                return;
            }
        }
        // case where you click on the active piece
        if (activeIndex === index) {
            setActiveIndex(-1);
            return;
        }
        // case where you click a new square
        console.log("Tried to make move:", activeIndex, " : ", index);
        // make new move
        if (boardState.turn) {
            for (var move of boardState.possibleMoves) {
                console.log("Possible Move:", move.oldSquare, " : ", move.newSquare)
                if (move.oldSquare == activeIndex && move.newSquare == index) {
                    moveHandler(move);
                    break;
                }
            }
        }

        setActiveIndex(-1);
        return;
    };

    const displayReversed = sideColor == 0;

    return (
        <div className="chess-board">
            {boardState.board.map((_, index, board) => {
                const currIndex = displayReversed ? 64 - 1 - index : index;
                const color: string = ((currIndex % 2 + Math.floor(currIndex / 8)) % 2 == 1) ? "#B7C0D8" : "#E8EDF9";
                return (
                    <ChessSquare 
                        key={index}
                        color={color} 
                        isActive={activeIndex === currIndex}
                        clickHandler={() => {handleClick(currIndex)}}>
                        {pieces[board[currIndex]+6]}
                    </ChessSquare>
                )
            })}
        </div>
    );
};

export default ChessBoard;