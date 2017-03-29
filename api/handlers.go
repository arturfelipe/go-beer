package api

import (
	"fmt"
	"net/http"
)

type ListBeerHandler struct{}

func (h *ListBeerHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "OK")
}
