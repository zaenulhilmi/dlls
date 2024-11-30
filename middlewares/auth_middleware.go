package middlewares

import (
	"context"
	"dlls/contracts"
	"net/http"
)

func AuthMiddleware(next http.HandlerFunc, jwt contracts.JWT) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {

		token, err := jwt.ExtractToken(r)

		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Validate the ExtractToken
		claims, err := jwt.ParseJWT(token)

		if err != nil {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "authenticated_user", claims)

		r = r.WithContext(ctx)

		next(w, r)
	}
}
