package middleware

import "net/http"

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
