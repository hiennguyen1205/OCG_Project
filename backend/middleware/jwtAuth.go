package middleware

import (
	"bt/project/controller"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwt")

		if err != nil {
			fmt.Println(err)
			json.NewEncoder(w).Encode(err)
		} else {
			token, err := jwt.ParseWithClaims(cookie.Value, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
				return []byte(controller.SecretKey), nil
			})
			if err != nil {
				fmt.Println("Unauthenticated 2")
			} else {
				claims := token.Claims
				json.NewEncoder(w).Encode(claims)
				next.ServeHTTP(w, r)
			}

		}

	})
}
