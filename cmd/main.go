package main

import (
	"io"
	"log"
	"net/http"
	//envflag "github.com/cathy-go-learning/http-server/pkg/envflag/flags.go"
)

func main() {
	http.HandleFunc("/home", home)
	err := http.ListenAndServe(":8080", nil)
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
	//writer.Header().Add("VERSION", envflag.)
	_, err := io.WriteString(writer, "hello server!")
	if err != nil {
		log.Fatal(err)
	}
}
