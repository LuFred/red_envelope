package core

import (
	"errors"
	"sync"

	"github.com/lufred/red_envelope/Service/red_envelope_service/config"
	"upper.io/db.v3/lib/sqlbuilder"
	"upper.io/db.v3/mysql"
)

//自定义Error
//不存在和类型匹配的表名
var ErrNoneTb = errors.New("mapper:not exist in the tablename")

//类型映射map
var typeMapper = map[interface{}]string{}

//连接字符串对象
var settings = mysql.ConnectionURL{}
var dbClient sqlbuilder.Database
var m *sync.Mutex

func init() {
	m = new(sync.Mutex)
}

//getDBSetting 获取数据库连接设置对象
func getDBSetting() *mysql.ConnectionURL {
	settings = mysql.ConnectionURL{
		User:     config.ProConfig.Mysql.UserName,
		Password: config.ProConfig.Mysql.PWD,
		Host:     config.ProConfig.Mysql.Host,
		Database: config.ProConfig.Mysql.DB,
	}
	return &settings
}

//GetDBSession 获取数据库连接对象
func GetDBSession() (sqlbuilder.Database, error) {
	if dbClient == nil {
		m.Lock()
		if dbClient == nil {
			db, err := mysql.Open(getDBSetting())
			if err != nil {
				m.Unlock()
				return nil, err
			}
			dbClient = db
		}
		m.Unlock()
	}
	return dbClient, nil
}