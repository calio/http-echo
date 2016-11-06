package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
)

var port int

func init() {
	flag.IntVar(&port, "port", 10101, "Listening port")
}

func hello(w http.ResponseWriter, r *http.Request) {
	//w.Header().Add(":method", r.Method)
	w.Header().Add("X-Host", r.Host)
	for header, value := range r.Header {
		for _, v1 := range value {
			w.Header().Add("X-"+header, v1)
		}
	}
	w.WriteHeader(http.StatusOK)
	io.Copy(w, r.Body)
	err := r.Body.Close()
	if err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()
	log.Printf("Listening on :%v", port)
	http.HandleFunc("/", hello)
	http.ListenAndServe(fmt.Sprintf(":%v", port), nil)
}
