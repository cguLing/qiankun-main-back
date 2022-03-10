package service

import (
	"bus-backend-go/model"
	"bus-backend-go/repo"
	"github.com/sirupsen/logrus"
)

type ElService interface {
	FindEnshrineList(userName string) ([]model.EnshrineList, error)
	AddEnshrineList(enshrineList model.EnshrineList)(model.EnshrineList, error)
	DeleteEnshrineList(model.EnshrineList)(model.EnshrineList, error)
}

type elService struct{
	log *logrus.Logger
	elRepo repo.EnshrineListRepository
}

func NewElService(log *logrus.Logger) ElService {
	return &elService{log: log, elRepo: repo.NewEnshrineListRepository()}
}


func (s elService) FindEnshrineList(userName string) ([]model.EnshrineList, error){
	return s.elRepo.FindEnshrineList(userName)
}

func (s elService) AddEnshrineList(enshrineList model.EnshrineList)(model.EnshrineList, error){
	return s.elRepo.AddEnshrineList(enshrineList)
}

func (s elService) DeleteEnshrineList(el model.EnshrineList)(model.EnshrineList, error){
	return s.elRepo.DeleteEnshrineList(el)
}