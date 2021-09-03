package server

import (
	"crypto/tls"
	"net/http"
	"time"
)

func NewServer(mux *http.ServeMux, serverAddress string) *http.Server {
	// Read https://blog.cloudflare.com/exposing-go-on-the-internet/
	tlsConfig := &tls.Config{
		PreferServerCipherSuites: true,
		CurvePreferences: []tls.CurveID{
			tls.CurveP256,
			tls.X25519,
		},
		MinVersion: tls.VersionTLS12,
		CipherSuites: []uint16{
			tls.TLS_ECDHE_ECDSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_RSA_WITH_AES_256_GCM_SHA384,
			tls.TLS_ECDHE_ECDSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_RSA_WITH_CHACHA20_POLY1305,
			tls.TLS_ECDHE_ECDSA_WITH_AES_128_GCM_SHA256,
			tls.TLS_ECDHE_RSA_WITH_AES_128_GCM_SHA256,
		},
	}

	// IF you wan't custom timeouts, use a custom http.Server.
	// http.ListenAndServe provides a default server.
	srv := &http.Server{
		Addr:         serverAddress, //":8080",
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		IdleTimeout:  120 * time.Second,
		TLSConfig:    tlsConfig,
		Handler:      mux,
	} // The timeouts are quite large by default ~ 10 minutes.

	return srv
}
