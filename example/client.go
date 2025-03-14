//go:build go1.7
// +build go1.7

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	otrestful "github.com/opentracing-contrib/go-restful"
	"github.com/opentracing-contrib/go-stdlib/nethttp"
	zipkin "github.com/openzipkin-contrib/zipkin-go-opentracing"
)

func main() {
	// 1) Create a opentracing.Tracer that sends data to Zipkin
	collector, _ := zipkin.NewHTTPCollector(fmt.Sprintf("http://127.0.0.1:9411/api/v1/spans"))
	tracer, _ := zipkin.NewTracer(
		zipkin.NewRecorder(collector, true, "0.0.0.0:0", "trivial"))

	// 2) Demonstrate nethttp client-side OpenTracing instrumentation works
	client := &http.Client{Transport: &nethttp.Transport{}}
	req, err := http.NewRequest("GET", "http://127.0.0.1:8080/hello", nil)
	if err != nil {
		log.Fatal(err)
	}

	req, ht := nethttp.TraceRequest(tracer, req,
		nethttp.OperationName(otrestful.DefaultOperationName), nethttp.ComponentName(otrestful.DefaultComponentName))

	_, err = client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	ht.Finish()

	// ... give Zipkin ample time to flush
	time.Sleep(2 * time.Second)
}
