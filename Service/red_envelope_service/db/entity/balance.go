package entity

import (
	"fmt"
	"log"
	"strconv"

	dbcore "github.com/lufred/red_envelope/Service/red_envelope_service/db/core"
	upper "upper.io/db.v3"
)

//BalanceEntity 余额实体
type BalanceEntity struct {
	ID          int32 `db:"id" json:"db"`
	UserID      int32 `db:"user_id" json:"user_id"`
	Balance     int32 `db:"balance" json:"balance"`
	GMTCreate   int64 `db:"gmt_create" json:"gmt_create"`     //创建时间
	GMTModified int64 `db:"gmt_modified" json:"gmt_modified"` //修改时间
}

var BalanceTableName = "balance"

//InsertBalance insert Balance
func InsertBalance(ety *BalanceEntity) (*BalanceEntity, error) {
	var dbError *dbcore.DbError
	var id int
	var sid string
	var _id interface{}
	sess, err := dbcore.GetDBSession()
	log.Println(sess)
	if err != nil {
		goto ERREND
	}
	_id, err = sess.Collection(BalanceTableName).Insert(ety)
	if err != nil {
		goto ERREND
	}
	sid = fmt.Sprint(_id)
	id, err = strconv.Atoi(sid)
	ety.ID = int32(id)
	return ety, nil
ERREND:
	dbError = &dbcore.DbError{
		Op:     "InsertBalance",
		Entity: BalanceTableName,
		Err:    err,
	}
	return nil, dbError
}

//QueryBalanceOne query a single record
func QueryBalanceOne(selecter upper.Cond) (*BalanceEntity, error) {
	var dbError *dbcore.DbError
	ety := &BalanceEntity{}
	sess, err := dbcore.GetDBSession()
	if err != nil {
		goto ERREND
	}
	err = sess.Collection(BalanceTableName).Find(selecter).One(ety)
	if err != nil {
		if err == upper.ErrNoMoreRows {
			return nil, nil
		}
		goto ERREND
	}
	return ety, nil
ERREND:
	dbError = &dbcore.DbError{
		Op:     "QueryBalanceOne",
		Entity: BalanceTableName,
		Err:    err,
	}
	return nil, dbError
}

//DeleteBalance delete Balance
func DeleteBalance(id int32) error {
	sess, err := dbcore.GetDBSession()
	if err != nil {
		goto ERREND
	}
	err = sess.Collection(BalanceTableName).Find(upper.Cond{
		"id": id,
	}).Delete()
	if err != nil {
		goto ERREND
	}
	return nil
ERREND:
	var dbError *dbcore.DbError
	dbError = &dbcore.DbError{
		Op:     "DeleteBalance",
		Entity: BalanceTableName,
		Err:    err,
	}
	return dbError
}

//QueryBalancesAll query Balance
func QueryBalancesAll(selecter upper.Cond, order ...interface{}) ([]BalanceEntity, error) {
	sess, err := dbcore.GetDBSession()
	var Balancees []BalanceEntity
	if err != nil {
		goto ERREND
	}
	err = sess.Collection(BalanceTableName).Find(selecter).OrderBy(order...).All(&Balancees)
	if err != nil {
		goto ERREND
	}
	return Balancees, nil
ERREND:
	var dbError *dbcore.DbError
	dbError = &dbcore.DbError{
		Op:     "QueryBalancesAll",
		Entity: BalanceTableName,
		Err:    err,
	}
	return nil, dbError
}

//QueryBalances query Balance
func QueryBalances(selecter upper.Cond, offset, limit int, order ...interface{}) ([]BalanceEntity, error) {
	sess, err := dbcore.GetDBSession()
	var Balancees []BalanceEntity
	if err != nil {
		goto ERREND
	}

	err = sess.Collection(BalanceTableName).Find(selecter).OrderBy(order...).Offset(offset).Limit(limit).All(&Balancees)
	if err != nil {
		goto ERREND
	}
	return Balancees, nil
ERREND:
	var dbError *dbcore.DbError
	dbError = &dbcore.DbError{
		Op:     "QueryBalances",
		Entity: BalanceTableName,
		Err:    err,
	}
	return nil, dbError
}

//QueryBalanceCount query Balance count
func QueryBalanceCount(selecter upper.Cond) (uint64, error) {
	sess, err := dbcore.GetDBSession()
	var c uint64
	if err != nil {
		goto ERREND
	}

	c, err = sess.Collection(BalanceTableName).Find(selecter).Count()
	if err != nil {
		goto ERREND
	}
	return c, nil
ERREND:
	var dbError *dbcore.DbError
	dbError = &dbcore.DbError{
		Op:     "QueryBalanceCount",
		Entity: BalanceTableName,
		Err:    err,
	}
	return 0, dbError
}

//UpdateBalance update Balance
func UpdateBalance(ety *BalanceEntity) error {
	var dbError *dbcore.DbError
	sess, err := dbcore.GetDBSession()
	if err != nil {
		goto ERREND
	}
	err = sess.Collection(BalanceTableName).UpdateReturning(ety)
	if err != nil {
		goto ERREND
	}
	return nil
ERREND:
	dbError = &dbcore.DbError{
		Op:     "UpdateBalance",
		Entity: BalanceTableName,
		Err:    err,
	}
	return dbError

}
