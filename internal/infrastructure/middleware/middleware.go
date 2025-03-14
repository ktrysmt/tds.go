package middleware

import (
	"context"
	"net/http"
	"time"

	"tds.go/internal/infrastructure/logger"
)

func Logger(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)

		logger.Info("HTTP Request",
			logger.Field("method", r.Method),
			logger.Field("path", r.URL.Path),
			logger.Field("duration", time.Since(start).String()),
		)
	})
}

func Auth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		token := r.Header.Get("Authorization")
		if token == "" {
			logger.Warn("Unauthorized access attempt",
				logger.Field("ip", r.RemoteAddr),
				logger.Field("path", r.URL.Path),
			)
			http.Error(w, "unauthorized", http.StatusUnauthorized)
			return
		}

		ctx := context.WithValue(r.Context(), "user_id", "some-user-id")
		next.ServeHTTP(w, r.WithContext(ctx))
	})
}
