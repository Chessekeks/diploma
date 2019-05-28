package handler

import (
	"diploma/domain"
	"diploma/service"
	"diploma/utils"
	"net/http"
	"strconv"
)

func Question(w http.ResponseWriter, r *http.Request) {
	serv := service.NewQuestionService(Conn)

	if r.Method == http.MethodGet {
		typ := r.URL.Query().Get("type")
		idStr := r.URL.Query().Get("id")

		id, err:= strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		switch typ {
		case "Single":
			question, err := serv.FindQuestion(id)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			b,err := utils.MarshalJSON(question)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Write(b)
			return
		case "All":
			typ := r.URL.Query().Get("type")

			var request domain.QuestionSearch
			switch typ {
			case "lawyer":
				request.IsLawyer = true
			case "user":
				request.IsLawyer = false
			}

			questions := serv.AllQuestions(request)
			if questions == nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			b,err := utils.MarshalJSON(questions)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Write(b)
			return
		}
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if r.Method == http.MethodPost {
		var q domain.Question

		err := utils.UnmarshalJSON(r.Body,&q)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = serv.CreateQuestion(q)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		return
	}

	if r.Method == http.MethodPut {
		var q domain.Question

		err := utils.UnmarshalJSON(r.Body,&q)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = serv.UpdateQuestion(q.ID,q)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	if r.Method == http.MethodDelete {
		idStr := r.URL.Query().Get("id")
		id,err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		err = serv.DeleteQuestion(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	return
}

