package main

import (
	"flag"
	"net"

	pb "github.com/lufred/red_envelope/Service/red_envelope_service/proto/pb"

	"github.com/lufred/red_envelope/Service/red_envelope_service/config"
	"github.com/lufred/red_envelope/Service/red_envelope_service/handle"
	"github.com/lufred/red_envelope/util/log"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", config.ProConfig.Listen)
	if err != nil {
		log.Error(err)
		panic(err)
	}
	var s *grpc.Server
	s = grpc.NewServer()
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
