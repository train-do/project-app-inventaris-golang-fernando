package middleware

import (
	"net/http"

	handler "github.com/train-do/project-app-inventaris-golang-fernando/handler/cms"
)

func Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("token")
		if err != nil || err == http.ErrNoCookie {
			http.Redirect(w, r, "/login", http.StatusUnauthorized)
			return
		}
		if cookie.Value != handler.Token {
			http.Redirect(w, r, "/login", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})
}
