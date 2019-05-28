package handler

import (
	"diploma/domain"
	"diploma/service"
	"diploma/utils"
	"net/http"
	"strconv"
)

func Answer(w http.ResponseWriter, r *http.Request) {
	serv := service.NewAnswerService(Conn)

	if r.Method == http.MethodGet {
		typ := r.URL.Query().Get("type")
		idStr := r.URL.Query().Get("id")

		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		switch typ {
		case "Single":
			question, err := serv.FindAnswer(id)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			b, err := utils.MarshalJSON(question)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			w.Write(b)
			return
		case "All":
			idStr := r.URL.Query().Get("question")
			id, err := strconv.Atoi(idStr)

			request := domain.AnswerSearch{
				Question: id,
			}

			answers := serv.AllAnswers(request)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				return
			}
			b, err := utils.MarshalJSON(answers)
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
		var a domain.Answer

		err := utils.UnmarshalJSON(r.Body, &a)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = serv.CreateAnswer(a)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		return
	}

	if r.Method == http.MethodPut {
		var a domain.Answer

		err := utils.UnmarshalJSON(r.Body, &a)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = serv.UpdateAnswer(a.ID, a)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	if r.Method == http.MethodDelete {
		idStr := r.URL.Query().Get("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		err = serv.DeleteAnswer(id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		return
	}

	w.WriteHeader(http.StatusBadRequest)
	return
}
