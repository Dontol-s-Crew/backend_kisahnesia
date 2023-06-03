package controller

import (
	"database/sql"

	AuthControllerApiPkg "github.com/Dontol-s-Crew/backend_kisahnesia.git/internal/auth/controller"
	AuthRepositoryApiPkg "github.com/Dontol-s-Crew/backend_kisahnesia.git/internal/auth/repository/impl"
	AuthServiceApiPkg "github.com/Dontol-s-Crew/backend_kisahnesia.git/internal/auth/service/impl"
	CeritaControllerApiPkg "github.com/Dontol-s-Crew/backend_kisahnesia.git/internal/cerita/controller"
	CeritaRepositoryApiPkg "github.com/Dontol-s-Crew/backend_kisahnesia.git/internal/cerita/repository/impl"
	CeritaServiceApiPkg "github.com/Dontol-s-Crew/backend_kisahnesia.git/internal/cerita/service/impl"
	"github.com/gorilla/mux"
)

func InitializeController(router *mux.Router, db *sql.DB) {
	webrouter := router.NewRoute().Subrouter()

	AuthRepository := AuthRepositoryApiPkg.ProvideAuthRepisitoryImpl(db)
	AuthService := AuthServiceApiPkg.ProvideAuthServiceImpl(AuthRepository)
	AuthController := AuthControllerApiPkg.ProvideAuthController(webrouter, AuthService)
	AuthController.InitializeController()
	AuthController.InitializeControllerAuth()

	CeritaRepository := CeritaRepositoryApiPkg.ProvideRepisitoryCeritaImpl(db)
	CeritaServiceApiPkg := CeritaServiceApiPkg.ProvideServiceCeritaImpl(CeritaRepository)
	CeritaController := CeritaControllerApiPkg.ProvideCeritaController(webrouter, *CeritaServiceApiPkg)
	CeritaController.InitializeController()
}
