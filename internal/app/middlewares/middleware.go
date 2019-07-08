package middlewares

import (
	"context"
	"fmt"
	"net/http"

	"github.com/nandaryanizar/golang-webservice-example/internal/app/provider"

	"github.com/nandaryanizar/golang-webservice-example/internal/app/helpers"
	"go.uber.org/zap"
)

// PanicRecoveryMiddleware handles the panic in the handlers.
func PanicRecoveryMiddleware(h http.HandlerFunc, logger *zap.Logger) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if rec := recover(); rec != nil {
				// log the error
				logger.Error(fmt.Sprint(rec))

				// write the error response
				helpers.JSONResponse(w, http.StatusInternalServerError, map[string]interface{}{
					"error": "Internal Error",
				})
			}
		}()

		h(w, r)
	}
}

type key int

const (
	userID key = iota
)

// JwtAuthentication middleware
func JwtAuthentication(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		urlExempt := []string{"/token"}
		requestPath := r.URL.Path

		for _, value := range urlExempt {
			if value == requestPath {
				h.ServeHTTP(w, r)
				return
			}
		}

		tokenHeader := r.Header.Get("Authorization")
		claims := provider.NewClaims()
		if err := provider.ParseAndValidateToken(tokenHeader, claims); err != nil {
			helpers.JSONResponse(w, http.StatusForbidden, map[string]interface{}{
				"error": err.Error(),
			})
			return
		}

		ctx := context.WithValue(r.Context(), userID, claims.UserID)
		r = r.WithContext(ctx)
		h.ServeHTTP(w, r)
	})
}
