package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"
	"strings"

	"github.com/Dontol-s-Crew/backend_kisahnesia.git/pkg/database"
	"github.com/gorilla/mux"
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

func handler(w http.ResponseWriter, r *http.Request) {
	return
}

func main() {
	envVariable := getEnvVariableValues()
	_ = initializeDatabase(envVariable)
	r := initializeGlobalRouter(envVariable)
	http.ListenAndServe(":8000", r)
	fmt.Print("work")

}
