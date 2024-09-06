package middlewares

import (
	"context"
	"net/http"
	"shortly/internal/utils"

	"github.com/golang-jwt/jwt/v4"
)

// protects routes; parses user ID from token and stores in context
func Authenticate(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token, err := utils.GetJWTFromCookie(r)
		if err == nil {
			claims, ok := token.Claims.(jwt.MapClaims)
			if ok && token.Valid {
				// get user ID from token
				if userID, exists := claims["user_id"].(float64); exists {
					// store userID in context
					ctx := context.WithValue(r.Context(), UserIDKey{}, uint(userID))
					r = r.WithContext(ctx)
				}
			}
		}
		next.ServeHTTP(w, r)
	})
}
