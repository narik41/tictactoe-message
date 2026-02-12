package core

import (
	"encoding/json"
	"log"
)

func EncodeMessage(ticTacToeMessage TicTacToeMessage) ([]byte, error) {
	payloadBytes, err := json.Marshal(ticTacToeMessage)
	if err != nil {
		log.Println("Error while seralizating the message", err.Error())
		return nil, err
	}

	return payloadBytes, nil
}

func DecodeMessage(data []byte) (*TicTacToeMessage, error) {
	var msg TicTacToeMessage
	if err := json.Unmarshal(data, &msg); err != nil {
		return nil, err
	}
	return &msg, nil
}
