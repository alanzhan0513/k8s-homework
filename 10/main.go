package main

import (
	"log"
	"math/rand"
	"net"
	"net/http"
	"os"
	"strings"
	"time"

	"httpserver/metrics"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

func main() {
	VERSION := os.Getenv("VERSION")

	metrics.Register()
	http.Handle("/metrics", promhttp.Handler())

	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Add("Version", VERSION)
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		delay := rand.Intn(2001)
		time.Sleep(time.Millisecond * time.Duration(delay))
		w.Header().Add("Version", VERSION)
		status := http.StatusOK
		for name, values := range r.Header {
			for _, value := range values {
				w.Header().Add(name, value)
				// log.Printf("Name: %s, Value: %s", name, value)
			}
		}
		w.WriteHeader(http.StatusOK)
		log.Printf("Ip: %s, Response Status Code: %d\n", ClientIP(r), status)
	})

	http.ListenAndServe(":80", nil)
}

func ClientIP(r *http.Request) string {
	xForwardedFor := r.Header.Get("X-Forwarded-For")
	ip := strings.TrimSpace(strings.Split(xForwardedFor, ",")[0])
	if ip != "" {
		return ip
	}

	ip = strings.TrimSpace(r.Header.Get("X-Real-Ip"))
	if ip != "" {
		return ip
	}

	if ip, _, err := net.SplitHostPort(strings.TrimSpace(r.RemoteAddr)); err == nil {
		return ip
	}

	return ""
}
