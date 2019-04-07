package main

import (
    "net/http"
    "log"
    "github.com/gorilla/mux"
)



func hello(w http.ResponseWriter, r *http.Request) {
  vars := mux.Vars(r)

  w.Write([]byte("Vars: \n" + vars["id"]))

}

func router() {
    router := mux.NewRouter()
    // Routes consist of a path and a handler function.
    router.HandleFunc("/", YourHandler)
    router.HandleFunc("/hello/{id}", hello).Methods("GET")

    // Bind to a port and pass our router in
    log.Fatal(http.ListenAndServe(":8000", router))
}
