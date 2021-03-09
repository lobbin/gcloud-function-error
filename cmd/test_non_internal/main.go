package test_non_internal

import (
	"net/http"

	"github.com/lobbin/gcloud-function-error/lib/lib"
)


func Test(w http.ResponseWriter, r *http.Request) {
	lib.LibTest(w, r)
}
