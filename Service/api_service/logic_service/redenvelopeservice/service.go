package redenvelopeservice

import (
	lctx "github.com/lufred/red_envelope/Service/api_service/core/context"
	logiccore "github.com/lufred/red_envelope/Service/api_service/logic_service/core"
	microclient "github.com/lufred/red_envelope/Service/api_service/microservice_client"
	redemodel "github.com/lufred/red_envelope/Service/api_service/model/redenvelope"
	redepb "github.com/lufred/red_envelope/Service/red_envelope_service/proto/pb"
	"github.com/lufred/red_envelope/util/log"
	"golang.org/x/net/context"
)

//Create 创建红包
func Create(ctx context.Context, user *lctx.User, dto *redemodel.CreateRequest) (*redemodel.CreateResponse, error) {
	var getCardBalanceReply *redepb.GetCardBalanceReply
	var secretCode string
	var createReply *redepb.CreateRedeReply
	resp := new(redemodel.CreateResponse)
	micServiceClients := microclient.MicroService{}
	redeClient, err := micServiceClients.GetRedEnvelopeServiceClient()
	if err != nil {
		goto ERREND
	}
	getCardBalanceReply, err = redeClient.GetCardBalance(ctx, &redepb.GetCardBalanceRequest{
		UserId: user.UserID,
	})
	if err != nil {
		goto ERREND
	}
	if getCardBalanceReply.Money < dto.Amount {
		err = &logiccore.LogicError{
			Code:    101,
			Message: "余额不足",
			Type:    1,
		}
		return nil, err
	}
	//todo 获取8位口令
	secretCode = "xxxx"
	createReply, err = redeClient.CreateRede(ctx, &redepb.CreateRedeRequest{
		UserId:     user.UserID,
		SecretCode: secretCode,
		Amount:     dto.Amount,
		Count:      dto.Count,
	})
	if err != nil {
		goto ERREND
	}

	return resp, nil
ERREND:
	err = &logiccore.LogicError{
		SeriviceName: serviceName,
		Op:           "Create",
		Err:          err,
	}
	log.Error(err)
	return nil, err
}

//Take 抢红包
func Take(ctx context.Context, user *lctx.User, dto *redemodel.TakeRequest) (*redemodel.TakeResponse, error) {
	resp := new(redemodel.TakeResponse)
	//var err error

	return resp, nil
	// ERREND:
	// 	err = &logiccore.LogicError{
	// 		SeriviceName: serviceName,
	// 		Op:           "Create",
	// 		Err:          err,
	// 	}
	// 	log.Error(err)
	// 	return nil, err
}

//GetBalance 查询余额
func GetBalance(ctx context.Context, user *lctx.User, dto *redemodel.GetBalanceRequest) (*redemodel.GetBalanceResponse, error) {
	resp := new(redemodel.GetBalanceResponse)
	//var err error

	return resp, nil
	// ERREND:
	// 	err = &logiccore.LogicError{
	// 		SeriviceName: serviceName,
	// 		Op:           "Create",
	// 		Err:          err,
	// 	}
	// 	log.Error(err)
	// 	return nil, err
}

//GetRecord 查询红包记录
func GetRecord(ctx context.Context, user *lctx.User, dto *redemodel.GetRecordRequest) (*redemodel.GetRecordResponse, error) {
	resp := new(redemodel.GetRecordResponse)
	//var err error

	return resp, nil
	// ERREND:
	// 	err = &logiccore.LogicError{
	// 		SeriviceName: serviceName,
	// 		Op:           "Create",
	// 		Err:          err,
	// 	}
	// 	log.Error(err)
	// 	return nil, err
}

//insertRedEToRedis 添加新红包到redis
func insertRedEToRedis(reply *redepb.CreateRedeReply) error {

}
