package middleware

import (
	"log"
	"net/http"
)

type HttpHandler http.Handler

//LoggerMW writes logs to stdout
func LoggerMW(next HttpHandler) HttpHandler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("in LoggerMW")
		defer log.Println("LoggerMW ended")
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()
		log.Printf("request received: %v\trequest Method: %v\trequest Body: %+v\n", r.RequestURI, r.Method, r.Body)
		next.ServeHTTP(w, r)
	})
}

//HeaderValidatorMW validate request header
func HeaderValidatorMW(next HttpHandler) HttpHandler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Println("in HeaderValidatorMW")
		defer log.Println("HeaderValidatorMW ended")
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()
		//validating header
		contentType := r.Header.Get("Content-Type")
		if contentType == "" && (r.Method == "POST" || r.Method == "PUT") {
			http.Error(w, "Error: Request Header must contain Json Content-Type", http.StatusBadRequest)
			return
		}
		next.ServeHTTP(w, r)
	})
}
