package dao

import (
	"context"

	"github.com/pengzhong2010/go-server-base/common/pkg/client/mysqldb"

	"github.com/jinzhu/gorm"
)

type Dao interface {
	Close()
	Ping(ctx context.Context) (err error)
}

// New new a topic Dao
func New() (d Dao, cf func(), err error) {
	return newDao(mysqldb.Mysqldb)
}

type dao struct {
	db *gorm.DB
}

func newDao(db *gorm.DB) (d *dao, cf func(), err error) {
	d = &dao{
		db: db,
	}
	cf = d.Close
	return
}

// Close dao.
func (d *dao) Close() {}

// Ping check connection status.
func (d *dao) Ping(c context.Context) (err error) {
	return
}
