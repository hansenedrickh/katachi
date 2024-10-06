package middleware

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
)

func Auth(secretKey string) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			token, err := r.Cookie("token")
			if err != nil {
				fmt.Println("error getting token cookie:", err)
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}

			err = verifyToken(token.Value, secretKey)
			if err != nil {
				fmt.Println("error verifying token:", err)
				http.Redirect(w, r, "/login", http.StatusSeeOther)
				return
			}

			next.ServeHTTP(w, r)
		})
	}
}

func verifyToken(tokenStr, secretKey string) error {
	token, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})
	if err != nil {
		return err
	}

	if !token.Valid {
		return errors.New("invalid token")
	}

	return nil
}
