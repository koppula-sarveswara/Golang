package main

import (
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "contact.html")
	})

	http.HandleFunc("/contact", func(w http.ResponseWriter, r *http.Request) {
		// Parse the form data
		err := r.ParseForm()
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Display the form data in a tabular format
		w.Write([]byte("<html><head><title>Contact Us</title></head><body>"))
		w.Write([]byte("<h1>Contact Us</h1>"))
		w.Write([]byte("<table><tr><th>Field</th><th>Value</th></tr>"))
		for key, value := range r.Form {
			w.Write([]byte("<tr><td>" + key + "</td><td>" + value[0] + "</td></tr>"))
		}
		w.Write([]byte("</table></body></html>"))
	})

	http.ListenAndServe(":8081", nil)
}
