package service

import (
	"bus-backend-go/model"
	"bus-backend-go/repo"
	"github.com/sirupsen/logrus"
)

type SlService interface {
	FindServiceList()([]*re_serviceList, error)
	AddServiceList(serviceList model.ServiceList)(model.ServiceList, error)
	UpdateServiceList(serviceList model.ServiceList)(model.ServiceList, error)
	DeleteServiceList(model.ServiceList)(model.ServiceList, error)
}

type slService struct{
	log *logrus.Logger
	slRepo repo.ServiceListRepository
}

func NewSlService(log *logrus.Logger) SlService {
	return &slService{log: log, slRepo: repo.NewServiceListRepository()}
}

type re_serviceList struct {
	ID      uint      `json:"id"`
	Name        string      `json:"name"`
	Url         string      `json:"url"`
	Status      uint      `json:"status"`
	Desc        string      `json:"desc"`
	MenuId      uint        `json:"menu_id"`
	Menu        string      `json:"menu"`
	ClassId     uint        `json:"class_id"`
	ClassName   string      `json:"class_name"`
	PoPo        string      `json:"popo"`
	Doc         string      `json:"doc"`
}

func (s slService) FindServiceList()([]*re_serviceList, error){
	res, err := s.slRepo.FindServiceList()
	var response []*re_serviceList
	response = make([]*re_serviceList, 0)
	for _, value := range res{
		resp := new(re_serviceList)
		resp.ID = value.ID
		resp.Name = value.Name
		resp.Desc = value.Desc
		resp.Url = value.Url
		resp.PoPo = value.PoPo
		resp.Doc = value.Doc
		resp.Status = value.Status
		resp.MenuId = value.MenuId
		resp.Menu = value.MicroList.Name
		resp.ClassId = value.ClassId
		resp.ClassName = value.ServiceType.ClassName
		response = append(response, resp)
	}
	return response, err
}

func (s slService) AddServiceList(serviceList model.ServiceList)(model.ServiceList, error){
	return s.slRepo.AddServiceList(serviceList)
}

func (s slService) UpdateServiceList(serviceList model.ServiceList)(model.ServiceList, error){
	return s.slRepo.UpdateServiceList(serviceList)
}

func (s slService) DeleteServiceList(sl model.ServiceList)(model.ServiceList, error){
	return s.slRepo.DeleteServiceList(sl)
}