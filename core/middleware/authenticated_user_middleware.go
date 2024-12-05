package middleware

import (
	"context"
	"log"
	"movie_premiuem/core/utils"
	"net/http"
	"strings"
)

func AuthenticatedUserMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			log.Println("Authorization header is missing")
			utils.UnauthorizedError401(w)
			return
		}

		token := strings.TrimSpace(strings.TrimPrefix(authHeader, "Bearer"))
		if token == "" {
			log.Println("Bearer token is missing in the Authorization header")
			utils.UnauthorizedError401(w)
			return
		}

		// Verify the token and extract user ID
		userID, err := utils.VerifyToken(token)
		if err != nil {
			utils.UnauthorizedError401(w)
			return
		}

		// pass authenticated user id to context
		ctx := context.WithValue(r.Context(), "AuthenticatedUserID", userID)

		// pass the context
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
