package main

import (
    "net/http"
    "github.com/gorilla/mux"
)


type Route struct {
    Name        string
    Method      string
    HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
  Route{
      "/bounces",
      "POST",
      Bounces,
  },
    Route{
      "/complaints",
      "POST",
      Complaints,
    },
    Route{
      "/Deliveries",
      "POST",
      Deliveries,
    },
}
