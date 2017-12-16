package opentracing

import (
	"io"

	"github.com/lufred/red_envelope/Service/api_service/config"

	"github.com/lufred/red_envelope/util/log"
	opentracing "github.com/opentracing/opentracing-go"
	jaeger "github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/transport/zipkin"
)

var Tracer opentracing.Tracer
var tracerIO io.Closer

func RegisterJadgerTracer() {
	transport, err := zipkin.NewHTTPTransport(
		config.ProConfig.TracingTransportURL,
		zipkin.HTTPBatchSize(1),
		zipkin.HTTPLogger(jaeger.StdLogger),
	)
	if err != nil {
		log.Errorf("Cannot initialize HTTP transport: %v", err)
	}
	// create Jaeger tracer
	Tracer, tracerIO = jaeger.NewTracer(
		config.ProConfig.ServiceName,
		jaeger.NewConstSampler(true), // sample all traces
		jaeger.NewRemoteReporter(transport),
	)

}

//Close close tracer
func Close() {
	if tracerIO != nil {
		tracerIO.Close()
	}
}
