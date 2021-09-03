package main

import (
	"log"
	"net/http"
	"os"

	home "github.com/shubhamdwivedii/simple-go-microservice/homepage"
	server "github.com/shubhamdwivedii/simple-go-microservice/server"
)

var (
	CertFile    = "server.crt" //os.Getenv("CERT_FILE")
	KeyFile     = "server.key" //os.Getenv("KEY_FILE")
	ServiceAddr = ":8080"      //os.Getenv("SERVICE_ADDR")
)

func main() {
	// Its good practice to create a logger, using log directly everywhere will give trouble when you want specific formatting everywhere.
	logger := log.New(os.Stdout, "SHUBHAM: ", log.LstdFlags|log.Lshortfile) // | is bitwise OR operator. It combines LstdFlags and Lshortfile
	// log.New takes io.Writer, prefix string, and logging property flag (log.Lshortfile will add file name)

	mux := http.NewServeMux()

	// giving logger to home is "Dependency Injection" pattern.
	h := home.NewHandlers(logger)

	// mux.HandleFunc("/", h.Home)
	// Use below patten insted.

	// A good pattern
	h.SetupRoutes(mux)

	srv := server.NewServer(mux, ServiceAddr)

	logger.Println("Starting Server...") // Note log is used in homepage too

	// err := http.ListenAndServe(":8080", mux)
	// err := srv.ListenAndServe()
	err := srv.ListenAndServeTLS(CertFile, KeyFile)
	// This will only work on https. Redirect http requests to https

	logger.Fatal(err)
}
