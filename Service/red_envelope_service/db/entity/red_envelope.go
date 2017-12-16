package entity

import (
	"fmt"
	"log"
	"strconv"

	dbcore "github.com/lufred/red_envelope/Service/red_envelope_service/db/core"
	upper "upper.io/db.v3"
)

//RedEnvelopeEntity 红包实体
type RedEnvelopeEntity struct {
	ID          int32  `db:"id" json:"db"`
	UserID      int32  `db:"user_id" json:"user_id"`
	SecretCode  string `db:"secret_code" json:"secret_code"`
	Amount      int32  `db:"amount" json:"amount"`
	Count       int32  `db:"count" json:"count"`
	ExpireTime  int64  `db:"expire_time" json:"expire_time"`   //过期时间
	GMTCreate   int64  `db:"gmt_create" json:"gmt_create"`     //创建时间
	GMTModified int64  `db:"gmt_modified" json:"gmt_modified"` //修改时间
}

var RedEnvelopeTableName = "red_envelope"

//InsertRedEnvelope insert RedEnvelope
func InsertRedEnvelope(ety *RedEnvelopeEntity) (*RedEnvelopeEntity, error) {
	var dbError *dbcore.DbError
	var id int
	var sid string
	var _id interface{}
	sess, err := dbcore.GetDBSession()
	log.Println(sess)
	if err != nil {
		goto ERREND
	}
	_id, err = sess.Collection(RedEnvelopeTableName).Insert(ety)
	if err != nil {
		goto ERREND
	}
	sid = fmt.Sprint(_id)
	id, err = strconv.Atoi(sid)
	ety.ID = int32(id)
	return ety, nil
ERREND:
	dbError = &dbcore.DbError{
		Op:     "InsertRedEnvelope",
		Entity: RedEnvelopeTableName,
		Err:    err,
	}
	return nil, dbError
}

//QueryRedEnvelopeOne query a single record
func QueryRedEnvelopeOne(selecter upper.Cond) (*RedEnvelopeEntity, error) {
	var dbError *dbcore.DbError
	ety := &RedEnvelopeEntity{}
	sess, err := dbcore.GetDBSession()
	if err != nil {
		goto ERREND
	}
	err = sess.Collection(RedEnvelopeTableName).Find(selecter).One(ety)
	if err != nil {
		if err == upper.ErrNoMoreRows {
			return nil, nil
		}
		goto ERREND
	}
	return ety, nil
ERREND:
	dbError = &dbcore.DbError{
		Op:     "QueryRedEnvelopeOne",
		Entity: RedEnvelopeTableName,
		Err:    err,
	}
	return nil, dbError
}

//DeleteRedEnvelope delete RedEnvelope
func DeleteRedEnvelope(id int32) error {
	sess, err := dbcore.GetDBSession()
	if err != nil {
		goto ERREND
	}
	err = sess.Collection(RedEnvelopeTableName).Find(upper.Cond{
		"id": id,
	}).Delete()
	if err != nil {
		goto ERREND
	}
	return nil
ERREND:
	var dbError *dbcore.DbError
	dbError = &dbcore.DbError{
		Op:     "DeleteRedEnvelope",
		Entity: RedEnvelopeTableName,
		Err:    err,
	}
	return dbError
}

//QueryRedEnvelopesAll query RedEnvelope
func QueryRedEnvelopesAll(selecter upper.Cond, order ...interface{}) ([]RedEnvelopeEntity, error) {
	sess, err := dbcore.GetDBSession()
	var RedEnvelopees []RedEnvelopeEntity
	if err != nil {
		goto ERREND
	}
	err = sess.Collection(RedEnvelopeTableName).Find(selecter).OrderBy(order...).All(&RedEnvelopees)
	if err != nil {
		goto ERREND
	}
	return RedEnvelopees, nil
ERREND:
	var dbError *dbcore.DbError
	dbError = &dbcore.DbError{
		Op:     "QueryRedEnvelopesAll",
		Entity: RedEnvelopeTableName,
		Err:    err,
	}
	return nil, dbError
}

//QueryRedEnvelopes query RedEnvelope
func QueryRedEnvelopes(selecter upper.Cond, offset, limit int, order ...interface{}) ([]RedEnvelopeEntity, error) {
	sess, err := dbcore.GetDBSession()
	var RedEnvelopees []RedEnvelopeEntity
	if err != nil {
		goto ERREND
	}

	err = sess.Collection(RedEnvelopeTableName).Find(selecter).OrderBy(order...).Offset(offset).Limit(limit).All(&RedEnvelopees)
	if err != nil {
		goto ERREND
	}
	return RedEnvelopees, nil
ERREND:
	var dbError *dbcore.DbError
	dbError = &dbcore.DbError{
		Op:     "QueryRedEnvelopes",
		Entity: RedEnvelopeTableName,
		Err:    err,
	}
	return nil, dbError
}

//QueryRedEnvelopeCount query RedEnvelope count
func QueryRedEnvelopeCount(selecter upper.Cond) (uint64, error) {
	sess, err := dbcore.GetDBSession()
	var c uint64
	if err != nil {
		goto ERREND
	}

	c, err = sess.Collection(RedEnvelopeTableName).Find(selecter).Count()
	if err != nil {
		goto ERREND
	}
	return c, nil
ERREND:
	var dbError *dbcore.DbError
	dbError = &dbcore.DbError{
		Op:     "QueryRedEnvelopeCount",
		Entity: RedEnvelopeTableName,
		Err:    err,
	}
	return 0, dbError
}

//UpdateRedEnvelope update RedEnvelope
func UpdateRedEnvelope(ety *RedEnvelopeEntity) error {
	var dbError *dbcore.DbError
	sess, err := dbcore.GetDBSession()
	if err != nil {
		goto ERREND
	}
	err = sess.Collection(RedEnvelopeTableName).UpdateReturning(ety)
	if err != nil {
		goto ERREND
	}
	return nil
ERREND:
	dbError = &dbcore.DbError{
		Op:     "UpdateRedEnvelope",
		Entity: RedEnvelopeTableName,
		Err:    err,
	}
	return dbError

}
