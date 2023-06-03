package controller

import (
	"net/http"
	"time"

	"github.com/Dontol-s-Crew/backend_kisahnesia.git/internal/auth/dto"
	"github.com/Dontol-s-Crew/backend_kisahnesia.git/internal/global"
	"github.com/Dontol-s-Crew/backend_kisahnesia.git/pkg/middleware"
	"github.com/gorilla/mux"

	ServiceApiPkg "github.com/Dontol-s-Crew/backend_kisahnesia.git/internal/auth/service/api"
)

type AuthController struct {
	X *mux.Router
	S ServiceApiPkg.AuthServiceInterface
}

func ProvideAuthController(X *mux.Router, S ServiceApiPkg.AuthServiceInterface) *AuthController {
	return &AuthController{X: X, S: S}
}

func (AC AuthController) LoginHandler(rw http.ResponseWriter, r *http.Request) {
	user := dto.TransformBodyToUser(r.Body)
	ctx := r.Context()
	tokenjwt, err := AC.S.LoginServiceAuth(ctx, user)
	if err != nil {
		return
	}

	expire := time.Now().Add(11 * time.Minute)
	http.SetCookie(rw, &http.Cookie{
		Name:    "Authorization",
		Value:   tokenjwt,
		Expires: expire,
	})

	rw.Write([]byte("login succesful"))
	return
}
func (AC AuthController) RegisterHandler(rw http.ResponseWriter, r *http.Request) {
	user := dto.TransformBodyToUserRegister(r.Body)
	ctx := r.Context()
	tokenjwt, err := AC.S.RegisterServiceAuth(ctx, user)
	if err != nil {
		return
	}
	expire := time.Now().Add(11 * time.Minute)
	http.SetCookie(rw, &http.Cookie{
		Name:    "Authorization",
		Value:   tokenjwt,
		Expires: expire,
	})
	rw.Write([]byte("Register succesful"))
	return
}
func (AC AuthController) Ping(rw http.ResponseWriter, r *http.Request) {
	rw.Write([]byte("ndak mau asu jangan dipegang"))
}

func (AC AuthController) InitializeController() {
	routes := AC.X.PathPrefix(global.API_PATH_ROOT_AUTH).Subrouter()
	routes.HandleFunc(global.API_PATH_POST_LOGIN, AC.LoginHandler).Methods(http.MethodPost)
	routes.HandleFunc(global.API_PATH_POST_REGISTER, AC.RegisterHandler).Methods(http.MethodPost)

}

func (AC AuthController) InitializeControllerAuth() {
	test := AC.X.PathPrefix("/Ping").Subrouter()
	test.HandleFunc("", AC.Ping).Methods(http.MethodGet)
	test.Use(middleware.Middlewareauth)
}
