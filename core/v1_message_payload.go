package core

type Version1MessageType string

const (
	MSG_LOGIN_REQUEST    Version1MessageType = "LOGIN_REQUEST"        // means the login is required
	MSG_LOGIN_PAYLOAD    Version1MessageType = "LOGIN_PAYLOAD"        // means the login PAYLOAD is required
	MSG_LOGIN_RESPONSE   Version1MessageType = "LOGIN_RESPONSE"       // means the login response
	GAME_START           Version1MessageType = "GAME_START"           // means the login response
	GAME_END             Version1MessageType = "GAME_END"             // means the login response
	PLAYER_MOVE          Version1MessageType = "PLAYER_MOVE"          // means the login response
	PLAYER_MOVE_RESPONSE Version1MessageType = "PLAYER_MOVE_RESPONSE" // means the login response
	ERROR                Version1MessageType = "ERROR"                // means the login response
	WAITING_FOR_OPPONENT Version1MessageType = "WAITING_FOR_OPPONENT"
)

type Version1MessageLoginPayload struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Version1MessageLoginResponse struct {
	IsAuthenticated bool   `json:"is_authenticated"`
	Message         string `json:"message,omitempty"`
	PlayerId        string `json:"player_id,omitempty"`
}

type Version1GameStartPayload struct {
	GameId           string `json:"game_id"`
	YourSymbol       string `json:"your_symbol"`
	OpponentUsername string `json:"opponent_username"`
	YourTurn         bool   `json:"your_turn"`
}

type Version1GameEndPayload struct {
	GameId string `json:"game_id"`
	Result string `json:"result"`
	Winner string `json:"winner"`
}

type Version1PositionMoveRequestPayload struct {
	Position int    `json:"position"`
	Symbol   string `json:"symbol"`
}

type Version1ErrorInfoPayload struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}

type Version1PositionMovedResponsePayload struct {
	MovedToPosition int    `json:"moved_to_position"`
	MovedByUser     string `json:"moved_by_user"`
	TurnSymbol      string `json:"turn_symbol"`
}
