package server

import (
	"fmt"
	"net/http"
	"time"

	log "../logging"
)

// Router handle all of request
type Router struct {
}

// NewRouter will create a Router instance
func NewRouter() (router *Router) {
	router = &Router{}
	return
}

func (router *Router) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	start := time.Now().UnixNano()
	router.route(res, req)
	log.Debugf(
		"用时: %d 毫秒 ( %d 微秒 )\n  URI: %s\n",
		((time.Now().UnixNano() - start) / 1e6),
		((time.Now().UnixNano() - start) / 1e3),
		req.URL.Path,
	)
}

func (router *Router) route(res http.ResponseWriter, req *http.Request) {
	var result string
	var err error
	result = "Hello World !"
	if nil != err {
		res.WriteHeader(500)
		result = "Internal Server Error"
		log.Error(err)
	}
	fmt.Fprintln(res, result)
}
