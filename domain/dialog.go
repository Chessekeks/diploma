package domain

import (
	"database/sql"
)

//Dialog defines dialog
type Dialog struct {
	ID         int       `json:"id"`
	FirstUser  User      `json:"firsUser"`
	SecondUser User      `json:"secondUser"`
	Messages   []Message `json:"messages"`
}

//Message defines single message in dialog
type Message struct {
	ID        int            `json:"id"`
	Sender    User           `json:"sender"`
	Recipient User           `json:"recipient"`
	Text      string         `json:"text"`
	Time      sql.NullString `json:"time"`
}

//DialogStore provides methods to work with Dialog entity
type DialogStore interface {
	CreateDialog(firstUserToken string, secondUserToken string) error
	ReadAllDialogs(userToken string) []Dialog
	ReadDialog(dialogID string) []Message
	CreateMessage(senderToken string, recipientToken string, text string) error
}
