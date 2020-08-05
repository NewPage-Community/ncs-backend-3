package rpc

import (
	"backend/pkg/log"
	"fmt"
	"net/http"
)

// HealthCheck http health check handler
func HealthCheck(health func() bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		if health() {
			if _, err := fmt.Fprintln(w, "Healthy"); err != nil {
				log.Error(err)
			}
		} else {
			http.Error(w, "Unhealthy", 503)
		}
	}
}
