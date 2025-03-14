package main

import (
	"fmt"
	"io"
	"net/http"

	"github.com/emicklei/go-restful/v3"
	otrestful "github.com/opentracing-contrib/go-restful"
	zipkin "github.com/openzipkin-contrib/zipkin-go-opentracing"
)

func main() {
	// 1) Create a opentracing.Tracer that sends data to Zipkin
	collector, _ := zipkin.NewHTTPCollector(fmt.Sprintf("http://127.0.0.1:9411/api/v1/spans"))
	tracer, _ := zipkin.NewTracer(
		zipkin.NewRecorder(collector, true, "0.0.0.0:0", "trivial"))

	// install a global (=DefaultContainer) filter (processed before any webservice in the DefaultContainer)
	// to provide  OpenTracing instrument
	restful.Filter(otrestful.NewOTFilter(tracer))

	ws := new(restful.WebService)
	ws.Route(ws.GET("/hello").To(hello))

	restful.Add(ws)
	http.ListenAndServe(":8080", nil)
}

func hello(req *restful.Request, resp *restful.Response) {
	io.WriteString(resp, "world")
}
