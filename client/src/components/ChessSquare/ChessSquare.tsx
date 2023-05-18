import "./ChessSquare.css"
import { ReactNode } from "react";

type ChessSquareProps = {
    color: string;
    isActive: boolean;
    clickHandler: any;
    children: ReactNode;
};

const ChessSquare = ({color, isActive, clickHandler, children}: ChessSquareProps) => {

    const squareColor: string = (isActive) ? "#7B61FF" : color;

    return (
        <div 
            className="chess-square" 
            style={{backgroundColor: squareColor}}
            onClick={clickHandler}>
            {children}
        </div>
    );
};

export default ChessSquare;