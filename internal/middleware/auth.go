package middleware

import (
	"context"
	"net/http"
	"strings"

	authService "github.com/fahmiabd/go-order-api/internal/services/auth"
)

type contextKey string

const userIDKey contextKey = "user_id"

func AuthMiddleware(authService authService.IAuthService) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			authHeader := r.Header.Get("Authorization")
			if authHeader == "" {
				http.Error(w, "authorization header required", http.StatusUnauthorized)
				return
			}

			parts := strings.Split(authHeader, " ")
			if len(parts) != 2 || parts[0] != "Bearer" {
				http.Error(w, "invalid authorization header format", http.StatusUnauthorized)
				return
			}

			userID, err := authService.ValidateToken(parts[1])
			if err != nil {
				http.Error(w, "invalid or expired token", http.StatusUnauthorized)
				return
			}

			// inject user_id ke context
			ctx := context.WithValue(r.Context(), userIDKey, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		})
	}
}

func UserIDFromContext(ctx context.Context) uint {
	userID, ok := ctx.Value(userIDKey).(uint)
	if !ok {
		return 0
	}
	return userID
}
