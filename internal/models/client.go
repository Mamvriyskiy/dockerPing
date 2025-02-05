package models

type Client struct {
	ClientID string `db:"clientid"`
	Login    string `db:"login" json:"login`
	// Email string `db:"email" json:"email`
}

type ClientHandler struct {
	Client
	Password string `json:"password"`
}

type ClientService struct {
	Client
	Password string `db:"password" json:"password`
}

type ClientData struct {
	Client
}
