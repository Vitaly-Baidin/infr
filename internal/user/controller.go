package user

import (
	"encoding/json"
	"github.com/Vitaly-Baidin/infr/internal/response"
	"io"
	"net/http"
)

type Controller interface {
	Set(rw http.ResponseWriter, r *http.Request)
	Get(rw http.ResponseWriter, r *http.Request)
}

type Ctr struct {
	service Service
}

func NewController(s Service) *Ctr {
	return &Ctr{service: s}
}

func (c *Ctr) Set(rw http.ResponseWriter, r *http.Request) {
	var u UserGrade

	body, err := io.ReadAll(r.Body)
	if err != nil {
		response.SendErrorResponse(rw, http.StatusInternalServerError, err.Error())
		return
	}

	err = json.Unmarshal(body, &u)
	if err != nil {
		response.SendErrorResponse(rw, http.StatusInternalServerError, err.Error())
		return
	}

	err = c.service.Save(&u)
	if err != nil {
		response.SendErrorResponse(rw, http.StatusBadRequest, err.Error())
		return
	}
	response.SendResponseData(rw, u)
}

func (c *Ctr) Get(rw http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()

	key, ok := query["user_id"]
	if !ok || len(key) == 0 {
		response.SendErrorResponse(rw, http.StatusBadRequest, "not found query user_id")
		return
	}

	u, err := c.service.FindByKey(key[0])
	if err != nil {
		response.SendErrorResponse(rw, http.StatusNotFound, err.Error())
		return
	}

	response.SendResponseData(rw, u)
}
