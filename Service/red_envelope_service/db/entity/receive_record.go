package entity

import (
	"fmt"
	"log"
	"strconv"

	dbcore "github.com/lufred/red_envelope/Service/red_envelope_service/db/core"
	upper "upper.io/db.v3"
)

//ReceiveRecordEntity 领取记录实体
type ReceiveRecordEntity struct {
	ID            int32 `db:"id" json:"db"`
	UserID        int32 `db:"user_id" json:"user_id"`
	RedEnvelopeID int32 `db:"red_envelope_id" json:"red_envelope_id"`
	Amount        int32 `db:"amount" json:"amount"`
	GMTCreate     int64 `db:"gmt_create" json:"gmt_create"`     //创建时间
	GMTModified   int64 `db:"gmt_modified" json:"gmt_modified"` //修改时间
}

var ReceiveRecordTableName = "receive_record"

//InsertReceiveRecord insert ReceiveRecord
func InsertReceiveRecord(ety *ReceiveRecordEntity) (*ReceiveRecordEntity, error) {
	var dbError *dbcore.DbError
	var id int
	var sid string
	var _id interface{}
	sess, err := dbcore.GetDBSession()
	log.Println(sess)
	if err != nil {
		goto ERREND
	}
	_id, err = sess.Collection(ReceiveRecordTableName).Insert(ety)
	if err != nil {
		goto ERREND
	}
	sid = fmt.Sprint(_id)
	id, err = strconv.Atoi(sid)
	ety.ID = int32(id)
	return ety, nil
ERREND:
	dbError = &dbcore.DbError{
		Op:     "InsertReceiveRecord",
		Entity: ReceiveRecordTableName,
		Err:    err,
	}
	return nil, dbError
}

//QueryReceiveRecordOne query a single record
func QueryReceiveRecordOne(selecter upper.Cond) (*ReceiveRecordEntity, error) {
	var dbError *dbcore.DbError
	ety := &ReceiveRecordEntity{}
	sess, err := dbcore.GetDBSession()
	if err != nil {
		goto ERREND
	}
	err = sess.Collection(ReceiveRecordTableName).Find(selecter).One(ety)
	if err != nil {
		if err == upper.ErrNoMoreRows {
			return nil, nil
		}
		goto ERREND
	}
	return ety, nil
ERREND:
	dbError = &dbcore.DbError{
		Op:     "QueryReceiveRecordOne",
		Entity: ReceiveRecordTableName,
		Err:    err,
	}
	return nil, dbError
}

//DeleteReceiveRecord delete ReceiveRecord
func DeleteReceiveRecord(id int32) error {
	sess, err := dbcore.GetDBSession()
	if err != nil {
		goto ERREND
	}
	err = sess.Collection(ReceiveRecordTableName).Find(upper.Cond{
		"id": id,
	}).Delete()
	if err != nil {
		goto ERREND
	}
	return nil
ERREND:
	var dbError *dbcore.DbError
	dbError = &dbcore.DbError{
		Op:     "DeleteReceiveRecord",
		Entity: ReceiveRecordTableName,
		Err:    err,
	}
	return dbError
}

//QueryReceiveRecordsAll query ReceiveRecord
func QueryReceiveRecordsAll(selecter upper.Cond, order ...interface{}) ([]ReceiveRecordEntity, error) {
	sess, err := dbcore.GetDBSession()
	var ReceiveRecordes []ReceiveRecordEntity
	if err != nil {
		goto ERREND
	}
	err = sess.Collection(ReceiveRecordTableName).Find(selecter).OrderBy(order...).All(&ReceiveRecordes)
	if err != nil {
		goto ERREND
	}
	return ReceiveRecordes, nil
ERREND:
	var dbError *dbcore.DbError
	dbError = &dbcore.DbError{
		Op:     "QueryReceiveRecordsAll",
		Entity: ReceiveRecordTableName,
		Err:    err,
	}
	return nil, dbError
}

//QueryReceiveRecords query ReceiveRecord
func QueryReceiveRecords(selecter upper.Cond, offset, limit int, order ...interface{}) ([]ReceiveRecordEntity, error) {
	sess, err := dbcore.GetDBSession()
	var ReceiveRecordes []ReceiveRecordEntity
	if err != nil {
		goto ERREND
	}

	err = sess.Collection(ReceiveRecordTableName).Find(selecter).OrderBy(order...).Offset(offset).Limit(limit).All(&ReceiveRecordes)
	if err != nil {
		goto ERREND
	}
	return ReceiveRecordes, nil
ERREND:
	var dbError *dbcore.DbError
	dbError = &dbcore.DbError{
		Op:     "QueryReceiveRecords",
		Entity: ReceiveRecordTableName,
		Err:    err,
	}
	return nil, dbError
}

//QueryReceiveRecordCount query ReceiveRecord count
func QueryReceiveRecordCount(selecter upper.Cond) (uint64, error) {
	sess, err := dbcore.GetDBSession()
	var c uint64
	if err != nil {
		goto ERREND
	}

	c, err = sess.Collection(ReceiveRecordTableName).Find(selecter).Count()
	if err != nil {
		goto ERREND
	}
	return c, nil
ERREND:
	var dbError *dbcore.DbError
	dbError = &dbcore.DbError{
		Op:     "QueryReceiveRecordCount",
		Entity: ReceiveRecordTableName,
		Err:    err,
	}
	return 0, dbError
}

//UpdateReceiveRecord update ReceiveRecord
func UpdateReceiveRecord(ety *ReceiveRecordEntity) error {
	var dbError *dbcore.DbError
	sess, err := dbcore.GetDBSession()
	if err != nil {
		goto ERREND
	}
	err = sess.Collection(ReceiveRecordTableName).UpdateReturning(ety)
	if err != nil {
		goto ERREND
	}
	return nil
ERREND:
	dbError = &dbcore.DbError{
		Op:     "UpdateReceiveRecord",
		Entity: ReceiveRecordTableName,
		Err:    err,
	}
	return dbError

}
