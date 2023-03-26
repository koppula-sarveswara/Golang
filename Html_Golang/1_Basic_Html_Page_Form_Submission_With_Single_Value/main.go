package main

import (
    "fmt"
    "net/http"
)

func main() {
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        http.ServeFile(w, r, "index.html")
    })

    http.HandleFunc("/greet", greetHandler)

    http.ListenAndServe(":8081", nil)
}

func greetHandler(w http.ResponseWriter, r *http.Request) {
    name := r.FormValue("name")

	if name == "" {
        http.Error(w, "Please provide a name", http.StatusBadRequest)
        return
    }

    fmt.Fprintf(w, "Hello, %s!", name)
}
