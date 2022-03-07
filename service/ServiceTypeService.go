package service

import (
	"bus-backend-go/model"
	"bus-backend-go/repo"
	"github.com/sirupsen/logrus"
)

type StService interface {
	FindServiceType()([]model.ServiceType, error)
	AddServiceType(serviceType model.ServiceType)(model.ServiceType, error)
	UpdateServiceType(serviceType model.ServiceType)(model.ServiceType, error)
	DeleteServiceType(model.ServiceType)(model.ServiceType, error)
}

type stService struct{
	log *logrus.Logger
	stRepo repo.ServiceTypeRepository
}

func NewStService(log *logrus.Logger) StService {
	return &stService{log: log, stRepo: repo.NewServiceTypeRepository()}
}

func (s stService) FindServiceType()([]model.ServiceType, error){
	return s.stRepo.FindServiceType()
}

func (s stService) AddServiceType(serviceType model.ServiceType)(model.ServiceType, error){
	return s.stRepo.AddServiceType(serviceType)
}

func (s stService) UpdateServiceType(serviceType model.ServiceType)(model.ServiceType, error){
	return s.stRepo.UpdateServiceType(serviceType)
}

func (s stService) DeleteServiceType(st model.ServiceType)(model.ServiceType, error){
	return s.stRepo.DeleteServiceType(st)
}