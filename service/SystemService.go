package service

import (
	"bus-backend-go/model"
	repo "bus-backend-go/repo"
	"github.com/sirupsen/logrus"
)

type SystemService interface {
	// 超级管理员
	FindSuperAdminByLdap(string) (model.SuperAdmin, error)
	AddSuperAdmin(model.SuperAdmin) (model.SuperAdmin, error)
}

type systemService struct {
	log       *logrus.Logger
	Systemrep repo.SystemRepository
}

func NewSystemService(log *logrus.Logger) SystemService {
	return &systemService{log: log, Systemrep: repo.NewSystemRepository()}
}

// 通过ldap获取super admin记录
func (s systemService) FindSuperAdminByLdap(ldap string) (model.SuperAdmin, error) {
	return s.Systemrep.FindSuperAdminByLdap(ldap)
}

// 新增管理员
func (s systemService) AddSuperAdmin(admin model.SuperAdmin) (model.SuperAdmin, error) {
	// 检查是否已存在
	if adminObj, err := s.Systemrep.FindSuperAdminByLdap(admin.Ldap); err == nil {
		return adminObj, nil
	}
	return s.Systemrep.AddSuperAdmin(admin)
}