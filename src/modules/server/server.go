package server

import (
	"net/http"

	log "../logging"
)

// Listen for HTTP access
func Listen(addr string) (err error) {
	log.Noticef("BFF is started on %s", addr)
	err = http.ListenAndServe(addr, NewRouter())
	return
}
