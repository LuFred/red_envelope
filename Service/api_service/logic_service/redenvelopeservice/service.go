package redenvelopeservice

import (
	"encoding/json"
	"strconv"

	"github.com/go-redis/redis"
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
	secretCode = GetSecretCode()
	createReply, err = redeClient.CreateRede(ctx, &redepb.CreateRedeRequest{
		UserId:     user.UserID,
		SecretCode: secretCode,
		Amount:     dto.Amount,
		Count:      dto.Count,
	})
	if err != nil {
		goto ERREND
	}
	log.Error("--", createReply)
	if createReply.Error != nil {

		//存在错误
		return resp, nil
	}
	go func() {
		//红包存入redis
		rede := &Rede{
			ID:         createReply.Id,
			Amount:     createReply.Amount,
			Count:      createReply.Count,
			SecretCode: createReply.SecretCode,
			Expire:     createReply.ExpireTime,
		}
		err := insertRedEToRedis(rede)
		log.Error("redis err=", err)
	}()
	resp.Amount = createReply.Amount
	resp.Count = createReply.Count
	resp.ID = createReply.Id
	resp.SecretCode = createReply.SecretCode
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
	var rede *Rede    //redis红包对象
	var am int32      //红包随机获取金额
	var redebs []byte //第一次查询红包信息
	redisClient := GetRedisClient()
	micServiceClients := microclient.MicroService{}
	redeClient, err := micServiceClients.GetRedEnvelopeServiceClient()
	if err != nil {
		goto ERREND
	}

	redebs, err = redisClient.Get(strconv.Itoa(int(dto.RedeID))).Bytes()
	if err != nil {
		if err == redis.Nil {
			err = &logiccore.LogicError{
				Code:    100,
				Message: "红包不存在或已过期",
				Type:    1,
			}
			return nil, err
		}
		goto ERREND
	}
	rede = &Rede{}
	json.Unmarshal(redebs, rede)
	if rede.Count > 0 && rede.SecretCode == dto.SecretCode {
		if Lock(rede.ID) {
			rebytes, err := redisClient.Get(strconv.Itoa(int(dto.RedeID))).Bytes()
			if err != nil {
				UnLock(rede.ID)
				if err == redis.Nil {
					err = &logiccore.LogicError{
						Code:    100,
						Message: "红包不存在或已过期",
						Type:    1,
					}
					return nil, err
				}
				goto ERREND
			}
			r := &Rede{}
			json.Unmarshal(rebytes, r)
			if r.Count > 0 {
				resu, err := InsertUserToGroup(r.ID, user.UserID)
				if err != nil {
					UnLock(r.ID)
					goto ERREND
				}
				if resu > 0 {
					am = GetRandAmount(rede.Amount, rede.Count)
					r.Count = r.Count - 1
					r.Amount = r.Amount - am
					insertRedEToRedis(r)
					go redeClient.CreateRecord(ctx, &redepb.CreateRecordRequest{
						UserId: user.UserID,
						RedeId: r.ID,
						Amount: am,
					})
					go redeClient.UpdateBalanceByUID(ctx, &redepb.UpdateBalanceByUIDRequest{
						UserId: user.UserID,
						Amount: am,
					})
				} else {
					UnLock(r.ID)
					err = &logiccore.LogicError{
						Code:    101,
						Message: "不可重复领取同一个红包",
						Type:    1,
					}
					return nil, err
				}
			}
			UnLock(r.ID)
			resp.Amount = am
			return resp, nil
		}
		err = &logiccore.LogicError{
			Code:    200,
			Message: "系统繁忙",
			Type:    1,
		}
		return nil, err
	}
	if rede.SecretCode != dto.SecretCode {
		err = &logiccore.LogicError{
			Code:    103,
			Message: "口令错误",
			Type:    1,
		}
		return nil, err
	}
	err = &logiccore.LogicError{
		Code:    102,
		Message: "红包已领完",
		Type:    1,
	}
	return nil, err
ERREND:
	err = &logiccore.LogicError{
		SeriviceName: serviceName,
		Op:           "Create",
		Err:          err,
	}
	log.Error(err)
	return nil, err
}

//GetBalance 查询余额
func GetBalance(ctx context.Context, user *lctx.User, dto *redemodel.GetBalanceRequest) (*redemodel.GetBalanceResponse, error) {
	var reply *redepb.GetBalanceByUIDReply
	resp := new(redemodel.GetBalanceResponse)
	micServiceClients := microclient.MicroService{}
	redeClient, err := micServiceClients.GetRedEnvelopeServiceClient()
	if err != nil {
		goto ERREND
	}
	reply, err = redeClient.GetBalanceByUID(ctx, &redepb.GetBalanceByUIDRequest{
		UserId: user.UserID,
	})
	if err != nil {
		goto ERREND
	}
	resp.Balance = reply.Balance
	return resp, nil
ERREND:
	err = &logiccore.LogicError{
		SeriviceName: serviceName,
		Op:           "GetBalance",
		Err:          err,
	}
	log.Error(err)
	return nil, err
}

//GetRecord 查询红包记录
func GetRecord(ctx context.Context, user *lctx.User, dto *redemodel.GetRecordRequest) (*redemodel.GetRecordResponse, error) {
	var reply *redepb.GetRecordByUIDReply
	resp := new(redemodel.GetRecordResponse)
	micServiceClients := microclient.MicroService{}
	redeClient, err := micServiceClients.GetRedEnvelopeServiceClient()
	if err != nil {
		goto ERREND
	}
	reply, err = redeClient.GetRecordByUID(ctx, &redepb.GetRecordByUIDRequest{
		UserId: user.UserID,
	})
	if err != nil {
		goto ERREND
	}
	resp.Data = make([]*redemodel.Record, 0)
	for i := range reply.Data {
		resp.Data = append(resp.Data, &redemodel.Record{
			RedeID: reply.Data[i].RedeId,
			Amount: reply.Data[i].Amount,
			Time:   reply.Data[i].Time,
		})
	}
	return resp, nil
ERREND:
	err = &logiccore.LogicError{
		SeriviceName: serviceName,
		Op:           "GetRecord",
		Err:          err,
	}
	log.Error(err)
	return nil, err
}
