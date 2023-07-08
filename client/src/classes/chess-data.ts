export type ChessInfo = {
    gameID: string;
    connectionStatus: string;
    statusMessage: string;
    turn: boolean;
    gameEnd: string;
};

export type MoveSearch = {
    messageType: number;
    index: number;
}

export type ChessMove = {
    moveType: string;
    oldSquare: number;
    newSquare: number;
};

export type ChessState = {
    turn: boolean;
    board: number[];
    possibleMoves: ChessMove[];
    previousMoves: ChessMove[];
};

export type ChessMessage = {
    messageType: number;
    messageContent: string;
    gameState: ChessState;
};

export const DefaultChessState: ChessState = {
    turn: false,
    board: [-4, -2, -3, -5, -6, -3, -2, -4,
        -1, -1, -1, -1, -1, -1, -1, -1,
        0, 0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0,
        0, 0, 0, 0, 0, 0, 0, 0,
        1, 1, 1, 1, 1, 1, 1, 1,
        4, 2, 3, 5, 6, 3, 2, 4,
    ],
    previousMoves: [],
    possibleMoves: [],
};