package server

import (
	"fmt"
	"net/http"
)

func GetFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World")
}
