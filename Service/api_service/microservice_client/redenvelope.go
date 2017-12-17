package microservice_client

import (
	"github.com/lufred/red_envelope/Service/api_service/config"
	redepb "github.com/lufred/red_envelope/Service/red_envelope_service/proto/pb"
	"github.com/lufred/red_envelope/util/log"
	"google.golang.org/grpc"
)

//registerRedEnvelopeClientConn
func registerRedEnvelopeClientConn() {
	var err error
	RedeClientConn, err = grpc.Dial(
		config.ProConfig.Microservice.RedEnvelopeHost,
		grpc.WithInsecure(),
	)
	if err != nil {
		log.Errorf("did not connect red_envelope service host: %v", err)
	}
}

//GetRedEnvelopeServiceClient
func (m *MicroService) GetRedEnvelopeServiceClient() (redepb.RedEnvelopeClient, error) {
	return redepb.NewRedEnvelopeClient(RedeClientConn), nil
}
