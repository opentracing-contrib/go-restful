package otrestful

import (
	"github.com/emicklei/go-restful"
	"github.com/opentracing/opentracing-go"
	"github.com/opentracing/opentracing-go/ext"
)

// go-restful FilterFunction does not allow to opt in third-party parameters,
// so using global variables for the moment
// TODO(mqliang): using parameter instead of global variable to set this
var (
	Tracer            opentracing.Tracer
	OperationNameFunc func(r *restful.Request) string
	ComponentName     string
)

var (
	DefaulOperationNameFunc = func(r *restful.Request) string {
		return "HTTP " + r.Request.Method
	}
)

// OTFilter if a filter which add OpenTracing instrument
// "filter" in go-restful is similar with "middleware" mechanism in modern web framework
func OTFilter(req *restful.Request, resp *restful.Response, chain *restful.FilterChain) {
	defer chain.ProcessFilter(req, resp)
	if Tracer == nil {
		// if Tracer global variable is nil, skip tracing
		return
	}
	ctx, _ := Tracer.Extract(opentracing.HTTPHeaders, opentracing.HTTPHeadersCarrier(req.Request.Header))
	if OperationNameFunc == nil {
		OperationNameFunc = DefaulOperationNameFunc
	}
	// record operation name
	sp := Tracer.StartSpan(OperationNameFunc(req), ext.RPCServerOption(ctx))
	// record HTTP method
	ext.HTTPMethod.Set(sp, req.Request.Method)
	// record HTTP url
	ext.HTTPUrl.Set(sp, req.Request.URL.String())
	// record component name
	if ComponentName == "" {
		ComponentName = DefaultComponentName
	}
	ext.Component.Set(sp, ComponentName)
	// record HTTP status code
	ext.HTTPStatusCode.Set(sp, uint16(resp.StatusCode()))
	sp.Finish()
}
