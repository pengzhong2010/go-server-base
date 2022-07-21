package mysqldb

import (
	"github.com/jinzhu/gorm"
	"github.com/pengzhong2010/go-server-base/app/app1/conf"
	"github.com/pengzhong2010/go-server-base/common/pkg/log"
)

var Mysqldb *gorm.DB

func New() (db *gorm.DB, closeFunc func(), err error) {
	Mysqldb, err = gorm.Open("mysql", conf.MYSQL_USER+":"+conf.MYSQL_PASSWD+"@tcp("+conf.MYSQL_HOST+":"+conf.MYSQL_PORT+")/"+conf.MYSQL_DB_NAME+"?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		log.Error(err)
		return
	}
	Mysqldb.DB().SetMaxIdleConns(20)
	Mysqldb.DB().SetMaxOpenConns(20)
	closeFunc = CloseFunc
	return Mysqldb, closeFunc, err
}
func CloseFunc() {
	Mysqldb.Close()
}
