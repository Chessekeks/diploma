package middleware

import (
	"log"
	"net/http"
	"time"
)

//Logging is middleware for inspect URL and spent time of request
func Logging(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next(w, r)
		delay := time.Since(start)
		log.Println("URL ", r.URL, " Time: ", delay)
	}
}
