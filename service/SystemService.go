package service

import (
	"bus-backend-go/model"
	"bus-backend-go/repo"
	"github.com/sirupsen/logrus"
)

type SystemService interface {
	FindUserAdminByLdap(string) (model.UserAdmin, error)
	AddUserAdmin(model.UserAdmin) (model.UserAdmin, error)
	UpdateUserAdmin(ua model.UserAdmin) (uas model.UserAdmin, err error)
}

type systemService struct {
	log       *logrus.Logger
	Systemrep repo.SystemRepository
}

func NewSystemService(log *logrus.Logger) SystemService {
	return &systemService{log: log, Systemrep: repo.NewSystemRepository()}
}

// 通过ldap获取记录，不存在则创建
func (s systemService) FindUserAdminByLdap(ldap string) (model.UserAdmin, error) {
	return s.Systemrep.FindUserAdminByLdap(ldap)
}

func (s systemService) AddUserAdmin(admin model.UserAdmin) (model.UserAdmin, error) {
	// 检查是否已存在
	if adminObj, err := s.Systemrep.FindUserAdminByLdap(admin.Ldap); err == nil {
		return adminObj, nil
	}
	return s.Systemrep.AddUserAdmin(admin)
}

func (s systemService) UpdateUserAdmin(userAdmin model.UserAdmin)(model.UserAdmin, error){
	return s.Systemrep.UpdateUserAdmin(userAdmin)
}