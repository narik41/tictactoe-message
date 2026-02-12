package core

type Version1MessageType string

const (
	MSG_LOGIN_REQUEST  Version1MessageType = "LOGIN_REQUEST"  // means the login is required
	MSG_LOGIN_PAYLOAD  Version1MessageType = "LOGIN_PAYLOAD"  // means the login PAYLOAD is required
	MSG_LOGIN_RESPONSE Version1MessageType = "LOGIN_RESPONSE" // means the login response
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
