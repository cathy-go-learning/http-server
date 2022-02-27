package main

import (
	"github.com/cathy-go-learning/http-server/pkg/envflag"
	"io"
	"log"
	"net/http"
	"time"
)

func main() {
	envflag.Parse()

	http.HandleFunc("/home", home)
	http.HandleFunc("/healthz", healthz)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthz(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(200)
	_, err := io.WriteString(writer, "hello server!")
	if err != nil {
		log.Fatal(err)
	}
}

func home(writer http.ResponseWriter, request *http.Request) {
	for k, values := range request.Header {
		log.Printf("%s:%s", k, values)
		for _, v := range values {
			writer.Header().Add(k, v)
		}
	}
	//writer.Header().Add("VERSION", *envflag.Version)
	_, err := io.WriteString(writer, "hello server!")
	if err != nil {
		log.Fatal(err)
	}
}
