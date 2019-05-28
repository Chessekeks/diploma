package main

import (
	"diploma/handler"
	"diploma/middleware"

	//"diploma/middleware"
	"diploma/repo"
	"log"
	"net/http"
)

func init() {
	handler.Conn = repo.OpenMysql()
}

func main() {
	r := http.NewServeMux()
	r.HandleFunc("/", middleware.Logging(handler.MainHandler))
	r.HandleFunc("/user", middleware.Logging(middleware.CORS(handler.User)))
	r.HandleFunc("/registration", middleware.Logging(middleware.CORS(handler.Registration)))
	r.HandleFunc("/login", middleware.Logging(middleware.CORS(handler.Login)))
	r.HandleFunc("/question", middleware.Logging(middleware.CORS(handler.Question)))
	r.HandleFunc("/answer",middleware.Logging(middleware.CORS(handler.Answer)))
	r.HandleFunc("/profile", middleware.Logging(middleware.CORS(handler.Profile)))

	log.Println("Start serving")
	err := http.ListenAndServe(":8085", r)
	if err != nil {
		log.Fatal(err.Error())
	}
}
