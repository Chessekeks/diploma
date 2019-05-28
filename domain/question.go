package domain

type Question struct {
	ID             int      `json:"ID"`
	Author         User     `json:"author"`
	Title          string   `json:"title"`
	Text           string   `json:"text"`
	Answers        []Answer `json:"answers"`
}

type Answer struct {
	ID int `json:"id"`
	Author     User   `json:"author"`
	QuestionID int    `json:"questionID"`
	Text       string `json:"text"`
}

type QuestionSearch struct {
	IsLawyer bool
}

type AnswerSearch struct {
	Question int
}

type QuestionService interface {
	CreateQuestion(q Question) error
	UpdateQuestion (id int,q Question) error
	FindQuestion (id int) (Question, error)
	AllQuestions (req QuestionSearch) []Question
	DeleteQuestion(id int) error
}

type AnswerService interface {
	CreateAnswer(a Answer) error
	UpdateAnswer(id int, a Answer) error
	FindAnswer(id int) (Answer,error)
	AllAnswers(req AnswerSearch) []Answer
	DeleteAnswer(id int) error
}
