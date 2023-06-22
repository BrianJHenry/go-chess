package sockets

const (
	UpdateStateMessage = iota
	GameInfoMessage
	MiscMessage
)

type Message struct {
	MessageType    int      `json:"messageType"`
	MessageContent string   `json:"messageContent"`
	GameState      APIState `json:"gameState"`
}

func NewMessage(messageType int, messageContent string, gameState APIState) Message {
	return Message{
		MessageType:    messageType,
		MessageContent: messageContent,
		GameState:      gameState,
	}
}
