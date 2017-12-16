package microservice_client

import "google.golang.org/grpc"

var (
	RedeClientConn *grpc.ClientConn
)

func RegisterMicroService() {
	//实例化red_envelope服务连接对象
	registerRedEnvelopeClientConn()

}