package middleware

import (
	"context"
	"net/http"
	"task-tracker-backend/internal/repository"
	"task-tracker-backend/internal/utils"
)

// TODO singleton repository
var userRepository = repository.UserRepository{}

// Let only requests from authenticated users pass through.
//
// Checks jwt token from header Authorization field and send error something wrong with it.
func OnlyAuthenticated() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			if header == "" {
				http.Error(w, "No Authorization provided", http.StatusUnauthorized)
				return
			}

			username, err := utils.ParseToken(header)
			if err != nil {
				http.Error(w, err.Error(), http.StatusUnauthorized)
				return
			}

			user, err := userRepository.GetByUsername(username)
			if err != nil {
				http.Error(w, "Wrong authorization token", http.StatusUnauthorized)
				return
			}
			ctx := context.WithValue(r.Context(), userCtxKey, user)

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}
