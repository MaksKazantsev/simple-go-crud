package main

import (
	"github.com/MaksKazantsev/go-crud/internal/helper"
	"github.com/MaksKazantsev/go-crud/internal/log"
	"github.com/MaksKazantsev/go-crud/internal/routes"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

func main() {
	l := log.MustStart()
	r := mux.NewRouter()
	routes.RegisterRoutes(r)

	srv := http.Server{
		Addr:         ":8000",
		Handler:      r,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
	}

	// Launching server, using srv struct and handling an error
	l.Info("Server started")
	err := srv.ListenAndServe()
	helper.PanicIfErr(err, "Error, server starting failed.")
}
