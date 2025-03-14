module github.com/opentracing-contrib/go-restful/example

go 1.24.1

replace github.com/opentracing-contrib/go-restful => ../

require (
	github.com/emicklei/go-restful v2.16.0+incompatible
	github.com/opentracing-contrib/go-restful v0.0.0-00010101000000-000000000000
	github.com/opentracing-contrib/go-stdlib v1.1.0
	github.com/openzipkin-contrib/zipkin-go-opentracing v0.5.0
)

require (
	github.com/json-iterator/go v1.1.12 // indirect
	github.com/modern-go/concurrent v0.0.0-20180228061459-e0a39a4cb421 // indirect
	github.com/modern-go/reflect2 v1.0.2 // indirect
	github.com/opentracing-contrib/go-observer v0.0.0-20170622124052-a52f23424492 // indirect
	github.com/opentracing/opentracing-go v1.2.0 // indirect
	github.com/openzipkin/zipkin-go v0.4.1 // indirect
	google.golang.org/grpc v1.50.0 // indirect
)
