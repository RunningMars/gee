package main

import (
	"fmt"
	"gee/gee"
	"net/http"
)

func main() {
	r := gee.New()

	r.Get("/hello", func(w http.ResponseWriter, req *http.Request) {
		fmt.Fprintf(w, "hello")
	})

	r.Run("localhost:8061")

}
