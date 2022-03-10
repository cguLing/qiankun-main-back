package repo

import (
	"bus-backend-go/datasource"
	"bus-backend-go/model"
	"gorm.io/gorm"
)

type MicroListRepository interface {
	FindMicroList() (list []model.MicroList, err error)
	AddMicroList(ml model.MicroList) (mls model.MicroList, err error)
	UpdateMicroList(ml model.MicroList) (mls model.MicroList, err error)
	DeleteMicroList(microList model.MicroList) (ml model.MicroList, err error)
}

func NewMicroListRepository() MicroListRepository {
	return &mlRepository{db: datasource.GetDB()}
}

type mlRepository struct {
	db *gorm.DB
}

// FindMicroList 全量获取微服务列表
func (o mlRepository) FindMicroList() (list []model.MicroList, err error) {

	err = o.db.Find(&list).Error
	return
}

func (o mlRepository) AddMicroList(ml model.MicroList) (model.MicroList, error){
	err := o.db.Create(&ml).Error
	return ml, err
}

func (o mlRepository) UpdateMicroList(ml model.MicroList) (model.MicroList, error){
	err := o.db.Model(&model.MicroList{}).Where("`id` = ?", ml.ID).Updates(&ml).Error
	return ml, err
}

func (o mlRepository) DeleteMicroList(microList model.MicroList) (model.MicroList, error){
	err := o.db.Where("`id` = ?", microList.ID).Delete(&microList).Error
	return microList, err
}