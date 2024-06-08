package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{}

func main() {
	fmt.Println("Hello World!")

	defineRoutes(upgrader)

	http.ListenAndServe(":8080", nil)
}
