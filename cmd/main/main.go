package main

import (
	"database/sql"
	"net/http"
	"os"
	"strings"

	controller "github.com/Dontol-s-Crew/backend_kisahnesia.git/pkg/controller"
	"github.com/Dontol-s-Crew/backend_kisahnesia.git/pkg/database"
	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func getEnvVariableValues() map[string]string {
	envVariables := make(map[string]string)

	envVariables["SERVER_ADDRESS"] = os.Getenv("SERVER_ADDRESS")
	// envVariables["FIREBASE_CREDENTIALS_PATH"] = os.Getenv("FIREBASE_CREDENTIALS_PATH")

	envVariables["DB_ADDRESS"] = os.Getenv("DB_ADDRESS")
	envVariables["DB_USERNAME"] = os.Getenv("DB_USERNAME")
	envVariables["DB_PASSWORD"] = os.Getenv("DB_PASSWORD")
	envVariables["DB_NAME"] = os.Getenv("DB_NAME")

	envVariables["WHITELISTED_URLS"] = os.Getenv("WHITELISTED_URLS")

	return envVariables
}

func initializeDatabase(envVariables map[string]string) *sql.DB {
	return database.GetDatabase(
		envVariables["DB_ADDRESS"],
		envVariables["DB_USERNAME"],
		envVariables["DB_PASSWORD"],
		envVariables["DB_NAME"],
	)
}

func initializeGlobalRouter(envVariables map[string]string) *mux.Router {
	r := mux.NewRouter()

	arrayWhitelistedUrls := strings.Split(envVariables["WHITELISTED_URLS"], ",")

	whitelistedUrls := make(map[string]bool)

	for _, v := range arrayWhitelistedUrls {
		whitelistedUrls[v] = true
	}
	return r
}

func main() {
	envVariable := getEnvVariableValues()
	db := initializeDatabase(envVariable)
	r := initializeGlobalRouter(envVariable)
	controller.InitializeController(r, db)
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:2000"},
		AllowCredentials: true,
	})
	handler := c.Handler(r)
	http.ListenAndServe(":8000", handler)
}
