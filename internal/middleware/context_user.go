package middleware

import (
	"context"
	"task-tracker-backend/internal/model"
)

var userCtxKey = &contextKey{"user"}

type contextKey struct {
	name string
}

// Finds the user from the context.
func GetUserFromContext(ctx context.Context) *model.User {
	user, _ := ctx.Value(userCtxKey).(*model.User)
	return user
}
