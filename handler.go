package main

import "net/http"
import "regexp"

func index(w http.ResponseWriter, r *http.Request) {
	productId := r.PathValue("id")
	if productId == "" {
		http.Error(w, "ID is missing", http.StatusBadRequest)
		return
}

match, _ := regexp.MatchString("^[a-zA-Z]+$", productId )
if !match {
		http.Error(w, "Invalid ID format. Only letters are allowed", http.StatusBadRequest)
		return
}

	

}