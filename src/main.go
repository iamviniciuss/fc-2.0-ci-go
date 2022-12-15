package main

import (
	"fmt"
	"net/http"
)

func main() {

	fmt.Println("Starting app2")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>Hello</h1>"))
	})

	http.ListenAndServe(":80", nil)
}
