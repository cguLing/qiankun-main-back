package repo

import (
	"bus-backend-go/datasource"
	"bus-backend-go/model"
	"gorm.io/gorm"
)

type EnshrineListRepository interface {
	FindEnshrineList(userName string) (list []model.EnshrineList, err error)
	AddEnshrineList(el model.EnshrineList) (els model.EnshrineList, err error)
	DeleteEnshrineList(enshrineList model.EnshrineList) (el model.EnshrineList, err error)
}

func NewEnshrineListRepository() EnshrineListRepository {
	return &elRepository{db: datasource.GetDB()}
}

type elRepository struct {
	db *gorm.DB
}

// FindEnshrineList 获取对应用户的服务列表
func (o elRepository) FindEnshrineList(userName string) (list []model.EnshrineList, err error) {
	//var ids []uint
	//o.db.Table("super_admin").Select("super_admin.id").Where("ldap = ?", userName).Scan(&ids)
	//err = o.db.Where("`user_id` = ?", ids[0]).
	//	Preload("ServiceList").Find(&list, model.EnshrineList{}).Error
	err = o.db.Where("`user_name` = ?", userName).
		Preload("ServiceList").Find(&list, model.EnshrineList{}).Error
	return
}

func (o elRepository) AddEnshrineList(el model.EnshrineList) (model.EnshrineList, error){
	err := o.db.Create(&el).Error
	return el, err
}


func (o elRepository) DeleteEnshrineList(enshrineList model.EnshrineList) (model.EnshrineList, error){
	err := o.db.Where("`user_name` = ? and `service_id` = ?", enshrineList.UserName, enshrineList.ServiceId).Delete(&enshrineList).Error
	return enshrineList, err
}