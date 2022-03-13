package main

import (
	"fmt"
	"github.com/cathy-go-learning/http-server/metrics"
	"github.com/cathy-go-learning/http-server/pkg/envflag"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"io"
	"log"
	"math/rand"
	"net/http"
	"time"
)

func main() {
	envflag.Parse()

	metrics.Register()
	mux := http.NewServeMux()

	mux.HandleFunc("/", home)
	mux.HandleFunc("/healthz", healthz)
	mux.HandleFunc("/latency", latency)
	mux.Handle("/metrics", promhttp.Handler())
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

func latency(writer http.ResponseWriter, request *http.Request) {
	timer := metrics.NewTimer()
	defer timer.ObserveTotal()
	randInt := rand.Intn(2000)
	time.Sleep(time.Millisecond * time.Duration(randInt))
	_, err := writer.Write([]byte(fmt.Sprintf("<h1>%d<h1>", randInt)))
	if err != nil {
		log.Fatal(err)
	}
}
