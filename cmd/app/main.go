// Package classification Let's Encrypt Authorization Server API.
//
// Use this application to handle ACME challenges.
//
// Terms Of Service:
//
// there are no TOS at this moment, use at your own risk we take no responsibility
//
//     Schemes: http, https
//     Host: localhost:80
//     BasePath: /
//     Version: 0.0.1
//     License: MIT http://opensource.org/licenses/MIT
//     Contact: Ryan Gilles <gilles.ryan.dev@gmail.com>
//
//     Consumes:
//     - application/json
//
//     Produces:
//     - application/json
//	   - text/plain
//
// swagger:meta
package main

import (
	"log"
	"net/http"

	"github.com/rs/cors"

	"LEAuthorizationServer/router"
)

// setupGlobalMiddleware will setup CORS
func setupGlobalMiddleware(handler http.Handler) http.Handler {
	handleCORS := cors.Default().Handler
	return handleCORS(handler)
}

// our main function
func main() {
	// create router and start listen on port 8000
	router := router.NewRouter()
	log.Fatal(http.ListenAndServe(":80", setupGlobalMiddleware(router)))

	// @todo Serve HTTPS
	//log.Fatal(http.ListenAndServeTLS(":8001", "cert.pem", "key.pem", setupGlobalMiddleware(router)))
}
