package main

import (
	"log"
	"net"
	"net/http"
	"os"
	"strings"
)

func main() {
	VERSION := os.Getenv("VERSION")
	http.HandleFunc("/healthz", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		w.Write([]byte("OK"))
		w.Header().Add("Version", VERSION)
	})
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		status := http.StatusOK
		w.WriteHeader(http.StatusOK)
		for name, values := range r.Header {
			for _, value := range values {
				w.Header().Add(name, value)
				// log.Printf("Name: %s, Value: %s", name, value)
			}
		}
		w.Header().Add("Version", VERSION)
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
