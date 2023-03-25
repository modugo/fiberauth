package models

type SignInPayload struct {
	Login    string `json:"login"`
	Password string `json:"password"`
}
