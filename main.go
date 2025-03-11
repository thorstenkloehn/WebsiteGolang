package main

import (
    "net/http"
)

func main() {
  server := routereinrichtung()
    // Starten des HTTP-Servers auf Port 8080
    http.ListenAndServe(":8080", server)
}