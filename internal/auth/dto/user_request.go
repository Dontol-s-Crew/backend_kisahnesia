package dto

import (
	"encoding/json"
	"io"

	"github.com/Dontol-s-Crew/backend_kisahnesia.git/internal/auth/entity"
)

type User struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserRegister struct {
	Nama     string `json:"nama"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func TransformBodyToUser(r io.ReadCloser) User {
	var temp User
	json.NewDecoder(r).Decode(&temp)
	return temp
}

func TransformBodyToUserRegister(r io.ReadCloser) UserRegister {
	var temp UserRegister
	json.NewDecoder(r).Decode(&temp)
	return temp
}

func ChangeUserToAuthEntity(user User) entity.Userdb {
	var data entity.Userdb
	data.Email = user.Email
	data.Password = user.Password
	return data
}
func ChangeUserRegisterToAuthEntity(user UserRegister) entity.Userdb {
	var data entity.Userdb
	data.Email = user.Email
	data.Password = user.Password
	data.Nama = user.Nama
	return data
}
