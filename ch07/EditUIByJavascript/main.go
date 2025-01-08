package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	http.Handle("/static/",
		http.StripPrefix("/static/", http.FileServer(http.Dir("ch07/EditUIByJavascript/static"))))

	http.HandleFunc("/todo", handleTodo)

	http.HandleFunc("/add", handleAdd) // <2>

	port := getPortNumber()
	fmt.Printf("listening port : %d\n", port)
	err := http.ListenAndServe(fmt.Sprintf(":%d", port), nil)
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}
