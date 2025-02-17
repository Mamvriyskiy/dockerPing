package models

type ClientHandler struct {
	Login    string `json:"login"`
	Email    string `json:"email`
	Password string `json:"password"`
}

type ClientService struct {
	Login    string
	Email    string
	Password string
}

type ClientData struct {
	ClientID string `db:"clientid" json:"clientid"`
	Login    string `db:"login" json:"login"`
	Email    string `db:"email" json:"email"`
	Password string `db:"password" json:"password"`
}
