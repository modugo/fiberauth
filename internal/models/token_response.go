package models

type TokenResponse struct {
	Type         string `json:"type"`
	AccessToken  string `json:"accessToken,omitempty"`
	RefreshToken string `json:"refreshToken,omitempty"`
	FactorsToken string `json:"factorsToken,omitempty"`
}
