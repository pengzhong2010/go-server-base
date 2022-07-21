package service

import (
	"github.com/pengzhong2010/go-server-base/app/app1/internal/dao"
	"github.com/pengzhong2010/go-server-base/common/pkg/log"
)

type Service struct {
	dao dao.Dao
}

var Svc *Service

func New() (s *Service, cf func(), err error) {
	d, daoClose, errT1 := dao.New()
	if errT1 != nil {
		err = errT1
		log.Error("incidentService New: %v", err)
		return
	}
	Svc = &Service{
		dao: d,
	}
	cf = func() {
		daoClose()
	}
	return Svc, cf, err
}
