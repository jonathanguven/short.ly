package middlewares

import (
	"context"
	"net/http"
	"shortly/internal/utils"

	"github.com/golang-jwt/jwt/v4"
)

// protects routes; checks JWT token in cookie
func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := utils.GetJWTFromCookie(r)
		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		claims, ok := token.Claims.(jwt.MapClaims)
		if ok && token.Valid {
			userID := uint(claims["user_id"].(float64)) // extract user ID from jwt token
			ctx := context.WithValue(r.Context(), UserIDKey{}, userID)
			next.ServeHTTP(w, r.WithContext(ctx))
		} else {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
		}
	})
}
