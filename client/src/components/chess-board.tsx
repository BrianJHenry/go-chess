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

const ChessBoard = () => {

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

    const [boardState, setBoardState] = useState([
        -4, -2, -3, -5, -6, -3, -2, -4,
        -1, -1, -1, -1, -1, -1, -1, -1,
        0, 0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0,
        1, 1, 1, 1, 1, 1, 1, 1,
        4, 2, 3, 5, 6, 3, 2, 4,
    ]);
    
    const [activeIndex, setActiveIndex] = useState(-1);

    const handleClick = (index: number) => {
        // case where no active square
        if (activeIndex === -1) {
            // check if the clicked square has a piece
            if (boardState[index] !== 0) {
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
        const newBoardState = [...boardState];
        newBoardState[index] = newBoardState[activeIndex];
        newBoardState[activeIndex] = 0;
        setBoardState(newBoardState);
        setActiveIndex(-1);
        return;
    };


    return (
        <div className="chess-board">
            {boardState.map((squareState, index) => {
                const color: string = ((index % 2 + Math.floor(index / 8)) % 2 == 1) ? "#B7C0D8" : "#E8EDF9";
                return (
                    <ChessSquare 
                        color={color} 
                        isActive={activeIndex === index}
                        clickHandler={() => {handleClick(index)}}>
                        {pieces[squareState+6]}
                    </ChessSquare>
                )
            })}
        </div>
    );
};

export default ChessBoard;