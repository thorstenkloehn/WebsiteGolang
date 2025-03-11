package main

import "net/http"

func routereinrichtung() * http.ServeMux {
		server := http.NewServeMux()
		fs := http.FileServer(http.Dir("static"))
		server.Handle("/static/", http.StripPrefix("/static/", fs))
	return server
}