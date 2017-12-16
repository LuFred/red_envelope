package handle

import (
	"fmt"
	"strconv"
	"time"

	dbcore "github.com/lufred/red_envelope/Service/red_envelope_service/db/core"
	"github.com/lufred/red_envelope/Service/red_envelope_service/db/entity"
	pb "github.com/lufred/red_envelope/Service/red_envelope_service/proto/pb"
	"github.com/lufred/red_envelope/util"
	"golang.org/x/net/context"
	upper "upper.io/db.v3"
	"upper.io/db.v3/lib/sqlbuilder"
)

//GetCardBalance 查询银行卡余额
func (s *Server) GetCardBalance(ctx context.Context, in *pb.GetCardBalanceRequest) (*pb.GetCardBalanceReply, error) {
	resp := new(pb.GetCardBalanceReply)
	resp.UserId = in.GetUserId()
	ety, err := entity.QueryBankCardOne(upper.Cond{
		"user_id": in.GetUserId(),
	})
	if err != nil {
		goto ERREND
	}
	if ety != nil {
		resp.Money = ety.Money
		resp.Id = ety.ID
	}
	return resp, nil
ERREND:
	return nil, &HandleError{
		Op:  "GetCardBalance",
		Err: err,
	}
}

//CreateRede 创建红包
func (s *Server) CreateRede(ctx context.Context, in *pb.CreateRedeRequest) (*pb.CreateRedeReply, error) {
	resp := new(pb.CreateRedeReply)
	var sess sqlbuilder.Database
	var tx sqlbuilder.Tx
	var redeEntity *entity.RedEnvelopeEntity
	var _id interface{}
	var id int
	ety, err := entity.QueryBankCardOne(upper.Cond{
		"user_id": in.GetUserId(),
	})
	if err != nil {
		goto ERREND
	}
	if ety == nil || ety.Money < in.Amount {
		resp.Error = &pb.ServerError{
			ErrorCode:    101,
			ErrorMessage: "余额不足",
		}
		return resp, nil
	}
	sess, err = dbcore.GetDBSession()
	if err != nil {
		goto ERREND
	}
	tx, err = sess.NewTx(nil)
	if err != nil {
		goto ERREND
	}
	ety.Money = ety.Money - in.Amount
	err = tx.Collection(entity.BankCardTableName).UpdateReturning(ety)
	if err != nil {
		tx.Rollback()
		goto ERREND
	}
	redeEntity = &entity.RedEnvelopeEntity{
		UserID:     in.GetUserId(),
		SecretCode: in.GetSecretCode(),
		Amount:     in.GetAmount(),
		Count:      in.GetCount(),
		ExpireTime: time.Now().Add(24*time.Hour).UnixNano() / 1e6,
		GMTCreate:  util.GetTimeMillisecond(),
	}
	_id, err = tx.Collection(entity.RedEnvelopeTableName).Insert(redeEntity)
	if err != nil {
		tx.Rollback()
		goto ERREND
	}
	err = tx.Commit()
	if err != nil {
		goto ERREND
	}
	id, err = strconv.Atoi(fmt.Sprint(_id))
	resp.Id = int32(id)
	resp.Amount = redeEntity.Amount
	resp.Count = redeEntity.Count
	resp.ExpireTime = redeEntity.ExpireTime
	resp.SecretCode = redeEntity.SecretCode
	resp.UserId = redeEntity.UserID
	resp.GmtCreate = redeEntity.GMTCreate
	return resp, nil
ERREND:
	return nil, &HandleError{
		Op:  "CreateRede",
		Err: err,
	}
}

//CreateRecord 创建红包领取记录
func (s *Server) CreateRecord(ctx context.Context, in *pb.CreateRecordRequest) (*pb.CreateRecordReply, error) {
	resp := new(pb.CreateRecordReply)
	ety := &entity.ReceiveRecordEntity{
		UserID:        in.GetUserId(),
		RedEnvelopeID: in.GetRedeId(),
		Amount:        in.GetAmount(),
		GMTCreate:     util.GetTimeMillisecond(),
	}
	ety, err := entity.InsertReceiveRecord(ety)
	if err != nil {
		goto ERREND
	}
	resp.Success = true
	return resp, nil
ERREND:
	return nil, &HandleError{
		Op:  "CreateRecord",
		Err: err,
	}
}

//GetRecordByUID 根据用户id查询红包记录
func (s *Server) GetRecordByUID(ctx context.Context, in *pb.GetRecordByUIDRequest) (*pb.GetRecordByUIDReply, error) {
	resp := new(pb.GetRecordByUIDReply)
	etys, err := entity.QueryReceiveRecordsAll(upper.Cond{
		"user_id": in.GetUserId(),
	}, "-gmt_create")
	if err != nil {
		goto ERREND
	}
	resp.Data = make([]*pb.RecordData, 0)
	for i := range etys {
		resp.Data = append(resp.Data, &pb.RecordData{
			RedeId: etys[i].RedEnvelopeID,
			Amount: etys[i].Amount,
			Time:   etys[i].GMTCreate,
		})
	}
	return resp, nil
ERREND:
	return nil, &HandleError{
		Op:  "CreateRecord",
		Err: err,
	}
}

//UpdateBalanceByUID 修改用户个人余额
func (s *Server) UpdateBalanceByUID(ctx context.Context, in *pb.UpdateBalanceByUIDRequest) (*pb.UpdateBalanceByUIDReply, error) {
	resp := new(pb.UpdateBalanceByUIDReply)
	var sess sqlbuilder.Database
	var tx sqlbuilder.Tx
	var coll upper.Collection
	var balenceEty = &entity.BalanceEntity{}
	sess, err := dbcore.GetDBSession()
	if err != nil {
		goto ERREND
	}
	tx, err = sess.NewTx(nil)
	if err != nil {
		goto ERREND
	}
	coll = tx.Collection(entity.BalanceTableName)
	err = coll.Find(upper.Cond{
		"user_id": in.GetUserId(),
	}).One(balenceEty)
	if err != nil {
		if err != upper.ErrNoMoreRows {
			goto ERREND
		}
		balenceEty = &entity.BalanceEntity{
			UserID:    in.GetUserId(),
			Balance:   in.GetAmount(),
			GMTCreate: util.GetTimeMillisecond(),
		}
		err = coll.InsertReturning(balenceEty)
		if err != nil {
			tx.Rollback()
			goto ERREND
		}
	} else {
		balenceEty.Balance = balenceEty.Balance + in.GetAmount()
		balenceEty.GMTModified = util.GetTimeMillisecond()
		err = coll.UpdateReturning(balenceEty)
		if err != nil {
			tx.Rollback()
			goto ERREND
		}
	}
	err = tx.Commit()
	if err != nil {
		goto ERREND
	}
	resp.Success = true
	return resp, nil
ERREND:
	return nil, &HandleError{
		Op:  "UpdateCardBaclance",
		Err: err,
	}
}

//GetBalanceByUID 查询银行卡余额
func (s *Server) GetBalanceByUID(ctx context.Context, in *pb.GetBalanceByUIDRequest) (*pb.GetBalanceByUIDReply, error) {
	resp := new(pb.GetBalanceByUIDReply)
	resp.UserId = in.GetUserId()
	ety, err := entity.QueryBalanceOne(upper.Cond{
		"user_id": in.GetUserId(),
	})
	if err != nil {
		goto ERREND
	}
	if ety != nil {
		resp.Balance = ety.Balance
		resp.UserId = ety.UserID
	}
	return resp, nil
ERREND:
	return nil, &HandleError{
		Op:  "GetCardBalance",
		Err: err,
	}
}
