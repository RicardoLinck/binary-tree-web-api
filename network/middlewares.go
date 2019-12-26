package network

import (
	"net/http"
	"time"

	"github.com/google/uuid"
	"go.uber.org/zap"
)

const correlationID string = "correlation-id"

type correlatedResponseWriter struct {
	http.ResponseWriter
	correlationID string
	statusCode    int
	timestamp     time.Time
}

func (crw *correlatedResponseWriter) WriteHeader(statusCode int) {
	crw.statusCode = statusCode
	crw.ResponseWriter.WriteHeader(statusCode)
}

func correlationIDMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		uuid := uuid.New()
		crw := &correlatedResponseWriter{w, uuid.String(), http.StatusOK, time.Now().UTC()}
		next.ServeHTTP(crw, r)
	})
}

func (lc *srv) loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		crw := w.(*correlatedResponseWriter)
		uuid := crw.correlationID
		next.ServeHTTP(w, r)
		finish := time.Since(start).Nanoseconds()
		lc.logger.Info("Request",
			zap.String("url", r.RequestURI),
			zap.Int("statusCode", crw.statusCode),
			zap.Int64("elapsedTime", finish),
			zap.String("correlationId", uuid))
	})
}
