package server

import (
	"fmt"
	"net/http"
)

func StartServer() {
	http.ListenAndServe(":8000", nil)
	fmt.Println("Server started")
}
