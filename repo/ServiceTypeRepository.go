package repo

import (
	"bus-backend-go/datasource"
	"bus-backend-go/model"
	"gorm.io/gorm"
)

type ServiceTypeRepository interface {
	FindServiceType() (list []model.ServiceType, err error)
	AddServiceType(st model.ServiceType) (sts model.ServiceType, err error)
	UpdateServiceType(st model.ServiceType) (sts model.ServiceType, err error)
	DeleteServiceType(serviceType model.ServiceType) (st model.ServiceType, err error)
}

func NewServiceTypeRepository() ServiceTypeRepository {
	return &stRepository{db: datasource.GetDB()}
}

type stRepository struct {
	db *gorm.DB
}

// FindServiceType 全量获取服务类型并以index升序
func (o stRepository) FindServiceType() (list []model.ServiceType, err error) {
	err = o.db.Order("`index`").Find(&list).Error
	return
}

func (o stRepository) AddServiceType(st model.ServiceType) (model.ServiceType, error){
	err := o.db.Create(&st).Error
	return st, err
}

func (o stRepository) UpdateServiceType(st model.ServiceType) (model.ServiceType, error){
	err := o.db.Model(&model.ServiceType{}).Where("class_name = ?", st.ClassName).Updates(&st).Error
	return st, err
}

func (o stRepository) DeleteServiceType(serviceType model.ServiceType) (model.ServiceType, error){
	err := o.db.Where("`id` = ?", serviceType.ID).Delete(&serviceType).Error
	return serviceType, err
}