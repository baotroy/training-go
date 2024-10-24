package main

import "net/http"

func handlerErr(w http.ResponseWriter, r *http.Request) {
	resonseWithError(w, http.StatusBadRequest, "something went wrong")
}
