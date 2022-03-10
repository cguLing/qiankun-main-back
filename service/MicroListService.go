package service

import (
	"bus-backend-go/model"
	"bus-backend-go/repo"
	"github.com/sirupsen/logrus"
)

type MlService interface {
	FindMicroList()([]model.MicroList, error)
	AddMicroList(microList model.MicroList)(model.MicroList, error)
	UpdateMicroList(microList model.MicroList)(model.MicroList, error)
	DeleteMicroList(model.MicroList)(model.MicroList, error)
}

type mlService struct{
	log *logrus.Logger
	mlRepo repo.MicroListRepository
}

func NewMlService(log *logrus.Logger) MlService {
	return &mlService{log: log, mlRepo: repo.NewMicroListRepository()}
}

func (s mlService) FindMicroList()([]model.MicroList, error){
	return s.mlRepo.FindMicroList()
}

func (s mlService) AddMicroList(microList model.MicroList)(model.MicroList, error){
	return s.mlRepo.AddMicroList(microList)
}

func (s mlService) UpdateMicroList(microList model.MicroList)(model.MicroList, error){
	return s.mlRepo.UpdateMicroList(microList)
}

func (s mlService) DeleteMicroList(ml model.MicroList)(model.MicroList, error){
	return s.mlRepo.DeleteMicroList(ml)
}