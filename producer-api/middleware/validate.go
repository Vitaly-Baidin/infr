package middleware

import (
	"bytes"
	"encoding/json"
	"github.com/Vitaly-Baidin/producer-api/response"
	"github.com/Vitaly-Baidin/producer-api/user"
	"github.com/go-playground/validator/v10"
	"io"
	"net/http"
)

func ValidateUser(next http.HandlerFunc) func(http.ResponseWriter, *http.Request) {
	return func(rw http.ResponseWriter, r *http.Request) {
		var u user.UserGrade

		body, err := io.ReadAll(r.Body)
		if err != nil {
			response.SendErrorResponse(rw, http.StatusBadRequest, err.Error())
			return
		}

		err = json.Unmarshal(body, &u)
		if err != nil {
			response.SendErrorResponse(rw, http.StatusInternalServerError, err.Error())
			return
		}

		v := validator.New()

		err = v.Struct(u)
		if err != nil {
			response.SendErrorResponse(rw, http.StatusBadRequest, err.Error())
			return
		}

		r.Body = io.NopCloser(bytes.NewBuffer(body))

		next.ServeHTTP(rw, r)
	}
}
