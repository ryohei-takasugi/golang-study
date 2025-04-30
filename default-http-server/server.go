package main

import (
	"default-http-server/handler"
	"fmt"
	"net/http"
)

func fuga(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "fuga")
}

func main() {
	http.Handle("/echo", handler.EchoHandler{})
	http.Handle("/weather", handler.WeatherEchoHandler{})

	http.HandleFunc("/fuga", fuga)

	fmt.Println("Starting server on :8080")
	fmt.Println(http.ListenAndServe(":8080", nil))
}
