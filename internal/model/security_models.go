package model

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RefreshToken struct {
	Token string `json:"token"`
}

// Convert object of type model.Credentials to model.NewUser
func CredsToNewUser(creds Credentials) NewUser {
	return NewUser{
		Username: creds.Username,
		Password: creds.Password,
	}
}
