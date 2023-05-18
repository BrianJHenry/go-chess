import "./ChessSquare.css"
import { ReactNode } from "react";

type ChessSquareProps = {
    color: string;
    children: ReactNode;
};

const ChessSquare = ({color, children}: ChessSquareProps) => {

    return (
        <div className="chess-square" style={{backgroundColor: color}}>
            {children}
        </div>
    );
};

export default ChessSquare;