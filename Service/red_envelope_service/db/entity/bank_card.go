package entity

import (
	"fmt"
	"log"
	"strconv"

	dbcore "github.com/lufred/red_envelope/Service/red_envelope_service/db/core"
	upper "upper.io/db.v3"
)

//BankCardEntity 银行卡实体
type BankCardEntity struct {
	ID          int32 `db:"id" json:"db"`
	UserID      int32 `db:"user_id" json:"user_id"`
	Money       int32 `db:"money" json:"money"`
	GMTCreate   int64 `db:"gmt_create" json:"gmt_create"`     //创建时间
	GMTModified int64 `db:"gmt_modified" json:"gmt_modified"` //修改时间
}

var BankCardTableName = "bank_card"

//InsertBankCard insert BankCard
func InsertBankCard(ety *BankCardEntity) (*BankCardEntity, error) {
	var dbError *dbcore.DbError
	var id int
	var sid string
	var _id interface{}
	sess, err := dbcore.GetDBSession()
	log.Println(sess)
	if err != nil {
		goto ERREND
	}
	_id, err = sess.Collection(BankCardTableName).Insert(ety)
	if err != nil {
		goto ERREND
	}
	sid = fmt.Sprint(_id)
	id, err = strconv.Atoi(sid)
	ety.ID = int32(id)
	return ety, nil
ERREND:
	dbError = &dbcore.DbError{
		Op:     "InsertBankCard",
		Entity: BankCardTableName,
		Err:    err,
	}
	return nil, dbError
}

//QueryBankCardOne query a single record
func QueryBankCardOne(selecter upper.Cond) (*BankCardEntity, error) {
	var dbError *dbcore.DbError
	ety := &BankCardEntity{}
	sess, err := dbcore.GetDBSession()
	if err != nil {
		goto ERREND
	}
	err = sess.Collection(BankCardTableName).Find(selecter).One(ety)
	if err != nil {
		if err == upper.ErrNoMoreRows {
			return nil, nil
		}
		goto ERREND
	}
	return ety, nil
ERREND:
	dbError = &dbcore.DbError{
		Op:     "QueryBankCardOne",
		Entity: BankCardTableName,
		Err:    err,
	}
	return nil, dbError
}

//DeleteBankCard delete BankCard
func DeleteBankCard(id int32) error {
	sess, err := dbcore.GetDBSession()
	if err != nil {
		goto ERREND
	}
	err = sess.Collection(BankCardTableName).Find(upper.Cond{
		"id": id,
	}).Delete()
	if err != nil {
		goto ERREND
	}
	return nil
ERREND:
	var dbError *dbcore.DbError
	dbError = &dbcore.DbError{
		Op:     "DeleteBankCard",
		Entity: BankCardTableName,
		Err:    err,
	}
	return dbError
}

//QueryBankCardsAll query BankCard
func QueryBankCardsAll(selecter upper.Cond, order ...interface{}) ([]BankCardEntity, error) {
	sess, err := dbcore.GetDBSession()
	var BankCardes []BankCardEntity
	if err != nil {
		goto ERREND
	}
	err = sess.Collection(BankCardTableName).Find(selecter).OrderBy(order...).All(&BankCardes)
	if err != nil {
		goto ERREND
	}
	return BankCardes, nil
ERREND:
	var dbError *dbcore.DbError
	dbError = &dbcore.DbError{
		Op:     "QueryBankCardsAll",
		Entity: BankCardTableName,
		Err:    err,
	}
	return nil, dbError
}

//QueryBankCards query BankCard
func QueryBankCards(selecter upper.Cond, offset, limit int, order ...interface{}) ([]BankCardEntity, error) {
	sess, err := dbcore.GetDBSession()
	var BankCardes []BankCardEntity
	if err != nil {
		goto ERREND
	}

	err = sess.Collection(BankCardTableName).Find(selecter).OrderBy(order...).Offset(offset).Limit(limit).All(&BankCardes)
	if err != nil {
		goto ERREND
	}
	return BankCardes, nil
ERREND:
	var dbError *dbcore.DbError
	dbError = &dbcore.DbError{
		Op:     "QueryBankCards",
		Entity: BankCardTableName,
		Err:    err,
	}
	return nil, dbError
}

//QueryBankCardCount query BankCard count
func QueryBankCardCount(selecter upper.Cond) (uint64, error) {
	sess, err := dbcore.GetDBSession()
	var c uint64
	if err != nil {
		goto ERREND
	}

	c, err = sess.Collection(BankCardTableName).Find(selecter).Count()
	if err != nil {
		goto ERREND
	}
	return c, nil
ERREND:
	var dbError *dbcore.DbError
	dbError = &dbcore.DbError{
		Op:     "QueryBankCardCount",
		Entity: BankCardTableName,
		Err:    err,
	}
	return 0, dbError
}

//UpdateBankCard update BankCard
func UpdateBankCard(ety *BankCardEntity) error {
	var dbError *dbcore.DbError
	sess, err := dbcore.GetDBSession()
	if err != nil {
		goto ERREND
	}
	err = sess.Collection(BankCardTableName).UpdateReturning(ety)
	if err != nil {
		goto ERREND
	}
	return nil
ERREND:
	dbError = &dbcore.DbError{
		Op:     "UpdateBankCard",
		Entity: BankCardTableName,
		Err:    err,
	}
	return dbError

}
