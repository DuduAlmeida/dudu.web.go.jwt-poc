package request

type LoginCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type RefreshTokenCredentials struct {
	RefreshToken string `json:"refresh_token"`
}
