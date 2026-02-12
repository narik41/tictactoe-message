package core

// TicTacToeMessage is the main message structure that will be transferred over the TCP connection
// between client and server. It wraps version-specific payloads with metadata.
type TicTacToeMessage struct {
	MessageId string           `json:"messageId"`
	Version   string           `json:"version"`
	Timestamp int64            `json:"timestamp"`
	Payload   interface{}      `json:"payload"`
	Error     ErrorInfoPayload `json:"error,omitempty"`
}

type ErrorInfoPayload struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

// Version1MessagePayload is the version 1 message payload. If the version is v1.0.0,
// then this payload will be present in TicTacToeMessage. Based on MessageType,
// the appropriate handler can process the message.
type Version1MessagePayload struct {
	MessageType     Version1MessageType `json:"messageType"`
	IsAuthenticated bool                `json:"isAuthenticated"`
	Payload         interface{}         `json:"payload"`
}
