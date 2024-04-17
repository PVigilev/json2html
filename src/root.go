package main

import "net/http"

func GetRootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html; charset=UTF-8")
	http.ServeFile(w, r, rootPageAddress)
}
