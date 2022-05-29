package middlewares

import (
	"net/http"
	"twittergo/domain/config"
)

func CheckDB(next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if config.CheckConection() == 0 {
			http.Error(writer, "Connection with db lost", 500)
			return
		}
		next.ServeHTTP(writer, request)
	}
}
