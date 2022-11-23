package usrcommand

import (
	"encoding/json"
	"github.com/Vitaly-Baidin/storage-api/user"
)

type Set struct {
	serv user.Service
}

func NewSetCommand(s user.Service) *Set {
	return &Set{serv: s}
}

func (c *Set) Execute(data []byte) error {
	var u user.UserGrade

	err := json.Unmarshal(data, &u)
	if err != nil {
		return err
	}

	return c.serv.Save(&u)
}
