package entity

import (
	"fmt"
	"log"
	"strconv"

	dbcore "github.com/lufred/red_envelope/Service/red_envelope_service/db/core"
	upper "upper.io/db.v3"
)

//ClearRedeEntity 红包过期清退实体
type ClearRedeEntity struct {
	ID            int32 `db:"id" json:"db"`
	UserID        int32 `db:"user_id" json:"user_id"`
	RedEnvelopeID int32 `db:"red_envelope_id" json:"red_envelope_id"`
	Amount        int32 `db:"amount" json:"amount"`
	Status        byte  `db:"status" json:"status"`
	GMTCreate     int64 `db:"gmt_create" json:"gmt_create"`     //创建时间
	GMTModified   int64 `db:"gmt_modified" json:"gmt_modified"` //修改时间
}

var ClearRedeTableName = "clear_rede"

//InsertClearRede insert ClearRede
func InsertClearRede(ety *ClearRedeEntity) (*ClearRedeEntity, error) {
	var dbError *dbcore.DbError
	var id int
	var sid string
	var _id interface{}
	sess, err := dbcore.GetDBSession()
	log.Println(sess)
	if err != nil {
		goto ERREND
	}
	_id, err = sess.Collection(ClearRedeTableName).Insert(ety)
	if err != nil {
		goto ERREND
	}
	sid = fmt.Sprint(_id)
	id, err = strconv.Atoi(sid)
	ety.ID = int32(id)
	return ety, nil
ERREND:
	dbError = &dbcore.DbError{
		Op:     "InsertClearRede",
		Entity: ClearRedeTableName,
		Err:    err,
	}
	return nil, dbError
}

//QueryClearRedeOne query a single record
func QueryClearRedeOne(selecter upper.Cond) (*ClearRedeEntity, error) {
	var dbError *dbcore.DbError
	ety := &ClearRedeEntity{}
	sess, err := dbcore.GetDBSession()
	if err != nil {
		goto ERREND
	}
	err = sess.Collection(ClearRedeTableName).Find(selecter).One(ety)
	if err != nil {
		if err == upper.ErrNoMoreRows {
			return nil, nil
		}
		goto ERREND
	}
	return ety, nil
ERREND:
	dbError = &dbcore.DbError{
		Op:     "QueryClearRedeOne",
		Entity: ClearRedeTableName,
		Err:    err,
	}
	return nil, dbError
}

//DeleteClearRede delete ClearRede
func DeleteClearRede(id int32) error {
	sess, err := dbcore.GetDBSession()
	if err != nil {
		goto ERREND
	}
	err = sess.Collection(ClearRedeTableName).Find(upper.Cond{
		"id": id,
	}).Delete()
	if err != nil {
		goto ERREND
	}
	return nil
ERREND:
	var dbError *dbcore.DbError
	dbError = &dbcore.DbError{
		Op:     "DeleteClearRede",
		Entity: ClearRedeTableName,
		Err:    err,
	}
	return dbError
}

//QueryClearRedesAll query ClearRede
func QueryClearRedesAll(selecter upper.Cond, order ...interface{}) ([]ClearRedeEntity, error) {
	sess, err := dbcore.GetDBSession()
	var ClearRedees []ClearRedeEntity
	if err != nil {
		goto ERREND
	}
	err = sess.Collection(ClearRedeTableName).Find(selecter).OrderBy(order...).All(&ClearRedees)
	if err != nil {
		goto ERREND
	}
	return ClearRedees, nil
ERREND:
	var dbError *dbcore.DbError
	dbError = &dbcore.DbError{
		Op:     "QueryClearRedesAll",
		Entity: ClearRedeTableName,
		Err:    err,
	}
	return nil, dbError
}

//QueryClearRedes query ClearRede
func QueryClearRedes(selecter upper.Cond, offset, limit int, order ...interface{}) ([]ClearRedeEntity, error) {
	sess, err := dbcore.GetDBSession()
	var ClearRedees []ClearRedeEntity
	if err != nil {
		goto ERREND
	}

	err = sess.Collection(ClearRedeTableName).Find(selecter).OrderBy(order...).Offset(offset).Limit(limit).All(&ClearRedees)
	if err != nil {
		goto ERREND
	}
	return ClearRedees, nil
ERREND:
	var dbError *dbcore.DbError
	dbError = &dbcore.DbError{
		Op:     "QueryClearRedes",
		Entity: ClearRedeTableName,
		Err:    err,
	}
	return nil, dbError
}

//QueryClearRedeCount query ClearRede count
func QueryClearRedeCount(selecter upper.Cond) (uint64, error) {
	sess, err := dbcore.GetDBSession()
	var c uint64
	if err != nil {
		goto ERREND
	}

	c, err = sess.Collection(ClearRedeTableName).Find(selecter).Count()
	if err != nil {
		goto ERREND
	}
	return c, nil
ERREND:
	var dbError *dbcore.DbError
	dbError = &dbcore.DbError{
		Op:     "QueryClearRedeCount",
		Entity: ClearRedeTableName,
		Err:    err,
	}
	return 0, dbError
}

//UpdateClearRede update ClearRede
func UpdateClearRede(ety *ClearRedeEntity) error {
	var dbError *dbcore.DbError
	sess, err := dbcore.GetDBSession()
	if err != nil {
		goto ERREND
	}
	err = sess.Collection(ClearRedeTableName).UpdateReturning(ety)
	if err != nil {
		goto ERREND
	}
	return nil
ERREND:
	dbError = &dbcore.DbError{
		Op:     "UpdateClearRede",
		Entity: ClearRedeTableName,
		Err:    err,
	}
	return dbError

}
