package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"
)

// define a home handler function which writes a byte slice containing
// "Hello from Snippetbox" as the response body.
func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello from Snippetbox"))
}

func snippetView(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.PathValue("id"))
	// Extract the value of the id wildcard from the request using r.PathValue()
  // and try to convert it to an integer using the strconv.Atoi() function. If
  // it can't be converted to an integer, or the value is less than 1, we
  // return a 404 page not found response.
	var invalid bool = (err != nil || id < 1)
	if invalid {
		http.NotFound(w, r)
		return
	}

	// Use the fmt.Sprintf() function to interpolate the id value with a
  // message, then write it as the HTTP response.
	msg := fmt.Sprintf("Display a specific input with ID %d...", id)
	w.Write([]byte(msg))
}

func snippetCreate(w http.ResponseWriter, r * http.Request) {
	w.Write([]byte("Display a form for creating a new snippet..."))
}

func main() {
	// Use the http.NewServeMux() function to initialise a new servemux, then
	// register the home function as the handler for the "/" URL pattern.
	mux := http.NewServeMux()
	mux.HandleFunc("/{$}", home)
	mux.HandleFunc("/snippet/view", snippetView)
	mux.HandleFunc("/snippet/view/{id}", snippetView)
	mux.HandleFunc("/snippet/create", snippetCreate)
	
	log.Print("starting server on :4000")
	
	// Use the http.ListenAndServe() function to start a new web server. We pass in
  // two parameters: the TCP network address to listen on (in this case ":4000")
  // and the servemux we just created. If http.ListenAndServe() returns an error
  // we use the log.Fatal() function to log the error message and exit. Note
  // that any error returned by http.ListenAndServe() is always non-nil.
	err := http.ListenAndServe(":4000", mux)
	log.Fatal(err)
}