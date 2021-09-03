package homepage

import (
	"log"
	"net/http"
	"time"
)

const message = "Hello Microservices!"

type Handlers struct {
	logger *log.Logger
}

func (h *Handlers) Home(w http.ResponseWriter, r *http.Request) {
	h.logger.Println("Request Processed...") // log is used in main too

	// Content-type detection takes a while if not explicitly provided.
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(message))
}

// A middleware
func (h *Handlers) Logger(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		startTime := time.Now() // Lets time the request
		defer h.logger.Printf("Request Proccessed in %s", time.Now().Sub(startTime))
		next(w, r)
	}
}

// Setup routes pattern
func (h *Handlers) SetupRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/", h.Logger(h.Home)) // Using a middleware.
	// if you want to change "/" to "/home" you can do it here. No need to go to main anymore.
}

func NewHandlers(logger *log.Logger) *Handlers {
	return &Handlers{
		logger: logger,
	}
}
