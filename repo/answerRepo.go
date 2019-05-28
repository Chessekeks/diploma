package repo

import "diploma/domain"

const createAnswerQuery = `INSERT INTO answers (question_id, text, author_id) VALUES (?,?, (SELECT id
																							FROM users
																							WHERE token =?));
							UPDATE questions SET lawer_id = (SELECT id FROM users WHERE token = ?) WHERE question_id = ?`
const updateAnswerQuery = `UPDATE answers SET text = ? WHERE id = ?`
const readAnswerQuery = ``
const deleteAnswerQuery = `DELETE FROM answers WHERE id = ?`

type AnswerRepo interface {
	CreateAnswer(a domain.Answer) error
	UpdateAnswer(id int, a domain.Answer) error
	FindAnswer(id int) (domain.Answer, error)
	AllAnswers(req domain.AnswerSearch) []domain.Answer
	DeleteAnswer(id int) error
}

func (m Mysql) CreateAnswer(a domain.Answer) error {
	strPre, err := m.DB.Prepare(createAnswerQuery)
	if err != nil {
		return err
	}
	defer strPre.Close()

	_, err = strPre.Exec()
	if err != nil {
		return err
	}

	return nil
}
func (m Mysql) UpdateAnswer(id int, a domain.Answer) error {
	strPre, err := m.DB.Prepare(updateAnswerQuery)
	if err != nil {
		return err
	}
	defer strPre.Close()

	_, err = strPre.Exec()
	if err != nil {
		return err
	}

	return nil
}
func (m Mysql) FindAnswer(id int) (domain.Answer, error) {
	var a domain.Answer

	strPre, err := m.DB.Prepare(readAnswerQuery)
	if err != nil {
		return a, err
	}
	defer strPre.Close()

	err = strPre.QueryRow(id).Scan()
	if err != nil {
		return a, err
	}

	return a, nil
}
func (m Mysql) AllAnswers(req domain.AnswerSearch) []domain.Answer {
	var answers []domain.Answer
	strPre, err := m.DB.Prepare(readAnswersOfQuestionQuery)
	if err != nil {
		return nil
	}
	defer strPre.Close()

	rows, err := strPre.Query()
	defer rows.Close()

	for rows.Next() {
		var q domain.Answer
		err = rows.Scan()
		if err != nil {
			return nil
		}
		answers = append(answers, q)
	}

	return answers
}
func (m Mysql) DeleteAnswer(id int) error {
	strPre, err := m.DB.Prepare(deleteAnswerQuery)
	if err != nil {
		return err
	}
	defer strPre.Close()

	_, err = strPre.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
