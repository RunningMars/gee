package gee

import (
	"fmt"
	"net/http"
)

type HandleFunc func(http.ResponseWriter, *http.Request)

type Engine struct {
	routerMap map[string]HandleFunc
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := e.routerMap[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

func New() *Engine {
	return &Engine{routerMap: make(map[string]HandleFunc)}
}

func (e *Engine) addRouter(method string, pattern string, f HandleFunc) {
	key := method + "-" + pattern
	e.routerMap[key] = f
}

func (e *Engine) Get(pattern string, f HandleFunc) {
	e.addRouter("GET", pattern, f)
}

func (e *Engine) Post(pattern string, f HandleFunc) {
	e.addRouter("POST", pattern, f)
}

// Run defines the method to start a http server
func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}
