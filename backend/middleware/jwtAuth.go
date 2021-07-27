package middleware

import (
	"bt/project/controller"
	"fmt"
	"log"
	"net/http"

	"github.com/dgrijalva/jwt-go"
)

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, err := r.Cookie("jwt")
		if err != nil {
			statusCode := http.StatusUnauthorized
			http.Error(w, "Token doesnt exist", statusCode)
			fmt.Println(err)

		} else {

			token, err := jwt.ParseWithClaims(c.Value, &jwt.StandardClaims{}, func(t *jwt.Token) (interface{}, error) {
				return []byte(controller.SecretKey), nil
			})
			if err != nil {
				statusCode := http.StatusUnauthorized
				http.Error(w, "Unauthenticated", statusCode)

			} else {
				claims := token.Claims
				log.Println(claims)
				// json.NewEncoder(w).Encode(claims)
				next.ServeHTTP(w, r)
			}

		}

	})
}
