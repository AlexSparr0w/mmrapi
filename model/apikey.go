package model

type Apikey struct {
	ID     string `json:"id"`
	UserID string `json:"user_id"`
	Apikey string `json:"apikey"`
}

type UserResponse struct {
	Email        string `json:"email"`
	TotalMatches int    `json:"matches"`
	TotalMmr     string `json:"mmr"`
	Apikey       string `json:"apikey"`
}
