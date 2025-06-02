module github.com/opentracing-contrib/go-restful/example

go 1.24.1

replace github.com/opentracing-contrib/go-restful => ../

require (
	github.com/emicklei/go-restful/v3 v3.12.2
	github.com/opentracing-contrib/go-restful v0.0.0-20250526164404-dca56b14add6
	github.com/opentracing-contrib/go-stdlib v1.1.0
	github.com/openzipkin-contrib/zipkin-go-opentracing v0.5.0
)

require (
	github.com/opentracing-contrib/go-observer v0.0.0-20170622124052-a52f23424492 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/openzipkin/zipkin-go v0.4.1 // indirect
	google.golang.org/grpc v1.56.3 // indirect
)
