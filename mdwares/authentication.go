package mdwares

import (
	"context"
	"net/http"

	"github.com/dee-ex/gocarest/mdwares/tokens"
)

// key type helps us avoid key collision when using context
type key int

// AuthMiddleware verifies the request is authenticated
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenStr := r.Header.Get("token")
		if len(tokenStr) == 0 {
			http.Error(w, "Missing authorization token", http.StatusUnauthorized)
			return
		}

		token, err := tokens.ValidateToken(tokenStr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		claims, err := tokens.ExtractClaims(token)
		if err != nil {
			http.Error(w, err.Error(), http.StatusUnauthorized)
			return
		}

		if claims["your_key"] == nil {
			http.Error(w, "Invalid authorization token", http.StatusUnauthorized)
			return
		}

		const (
			id     key = 0
			object key = 1
		)

		ctx := context.WithValue(r.Context(), id, 0)
		ctx = context.WithValue(ctx, object, "value")

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
