package main

import (
	"default-http-server/router"
	"fmt"
	"net/http"
)

func main() {
	fmt.Println("Starting server on :8080")
	fmt.Println(http.ListenAndServe(":8080", router.CreateRouter()))
}
