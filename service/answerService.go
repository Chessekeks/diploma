package service

import (
	"diploma/domain"
	"diploma/repo"
)

//AnswerService is service of Answer's scenarios
type AnswerService struct {
	rep repo.AnswerRepo
}

func NewAnswerService(conn repo.Mysql) AnswerService {
	return AnswerService{
		rep: conn,
	}
}

func (aServ AnswerService) CreateAnswer(a domain.Answer) error {
	err := aServ.rep.CreateAnswer(a)
	if err != nil {
		return err
	}
	return nil
}
func (aServ AnswerService) UpdateAnswer(id int, a domain.Answer) error {
	err := aServ.rep.UpdateAnswer(id, a)
	if err != nil {
		return err
	}
	return nil
}
func (aServ AnswerService) FindAnswer(id int) (domain.Answer, error) {
	a, err := aServ.rep.FindAnswer(id)
	if err != nil {
		return a, err
	}
	return a, nil
}
func (aServ AnswerService) AllAnswers(req domain.AnswerSearch) []domain.Answer {
	answers := aServ.rep.AllAnswers(req)
	return answers
}
func (aServ AnswerService) DeleteAnswer(id int) error {
	err := aServ.rep.DeleteAnswer(id)
	if err != nil {
		return err
	}
	return nil
}
