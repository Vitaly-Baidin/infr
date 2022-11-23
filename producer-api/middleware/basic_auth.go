package middleware

import (
	"github.com/Vitaly-Baidin/producer-api/response"
	"net/http"
)

func BasicAuth(username string, password string, next http.HandlerFunc) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		un, pass, ok := r.BasicAuth()
		if !ok {
			response.SendErrorResponse(rw, http.StatusBadRequest, "need basic auth")
			return
		}

		if un != username || pass != password {
			response.SendErrorResponse(rw, http.StatusUnauthorized, "incorrect username or password")
			return
		}
		next.ServeHTTP(rw, r)
	}
}
