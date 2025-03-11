package main

import "net/http"
// routereinrichtung ist eine Funktion, die einen http.ServeMux zurückgibt, der die Routen für die Anwendung einrichtet.
// Die Funktion gibt einen http.ServeMux zurück, der die Routen für die Anwendung einrichtet.
func routereinrichtung() * http.ServeMux {
		server := http.NewServeMux()
		fs := http.FileServer(http.Dir("static"))
		server.Handle("/static/", http.StripPrefix("/static/", fs))
	return server
}