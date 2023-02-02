package middleware

import (
	"log"
	"mime"
	"net/http"
)

func Logging(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.Printf("recv a %s request from %s", req.Method, req.RemoteAddr)
		next.ServeHTTP(w, req)
	})
}

func Validating(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		contextType := req.Header.Get("Content-Type")
		mediaType, _, err := mime.ParseMediaType(contextType)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
		if mediaType != "application/json" {
			http.Error(w, "invalid json Content-Type "+mediaType, http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, req)
	})

}
