package repo

import (
	"bus-backend-go/datasource"
	"bus-backend-go/model"
	"gorm.io/gorm"
)

type SystemRepository interface {
	FindSuperAdminByLdap(string) (model.SuperAdmin, error)
	AddSuperAdmin(model.SuperAdmin) (model.SuperAdmin, error)
}

func NewSystemRepository() SystemRepository {
	return &systemRepository{db: datasource.GetDB()}
}

type systemRepository struct {
	db *gorm.DB
}


// 对数据库的直接curd操作
// 根据ldap查询
func (o systemRepository) FindSuperAdminByLdap(ldap string) (superAdmin model.SuperAdmin, err error) {
	err = o.db.Where("ldap = ?", ldap).First(&superAdmin, model.SuperAdmin{}).Error
	return
}

// 新增超级管理员
func (o systemRepository) AddSuperAdmin(admin model.SuperAdmin) (model.SuperAdmin, error) {
	res := o.db.Create(&admin)
	return admin, res.Error
}

