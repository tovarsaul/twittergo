package middlewares

import (
	"net/http"
	"twittergo/domain"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if domain.CheckConection() == 0 {
			http.Error(writer, "Connection lost", 500)
			return
		}
		next.ServeHTTP(writer, request)
	}
}
