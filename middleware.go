package main

import (
	"log"
	"net/http"
	"os"
	"strings"
	"time"
)

// withCORS lets the FE call this API from the browser. ALLOWED_ORIGINS is a
// comma-separated list (e.g. the Netlify site URL in production); when it's
// unset the API stays wide open ("*") so local dev from localhost or a LAN IP
// (see project memory on the double-router dev setup) keeps working.
func withCORS(next http.Handler) http.Handler {
	allowed := map[string]bool{}
	for _, origin := range strings.Split(os.Getenv("ALLOWED_ORIGINS"), ",") {
		if origin = strings.TrimRight(strings.TrimSpace(origin), "/"); origin != "" {
			allowed[origin] = true
		}
	}

	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if len(allowed) == 0 {
			w.Header().Set("Access-Control-Allow-Origin", "*")
		} else {
			w.Header().Add("Vary", "Origin")
			if origin := r.Header.Get("Origin"); allowed[origin] {
				w.Header().Set("Access-Control-Allow-Origin", origin)
			}
		}
		w.Header().Set("Access-Control-Allow-Methods", "GET, PATCH, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func withLogging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		log.Printf("%s %s (%s)", r.Method, r.URL.Path, time.Since(start))
	})
}
