package repo

import (
	"bus-backend-go/datasource"
	"bus-backend-go/model"
	"gorm.io/gorm"
)

type SystemRepository interface {
	FindUserAdminByLdap(string) (model.UserAdmin, error)
	AddUserAdmin(model.UserAdmin) (model.UserAdmin, error)
	UpdateUserAdmin(ua model.UserAdmin) (uas model.UserAdmin, err error)
}

func NewSystemRepository() SystemRepository {
	return &systemRepository{db: datasource.GetDB()}
}

type systemRepository struct {
	db *gorm.DB
}


func (o systemRepository) FindUserAdminByLdap(ldap string) (userAdmin model.UserAdmin, err error) {
	err = o.db.Where("ldap = ?", ldap).FirstOrCreate(&userAdmin, model.UserAdmin{Ldap: ldap}).Error
	return
}

func (o systemRepository) AddUserAdmin(admin model.UserAdmin) (model.UserAdmin, error) {
	res := o.db.Create(&admin)
	return admin, res.Error
}

func (o systemRepository) UpdateUserAdmin(ua model.UserAdmin) (model.UserAdmin, error){
	err := o.db.Model(&model.UserAdmin{}).Where("`ldap` = ?", ua.Ldap).Updates(&ua).Update("`menu_way`", ua.MenuWay).Error
	return ua, err
}