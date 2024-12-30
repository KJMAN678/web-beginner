package main

import (
	"log"
	"net/http"
)

func main() {
	fileServer := http.FileServer(http.Dir("ch06/FileServer/static"))
	stripPrefix := http.StripPrefix("/static/", fileServer)
	http.Handle("/", stripPrefix)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal("failed to start:", err)
	}
}
