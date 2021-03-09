package lib

import (
	"fmt"
	"net/http"
)

func LibTest(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}