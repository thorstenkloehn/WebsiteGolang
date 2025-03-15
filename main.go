package main

import (
    "net/http"
)

func main() {
  server := routereinrichtung() // Einrichten der Routen  (siehe routereinrichtung.go)
    // Starten des HTTP-Servers auf Port 8080
    http.ListenAndServe(":8080", server)
}