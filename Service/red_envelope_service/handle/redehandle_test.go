package handle

import (
	"flag"
	"testing"

	"github.com/lufred/red_envelope/Service/red_envelope_service/config"
	pb "github.com/lufred/red_envelope/Service/red_envelope_service/proto/pb"
	"golang.org/x/net/context"
)

var (
	server *Server
)

func init() {
	flag.Parse()
	config.RegisterConfig()
	server = new(Server)
}
func TestCreateClass(t *testing.T) {
	reply, err := server.GetCardBalance(context.Background(), &pb.GetCardBalanceRequest{
		UserId: 333,
	})
	t.Logf("%v,%v", reply, err)
}

func TestCreateRede(t *testing.T) {
	reply, err := server.CreateRede(context.Background(), &pb.CreateRedeRequest{
		UserId:     333,
		Amount:     17,
		Count:      2,
		SecretCode: "xx123cSS",
	})
	t.Logf("%v,%v", reply, err)
}

func TestCreateRecord(t *testing.T) {
	reply, err := server.CreateRecord(context.Background(), &pb.CreateRecordRequest{
		UserId: 333,
		Amount: 17,
		RedeId: 2,
	})
	t.Logf("%v,%v", reply, err)
}

func TestGetRecordByUID(t *testing.T) {
	reply, err := server.GetRecordByUID(context.Background(), &pb.GetRecordByUIDRequest{
		UserId: 313,
	})
	t.Logf("%v,%v", reply, err)
}
func TestUpdateBalanceByUID(t *testing.T) {
	reply, err := server.UpdateBalanceByUID(context.Background(), &pb.UpdateBalanceByUIDRequest{
		UserId: 313,
		Amount: 300,
	})
	t.Logf("%v,%v", reply, err)
}
func TestGetBalanceByUID(t *testing.T) {
	reply, err := server.GetBalanceByUID(context.Background(), &pb.GetBalanceByUIDRequest{
		UserId: 313,
	})
	t.Logf("%v,%v", reply, err)
}
