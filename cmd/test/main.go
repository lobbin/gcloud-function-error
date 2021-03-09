package test

import (
	"net/http"

	"github.com/lobbin/gcloud-function-error/internal/lib"
)


func Test(w http.ResponseWriter, r *http.Request) {
	lib.LibTest(w, r)
}
