package utils

import (
	"crypto/sha256"
	"encoding/json"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

func ReadFromRequestBody(r io.Reader, d interface{}) error {
	buf, err := ioutil.ReadAll(r)
	if err != nil {
		return err
	}
	err = json.Unmarshal(buf, &d)
	if err != nil {
		return err
	}

	return nil
}

func Hashpassword(pass string) string {
	var headsedpass string
	supersecret := os.Getenv("salt")
	headsedpass = supersecret + pass + supersecret
	h := sha256.New()
	h.Write([]byte(headsedpass))
	bs := h.Sum(nil)
	return string(bs)
}

func Generatejwttoken(id int64, nama string, email string, is_admin bool) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["exp"] = time.Now().Add(1 * time.Hour).Unix()
	claims["user_id"] = id
	claims["nama"] = nama
	claims["authorized"] = true
	claims["email"] = email
	claims["is_admin"] = is_admin
	lubang := os.Getenv("KEY")
	tokenString, err := token.SignedString([]byte(lubang))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// fileExtensionilustrasi := filepath.Ext(data.Ilustrasiname.Filename)
// dstilu, err := os.Create("image/ilustrasi/" + iduuid.String() + fileExtensionilustrasi)
//
//	if err != nil {
//		return err
//	}
//
// _, err = io.Copy(dstilu, data.Cover)
// err = dstilu.Close()
//
//	if err != nil {
//		return err
//	}
func Createfile(filepath string, file multipart.File) error {
	dst, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer dst.Close()
	_, err = io.Copy(dst, file)
	if err != nil {
		return err
	}
	return nil
}
