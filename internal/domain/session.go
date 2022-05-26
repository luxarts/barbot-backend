package domain

type SessionDTO struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
type SessionResponse struct {
	Token string `json:"token"`
}
