package main

import (
	"net/http"

	"github.com/gorilla/mux"
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
		"Notifications",
		"/notifications",
		"POST",
		Notifications,
	},
	Route{
		"Notifications",
		"/notifications",
		"GET",
		NotificationIndex,
	},
	Route{
		"Notifications",
		"/notifications",
		"OPTIONS",
		NotificationIndex,
	},
}
