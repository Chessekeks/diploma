package service

import (
	"diploma/domain"
	"diploma/repo"
)

type QuestionService struct {
	rep repo.QuestionRepo
}

func NewQuestionService(conn repo.Mysql) QuestionService {
	return QuestionService{
		rep: conn,
	}
}

func (qServ QuestionService) CreateQuestion(q domain.Question) error {
	err := qServ.rep.CreateQuestion(q)
	if err != nil {
		return err
	}
	return nil
}
func (qServ QuestionService) UpdateQuestion(id int, q domain.Question) error {
	err := qServ.rep.UpdateQuestion(id, q)
	if err != nil {
		return err
	}
	return nil
}
func (qServ QuestionService) FindQuestion(id int) (domain.Question, error) {
	q, err := qServ.rep.FindQuestion(id)
	if err != nil {
		return q, err
	}
	return q, nil
}
func (qServ QuestionService) AllQuestions(req domain.QuestionSearch) []domain.Question {
	questions := qServ.rep.AllQuestions(req)
	return questions
}
func (qServ QuestionService) DeleteQuestion(id int) error {
	err := qServ.rep.DeleteQuestion(id)
	if err != nil {
		return err
	}
	return nil
}
