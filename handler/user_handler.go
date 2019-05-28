package handler

import (
	"diploma/domain"
	"diploma/service"
	"diploma/utils"
	"encoding/json"
	"log"
	"net/http"
)

//Registration handling process of registration
func Registration(w http.ResponseWriter, r *http.Request) {

	serv := service.NewUserService(Conn)

	if r.Method == http.MethodPost {
		var user domain.User

		err := utils.UnmarshalJSON(r.Body,&user)
		if err != nil {
			log.Println("Error in unmarshal json: Registration ", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		err = serv.Create(user)
		if err != nil {
			log.Println("Error in serv: UserService", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		return
	}
	w.WriteHeader(http.StatusBadRequest)
	return
}

//Login handling process of entrance
func Login(w http.ResponseWriter, r *http.Request) {
	serv := service.NewUserService(Conn)

	if r.Method == http.MethodPost {
		var user domain.User
		err := utils.UnmarshalJSON(r.Body, &user)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		checkUser,err := serv.FindByEmail(user.Email)
		if err != nil {
			w.WriteHeader(http.StatusForbidden)
			return
		}
		if checkUser.Password != user.Password {
			w.WriteHeader(http.StatusForbidden)
			return
		}

		token := utils.GenerateToken(15)
		checkUser.Token = token
		err = serv.Update(checkUser.ID,checkUser)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		response, err := json.Marshal(token)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
		w.Write([]byte(response))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	return
}

//Profile handling profile
func Profile(w http.ResponseWriter, r *http.Request) {
	serv := service.NewUserService(Conn)

	if r.Method == http.MethodGet {
		token := r.URL.Query().Get("token")

		profile,err := serv.FindByToken(token)
		if err != nil {
			log.Println("Profile error- ",err.Error())
			w.WriteHeader(http.StatusNotFound)
			return
		}

		b, err := utils.MarshalJSON(profile)
		if err != nil {
			log.Println("Profile:MarshalJSON error -", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(b)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	return
}

func User(w http.ResponseWriter, r *http.Request) {
	serv := service.NewUserService(Conn)

	if r.Method == http.MethodGet {
		token := r.URL.Query().Get("token")

		profile,err := serv.FindByToken(token)
		if err != nil {
			log.Println("Profile error- ",err.Error())
			w.WriteHeader(http.StatusNotFound)
			return
		}

		lawyer, err := serv.SpecialistProfile(token)
		if err != nil {
			log.Println("Profile error- ",err.Error())
			w.WriteHeader(http.StatusNotFound)
			return
		}

		response := domain.UserResponse{
			User: profile,
			Specialist:lawyer,
		}

		b, err := utils.MarshalJSON(response)
		if err != nil {
			log.Println("Profile:MarshalJSON error -", err.Error())
			w.WriteHeader(http.StatusInternalServerError)
			return
		}

		w.Write(b)
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	return
}