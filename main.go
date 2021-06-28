package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"

	ctrl "github.com/sourav977/avengers-backend/controllers"
	md "github.com/sourav977/avengers-backend/middleware"
)

func init() {
	if os.Getenv("MONGO_CONNECTION_URL") == "" {
		os.Setenv("MONGO_CONNECTION_URL", "mongodb://mongo:27017")
	}
}
func main() {
	var wait time.Duration
	mux := http.NewServeMux()

	getAllAvengersHandler := http.HandlerFunc(ctrl.GetAllAvengers)
	createNewAvengerHandler := http.HandlerFunc(ctrl.AddAvenger)

	healthcheckHandler := http.HandlerFunc(ctrl.Healthcheck)
	readinessHandler := http.HandlerFunc(ctrl.Readiness)

	//our two main apis
	mux.Handle("/avengers/getAllAvengers", md.LoggerMW(md.HeaderValidatorMW(getAllAvengersHandler)))
	mux.Handle("/avengers/createNewAvenger", md.LoggerMW(md.HeaderValidatorMW(md.AvengerMW(createNewAvengerHandler))))

	//container orch. support
	mux.Handle("/healthcheck", md.LoggerMW(healthcheckHandler))
	mux.Handle("/readiness", md.LoggerMW(readinessHandler))

	//server config
	srv := &http.Server{
		Addr:         "0.0.0.0:8000",
		WriteTimeout: time.Second * 15,
		ReadTimeout:  time.Second * 15,
		IdleTimeout:  time.Second * 60,
		Handler:      mux,
	}
	//start server
	go func() {
		log.Fatalln(srv.ListenAndServe())
	}()

	//graceful server shutdown
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)

	// Block until we receive our signal.
	<-c

	// Create a deadline to wait for.
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	srv.Shutdown(ctx)
	log.Println("shutting down")
	os.Exit(0)

}
