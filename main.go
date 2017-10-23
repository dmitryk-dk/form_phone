package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"database/sql"
	_ "github.com/go-sql-driver/mysql"

	"github.com/dmitryk-dk/form_phone/server/config"
	"github.com/dmitryk-dk/form_phone/server/database"
	appHandlers "github.com/dmitryk-dk/form_phone/server/handlers"
)

const (
	staticDir = "./build/"
	listen    = ":3000"
)

func dependenciesHandler() http.Handler {
	return http.StripPrefix("/", http.FileServer(http.Dir(staticDir)))
}

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)

	dbCfg := config.GetDBConfig()
	db, err := database.Connect(dbCfg)
	if err != nil {
		log.Fatalf("Error when connecting to DB %s", err)
	}
	// get static files
	depHandler := dependenciesHandler()

	// handle static files
	http.Handle("/", depHandler)

	// handle get request
	http.HandleFunc("/phones", appHandlers.GetHandler)
	//handle post request
	http.HandleFunc("/phone", appHandlers.PostHandler)
	//handle delete request
	http.HandleFunc("/delete", appHandlers.DeleteHandler)
	// ui config
	http.HandleFunc("/uiConfig", appHandlers.UiConfigHandler)

	// prepare server for shutdown
	prepareShutdown(db)

	fmt.Printf("Running server on port: %s\n Type Ctr-c to shutdown server.\n", listen)
	http.ListenAndServe(listen, nil)
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
