package middleware

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/golang-jwt/jwt/v4"
)

func SetContentTypeMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, rw *http.Request) {
		w.Header().Add("Content-Type", "application/json")
		w.Header().Add("Content-Type", "multipart/form-data")

		next.ServeHTTP(w, rw)
	})
}

func AllowCORSAll(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, rw *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		next.ServeHTTP(w, rw)
	})
}

func Middlewareauth(x http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		reqToken := r.Header.Get("Authorization")
		if reqToken == "" {

			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("error cookie gak ada kue"))
			return
		}
		splitToken := strings.Split(reqToken, "Bearer ")
		reqToken = splitToken[1]
		secret := os.Getenv("KEY")
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(reqToken, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		tokens := token.Claims.(jwt.MapClaims)
		if false == tokens["is_admin"] {
			r.Header.Set("is_admin", "0")
		} else {
			r.Header.Set("is_admin", "true")
		}
		// r.Header.Set("user_id",tokens[])
		if err != nil {
			fmt.Println(err)
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("error cookie kue jelek"))
			return
		}
		if !token.Valid {
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}
		x.ServeHTTP(rw, r)
	})
}

func MiddlewareauthAdmin(x http.Handler) http.Handler {
	return http.HandlerFunc(func(rw http.ResponseWriter, r *http.Request) {
		reqToken := r.Header.Get("Authorization")
		if reqToken == "" {

			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("error cookie gak ada kue"))
			return
		}
		splitToken := strings.Split(reqToken, "Bearer ")
		reqToken = splitToken[1]
		secret := os.Getenv("KEY")
		claims := jwt.MapClaims{}
		token, err := jwt.ParseWithClaims(reqToken, claims, func(t *jwt.Token) (interface{}, error) {
			return []byte(secret), nil
		})
		// fmt.Println(token)
		if false == token.Claims.(jwt.MapClaims)["is_admin"] {
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("bau bukan admin"))
			return
		}
		if err != nil {
			fmt.Println(err)
			rw.WriteHeader(http.StatusBadRequest)
			rw.Write([]byte("error cookie kue jelek"))
			return
		}
		if !token.Valid {
			rw.WriteHeader(http.StatusUnauthorized)
			return
		}
		x.ServeHTTP(rw, r)
	})
}
