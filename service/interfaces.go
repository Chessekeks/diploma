package service

import (
	"diploma/domain"
)

//Validator interface for validate services data
type Validator interface {
	Validate() error
}

//UserServiceInterface for methods of User entity
type UserServiceInterface interface {
	SingUp(u domain.User) error
	SingIn(u domain.User) error
}

//QuestionServiceInterface for methods of Question entity
type QuestionServiceInterface interface {
	AskAQuestion(q domain.Question) error
}

//AnswerServiceInterface for methods of Answer entity
type AnswerServiceInterface interface {
	AnswerTheQuestion(a domain.Answer) error
}
