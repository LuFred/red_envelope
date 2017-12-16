package main

import (
	"flag"
	"net"

	pb "github.com/lufred/red_envelope/Service/red_envelope_service/proto/pb"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_opentracing "github.com/grpc-ecosystem/go-grpc-middleware/tracing/opentracing"
	"github.com/lufred/red_envelope/Service/red_envelope_service/config"
	"github.com/lufred/red_envelope/Service/red_envelope_service/handle"
	"github.com/lufred/red_envelope/util/log"
	jaeger "github.com/uber/jaeger-client-go"
	"github.com/uber/jaeger-client-go/transport/zipkin"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", config.ProConfig.Listen)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	transport, err := zipkin.NewHTTPTransport(
		config.ProConfig.TracingTransportURL,
		zipkin.HTTPBatchSize(1),
		zipkin.HTTPLogger(jaeger.StdLogger),
	)
	var s *grpc.Server
	if err != nil {
		log.Errorf("Cannot initialize opentracing HTTP transport: %v", err)
		s = grpc.NewServer()
	} else {
		tracer, closer := jaeger.NewTracer(
			config.ProConfig.ServiceName,
			jaeger.NewConstSampler(true),
			jaeger.NewRemoteReporter(transport),
		)
		defer closer.Close()
		var opts []grpc_opentracing.Option
		opts = append(opts, grpc_opentracing.WithTracer(tracer))
		s = grpc.NewServer(grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_opentracing.UnaryServerInterceptor(opts...),
			)))
	}
	pb.RegisterRedEnvelopeServer(s, &handle.Server{})
	log.Infof("listen %s", config.ProConfig.Listen)
	if err := s.Serve(lis); err != nil {
		log.Errorf("failed to serve: %v", err)
	}
}

func init() {
	flag.Parse()
	config.RegisterConfig()
	log.Debugged = config.ProConfig.Debug
}
