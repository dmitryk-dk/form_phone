package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"database/sql"

	"github.com/dmitryk-dk/from_phone/server/config"
	"github.com/dmitryk-dk/from_phone/server/database"
	appHandlers "github.com/dmitryk-dk/from_phone/server/handlers"
	_ "github.com/go-sql-driver/mysql"
)

func dependenciesHandler() http.Handler {
	return http.StripPrefix("/", http.FileServer(http.Dir("../build/")))
}

func main() {
	// listening port
	const port = "3000"
	cfg := config.GetConfig()
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatalf("Error when connecting to DB %s", err)
	}
	// get static files
	depHandler := dependenciesHandler()

	// handle static files
	http.Handle("/", depHandler)

	// handle get request
	http.HandleFunc("/phones", appHandlers.GetHandler)
	//handle get request
	http.HandleFunc("/phone", appHandlers.PostHandler)

	// prepare server for shutdown
	prepareShutdown(db)

	fmt.Printf("Running server on port: %s\n Type Ctr-c to shutdown server.\n", port)
	http.ListenAndServe(":"+port, nil)
}

func prepareShutdown(db *sql.DB) {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGKILL, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		log.Printf("Got signal %d", <-sig)
		db.Close()
		os.Exit(0)
	}()
}
