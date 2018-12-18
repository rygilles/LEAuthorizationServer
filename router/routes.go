package router

import (
	"net/http"

	handler "LEAuthorizationServer/handlers"
)

// Route type description
type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

// Routes contains all routes
type Routes []Route

var routes = Routes{
	Route{
		"GetChallenges",
		"GET",
		"/challenges",
		handler.GetChallenges,
	},
	Route{
		"GetChallenge",
		"GET",
		"/challenge/{token}",
		handler.GetChallenge,
	},
	Route{
		"GetChallengeValue",
		"GET",
		"/.well-known/acme-challenge/{token}",
		handler.GetChallengeValue,
	},
	Route{
		"PostChallenge",
		"POST",
		"/challenge",
		handler.PostChallenge,
	},
}
