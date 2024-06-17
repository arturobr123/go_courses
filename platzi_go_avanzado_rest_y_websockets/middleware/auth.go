package middleware

import (
	"net/http"
	"platzi/go/rest_websockets/models"
	"platzi/go/rest_websockets/server"
	"strings"

	"github.com/golang-jwt/jwt"
)

var (
	NO_AUTH_NEEDED = []string{
		"/login",
		"/signup",
	}
)

func ShouldCheckToken(r *http.Request) bool {
	for _, path := range NO_AUTH_NEEDED {
		if path == r.URL.Path {
			return false
		}
	}
	return true
}

func CheckAuthMiddleware(s server.Server) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			if !ShouldCheckToken(r) {
				next.ServeHTTP(w, r)
				return
			}

			tokenString := strings.TrimSpace(r.Header.Get("Authorization"))

			_, err := jwt.ParseWithClaims(tokenString, &models.AppClaims{}, func(token *jwt.Token) (interface{}, error) {
				return []byte(s.Config().JwtSecret), nil
			})

			if err != nil {
				http.Error(w, "Unauthorized", http.StatusUnauthorized)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}
