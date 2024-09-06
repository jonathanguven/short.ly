package utils

import (
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

var jwtSecret = []byte(os.Getenv("JWT_SECRET"))

// generate JWT token for user given their user ID
func GenerateJWT(userID uint) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(),
	})
	return token.SignedString(jwtSecret)
}

// set JWT token in HTTP-only cookie
func SetCookie(w http.ResponseWriter, token string) {
	http.SetCookie(w, &http.Cookie{
		Name:     "token",
		Value:    token,
		HttpOnly: true,
		Expires:  time.Now().Add(time.Hour * 24 * 7),
		Path:     "/",
		SameSite: http.SameSiteLaxMode,
	})
}

// retrieve JWT token from a HTTP-only cookie
func GetJWTFromCookie(r *http.Request) (*jwt.Token, error) {
	cookie, err := r.Cookie("token")
	if err != nil {
		return nil, err
	}

	return jwt.Parse(cookie.Value, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
}
