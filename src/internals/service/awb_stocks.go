package service

import (
	"github.com/audricimanuel/awb-stock-allocation/src/internals/repository"
	"github.com/audricimanuel/awb-stock-allocation/src/model"
)

type (
	AWBStockService interface {
		// TODO: create the functions needed to implement here
		GetAWBStock() ([]*model.AWBStock, error)
	}

	AWBStockServiceImpl struct {
		awbStockRepository repository.AWBStockRepository
	}
)

func NewAWBStockService(awbStockRepository repository.AWBStockRepository) AWBStockService {
	return &AWBStockServiceImpl{
		awbStockRepository: awbStockRepository,
	}
}

func (s *AWBStockServiceImpl) GetAWBStock() ([]*model.AWBStock, error) {
	return s.awbStockRepository.GetAWBStock()
}
