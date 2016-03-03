package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Name        string
	Pattern     string
	Method      string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.
		Methods(route.Method).
		Path(route.Pattern).
		Name(route.Name).
		Handler(route.HandlerFunc)
	}
	return router
}

var routes = Routes{
	Route{
		"Bounces",
		"/bounces",
		"POST",
		Bounces,
	},
	Route{
		"Complaints",
		"/complaints",
		"POST",
		Complaints,
	},
	Route{
		"Deliviries",
		"/Deliveries",
		"POST",
		Deliveries,
	},
}
