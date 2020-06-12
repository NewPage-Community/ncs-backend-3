package rpc

import (
	"fmt"
	"net/http"
)

func HealthCheck(health func() bool) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain")
		if health() {
			fmt.Fprintln(w, "Healthy")
		} else {
			http.Error(w, "Unhealthy", 503)
		}
	}
}
