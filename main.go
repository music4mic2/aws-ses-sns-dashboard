package main

import (
    "fmt"
    "net/http"
    "github.com/gorilla/mux"
)

func bounces(w http.ResponseWriter, r *http.Request) {
}

func main() {
    router := mux.NewRouter().StrictSlash(true)
    router.HandleFunc("/bounces", bounces)
    http.ListenAndServe(":8000", router)
}
