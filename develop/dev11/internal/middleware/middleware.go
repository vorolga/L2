package middleware

import (
	"net/http"
	"time"

	"go.uber.org/zap"
)

type Middleware struct {
	logger *zap.SugaredLogger
}

func NewMiddleware(logger *zap.SugaredLogger) *Middleware {
	return &Middleware{
		logger: logger,
	}
}

func (m Middleware) AccessLogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		m.logger.Info(
			zap.String("URL", r.URL.Path),
			zap.String("METHOD", r.Method),
		)
		start := time.Now()
		next.ServeHTTP(w, r)
		responseTime := time.Since(start)

		m.logger.Info(
			zap.String("URL", r.URL.Path),
			zap.String("METHOD", r.Method),
			zap.Duration("TIME FOR ANSWER", responseTime),
		)
	})
}
