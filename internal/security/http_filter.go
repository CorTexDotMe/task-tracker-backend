package security

import (
	"context"
	"net/http"
	"task-tracker-backend/internal/model"
	"task-tracker-backend/internal/repository"
	"task-tracker-backend/internal/utils"
)

var userCtxKey = &contextKey{"user"}
var userRepository = repository.UserRepository{}

type contextKey struct {
	name string
}

func Filter() func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			header := r.Header.Get("Authorization")

			if header == "" {
				next.ServeHTTP(w, r)
				return
			}

			username, err := utils.ParseToken(header)
			utils.HandleError(err)

			user, err := userRepository.GetByUsername(username)
			if err != nil {
				next.ServeHTTP(w, r)
				return
			}
			ctx := context.WithValue(r.Context(), userCtxKey, user) // may not work cause user is a pointer already

			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		})
	}
}

// ForContext finds the user from the context. REQUIRES Middleware to have run.
func ForContext(ctx context.Context) *model.User {
	user, _ := ctx.Value(userCtxKey).(*model.User)
	return user
}
