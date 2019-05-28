package repo

import (
	"diploma/domain"
	"log"
)

const readAllDialogs = `SELECT dialogs.id,users.name, users.surname, users.email
						FROM users
						JOIN dialogs ON dialogs.second_user_id = users.id
						WHERE users.token = ? `
const readMessagesOfDialog = `SELECT id, text, time FROM messages WHERE dialog_id = ?`

//CreateDialog is realization of DialogStore interface
func (m Mysql) CreateDialog(firstUserToken string, secondUserToken string) error {
	return nil
}

//ReadAllDialogs is realization of DialogStore interface
func (m Mysql) ReadAllDialogs(userToken string) []domain.Dialog {
	var dialogs []domain.Dialog
	strPre, err := m.DB.Prepare(readAllDialogs)
	if err != nil {
		log.Println("Prepare statement error: ReadAllDialogs")
		return nil
	}
	defer strPre.Close()

	rows, err := strPre.Query(userToken)
	if err != nil {
		log.Println("Query error: ReadAllDialogs")
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var d domain.Dialog
		err := rows.Scan(&d.ID, &d.SecondUser.Name, &d.SecondUser.Surname, &d.SecondUser.Email)
		if err != nil {
			log.Println("Scan error: ReadAllDialogs")
			return nil
		}
		dialogs = append(dialogs, d)
	}
	return dialogs
}

//ReadDialog is realization of DialogStore interface
func (m Mysql) ReadDialog(dialogID string) []domain.Message {
	var messages []domain.Message
	tx, err := m.DB.Begin()
	if err != nil {
		log.Println("Transaction begin error: ReadDialog")
		tx.Rollback()
	}

	stmt, err := tx.Prepare(readMessagesOfDialog)
	if err != nil {
		log.Println("Transaction prepare error: ReadDialog")
		tx.Rollback()
	}
	defer stmt.Close()

	rows, err := stmt.Query(dialogID)
	if err != nil {
		log.Println("Transaction query error: ReadDialog")
		tx.Rollback()
	}
	defer rows.Close()

	for rows.Next() {
		var msg domain.Message
		err = rows.Scan(&msg.ID, &msg.Text, &msg.Time)
		if err != nil {
			log.Println("Transaction scan error: ReadDialog ", err.Error())
			tx.Rollback()
		}
		messages = append(messages, msg)
	}
	tx.Commit()
	return messages
}

//CreateMessage is realization of MessageStore interface
func (m Mysql) CreateMessage(senderToken string, recipientToken string, text string) error {
	return nil
}
