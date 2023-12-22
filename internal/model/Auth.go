package model

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RefreshToken struct {
	Token *string `json:"token,omitempty"`
}
