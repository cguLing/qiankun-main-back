package repo

import (
	"bus-backend-go/datasource"
	"bus-backend-go/model"
	"gorm.io/gorm"
)

type ServiceListRepository interface {
	FindServiceList() (list []model.ServiceList, err error)
	AddServiceList(sl model.ServiceList) (sls model.ServiceList, err error)
	UpdateServiceList(sl model.ServiceList) (sls model.ServiceList, err error)
	DeleteServiceList(serviceList model.ServiceList) (sl model.ServiceList, err error)
}

func NewServiceListRepository() ServiceListRepository {
	return &slRepository{db: datasource.GetDB()}
}

type slRepository struct {
	db *gorm.DB
}

// FindServiceList 全量获取服务列表
func (o slRepository) FindServiceList() (list []model.ServiceList, err error) {
	err = o.db.Preload("ServiceType").Preload("MicroList").Find(&list, model.ServiceList{}).Error
	return
}

func (o slRepository) AddServiceList(sl model.ServiceList) (model.ServiceList, error){
	err := o.db.Create(&sl).Error
	return sl, err
}

func (o slRepository) UpdateServiceList(sl model.ServiceList) (model.ServiceList, error){
	err := o.db.Model(&model.ServiceList{}).Where("`id` = ?", sl.ID).Updates(&sl).Error
	return sl, err
}

func (o slRepository) DeleteServiceList(serviceList model.ServiceList) (model.ServiceList, error){
	err := o.db.Where("`id` = ?", serviceList.ID).Delete(&serviceList).Error
	return serviceList, err
}