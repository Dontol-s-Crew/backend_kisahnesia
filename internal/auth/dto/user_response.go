package dto

import (
	"time"

	"github.com/Dontol-s-Crew/backend_kisahnesia.git/internal/auth/entity"
)

type UserResponse struct {
	Id           int64     `json:"Id"`
	Password     string    `json:"Password"`
	Email        string    `json:"Email"`
	Is_admin     bool      `json:"Is_admin"`
	Nama         string    `json:"nama"`
	Time_updated time.Time `json:"Time_updated"`
	Time_created time.Time `json:"Time_created"`
}

type UserResponses []*UserResponse

func FromEntityToJson(data entity.Userdbs) UserResponses {
	var temp UserResponses
	for _, val := range data {
		var temp2 UserResponse
		temp2.Email = val.Email
		temp2.Password = val.Password
		temp2.Id = val.Id
		temp2.Nama = val.Nama
		temp2.Is_admin = val.Is_admin
		temp2.Time_created = val.Time_created
		temp2.Time_updated = val.Time_updated
		temp = append(temp, &temp2)
	}
	return temp
}
