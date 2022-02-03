package middlewares

import (
	"crypto/subtle"
	"fmt"
	"net/http"

	"github.com/vsyakunin/sg-task/domain/models/config"

	log "github.com/sirupsen/logrus"
)

var (
	users  []config.User
)

func InitBasicAuthMiddleware(config *config.Config) {
	users = append(users, config.Auth.UserOne, config.Auth.UserTwo, config.Auth.Operator)
}

func BasicAuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		login, pass, ok := r.BasicAuth()
		if !ok {
			w.WriteHeader(http.StatusUnauthorized)
			log.Info("request without an authorization")
			return
		}

		if !checkCredentials(login, pass, users) {
			w.WriteHeader(http.StatusForbidden)
			log.Info(fmt.Sprintf("authentication failed. Login used: %s", login))
			return
		}

		next.ServeHTTP(w, r)
	})
}

func checkCredentials(login, pass string, users []config.User) bool {
	for i := range users {
		credentials := users[i]
		isLoginCorrect := subtle.ConstantTimeCompare([]byte(credentials.Login), []byte(login))
		isPassCorrect := subtle.ConstantTimeCompare([]byte(credentials.Password), []byte(pass))
		if isLoginCorrect == 1 && isPassCorrect == 1 {
			return true
		}
	}
	return false
}
