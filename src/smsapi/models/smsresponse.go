package models

type SMSResponse struct {
	Message string `json:"message"`
	Error   string `json:"error"`
}
