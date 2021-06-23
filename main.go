package main

import (
	"fmt"
	"gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(w http.ResponseWriter, req *http.Request){
		fmt.Fprintf(w, "URL.Path=%q\n", req.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, req *http.Request){
		fmt.Fprintf(w, "Header[%q]=%q\n", req.Header)
	})

	r.Run(":9999")
}
