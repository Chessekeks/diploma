package utils

import (
	"log"
	"net/http"
)

func ReadCookie(r *http.Request, name string) string {
	c, err := r.Cookie(name)
	if err != nil {
		log.Println(err.Error())
		return ""
	}
	return c.Value
}
