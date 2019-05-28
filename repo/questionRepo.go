package repo

import "diploma/domain"

const readQuestionQuery = `SELECT users.name,users.surname,questions.title,questions.text 
						  FROM questions INNER JOIN users ON users.id = questions.author_id
						  WHERE questions.id = ?`
const readAnswersOfQuestionQuery = `SELECT users.name,users.surname,answers.text 
								   FROM answers INNER JOIN questions ON questions.id = answers.question_id
								   INNER JOIN users ON answers.author_id = users.id
								   WHERE questions.id = ? `
const readAllQuestionsQuery = `SELECT questions.id,users.name,users.surname, questions.title,questions.text 
							  FROM questions INNER JOIN users ON users.id = questions.author_id
							  WHERE questions.lawer_id IS NULL`
const readActiveQuestionsQuery = `SELECT questions.id,users.name,users.surname, questions.text 
								 FROM questions INNER JOIN users ON users.id = questions.author_id
								 WHERE author_id = ?`
const createQuestionQuery = `INSERT INTO questions (author_id,title,text) VALUES (
										(SELECT id FROM users WHERE token =?),?,(SELECT id FROM specializations WHERE title =?),?)`
const updateQuestionQuery = `UPDATE questions SET title = ?, text = ? WHERE id = ?`
const deleteQuestionQuery = `DELETE FROM questions WHERE id = ?`

type QuestionRepo interface {
	CreateQuestion(q domain.Question) error
	UpdateQuestion (id int,q domain.Question) error
	FindQuestion (id int) (domain.Question, error)
	AllQuestions (req domain.QuestionSearch) []domain.Question
	DeleteQuestion(id int) error
}

func (m Mysql) CreateQuestion(q domain.Question) error {
	strPre,err := m.DB.Prepare(createQuestionQuery)
	if err != nil {
		return err
	}
	defer strPre.Close()

	_,err = strPre.Exec(q.Author.Token, q.Title, q.Text)
	if err != nil {
		return err
	}

	return nil
}
func (m Mysql) UpdateQuestion (id int,q domain.Question) error {
	strPre,err := m.DB.Prepare(updateQuestionQuery)
	if err != nil {
		return err
	}
	defer strPre.Close()

	_,err = strPre.Exec(q.Title, q.Text)
	if err != nil {
		return err
	}

	return nil
}
func (m Mysql) FindQuestion (id int) (domain.Question, error) {
	var q domain.Question

	strPre,err := m.DB.Prepare(readQuestionQuery)
	if err != nil {
		return q,err
	}
	defer strPre.Close()

	err = strPre.QueryRow(id).Scan(&q.Author.Name,&q.Author.Surname,&q.Title,&q.Text)
	if err != nil {
		return q,err
	}

	return q,nil
}
func (m Mysql) AllQuestions (req domain.QuestionSearch) []domain.Question {
	var questions []domain.Question
	strPre,err := m.DB.Prepare(readAllQuestionsQuery)
	if err != nil {
		return nil
	}
	defer strPre.Close()

	rows,err := strPre.Query()
	defer rows.Close()

	for rows.Next() {
		var q domain.Question
		err = rows.Scan(&q.Author.Name,&q.Author.Surname,&q.Title,&q.Text)
		if err != nil {
			return nil
		}
		questions = append(questions,q)
	}

	return questions
}
func (m Mysql) DeleteQuestion(id int) error {
	strPre,err := m.DB.Prepare(deleteQuestionQuery)
	if err != nil {
		return err
	}
	defer strPre.Close()

	_,err = strPre.Exec(id)
	if err != nil {
		return err
	}

	return nil
}
