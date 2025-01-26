package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", route)

	// file server is in place, we just need to point a url path at it
	fs := http.FileServer(http.Dir("/static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))

	//HTTP server and listen for connection on port localhost:2222
	http.ListenAndServe(":2222", nil)
}

func route(res http.ResponseWriter, req *http.Request) {
	// Any route you write in website you see that route here
	fmt.Fprintln(res, "Welcome to our HTTP Server:", req.URL.Path)
}
