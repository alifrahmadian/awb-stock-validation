package service

import (
	"github.com/audricimanuel/awb-stock-allocation/src/internals/repository"
)

type (
	AWBStockService interface {
		// TODO: create the functions needed to implement here
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
