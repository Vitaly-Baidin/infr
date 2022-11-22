package middleware

import (
	"github.com/Vitaly-Baidin/infr/internal/response"
	"net/http"
)

type BasicAuth struct {
	username string
	password string
	next     http.Handler
}

func NewBasicAuth(username string, password string, h http.Handler) *BasicAuth {
	return &BasicAuth{username: username, password: password, next: h}
}

func (m *BasicAuth) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	username, password, ok := r.BasicAuth()
	if !ok {
		response.SendErrorResponse(rw, http.StatusBadRequest, "need basic auth")
		return
	}

	if m.username != username || m.password != password {
		response.SendErrorResponse(rw, http.StatusUnauthorized, "incorrect username or password")
		return
	}

	m.next.ServeHTTP(rw, r)
}
