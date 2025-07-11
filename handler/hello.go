package handler

import (
	"fmt"
	"net/http"
)

// HelloHandler returns a Hello World message
func HelloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, World")
}
