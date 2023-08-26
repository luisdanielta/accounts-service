package handlers

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	"github.com/dgrijalva/jwt-go"
)

func ValidateRequiredFields(r *http.Request, fields []string) bool {
	for _, field := range fields {
		if r.FormValue(field) == "" {
			return false
		}
	}
	return true
}

func validateToken(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Validar el método de firma
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func JWTMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString := r.Header.Get("Authorization")
		if tokenString == "" {
			res := Response{
				Status:      http.StatusUnauthorized,
				Msg:         "Missing token.",
				Application: "json",
			}
			res.Send(w)
			return
		}

		// Validar el token
		token, err := validateToken(strings.TrimPrefix(tokenString, "Bearer "))
		if err != nil || !token.Valid {
			res := Response{
				Status:      http.StatusUnauthorized,
				Msg:         "Invalid token.",
				Application: "json",
			}
			res.Send(w)
			return
		}

		// Token válido, puedes acceder a los claims
		claims := token.Claims.(jwt.MapClaims)
		username := claims["username"].(string)

		// Agregar información adicional a la solicitud, si es necesario
		ctx := context.WithValue(r.Context(), "username", username)
		r = r.WithContext(ctx)

		// Continuar con el siguiente middleware o manejador
		next.ServeHTTP(w, r)
	})
}
