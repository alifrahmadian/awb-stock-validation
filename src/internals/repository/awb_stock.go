package repository

import (
	"sync"

	"github.com/audricimanuel/awb-stock-allocation/src/model"
)

type (
	AWBStockRepository interface {
		// TODO: create the functions needed to implement here
		GetAWBStock() ([]*model.AWBStock, error)
		// GetAWBStockByAWB(awbNumber string) (*model.AWBStock, error)
	}

	AWBStockRepositoryImpl struct {
		mu   sync.RWMutex
		list *[]model.AWBStock
	}
)

func NewAWBStockRepository(list *[]model.AWBStock) AWBStockRepository {
	return &AWBStockRepositoryImpl{
		list: list,
	}
}

func (r *AWBStockRepositoryImpl) GetAWBStock() ([]*model.AWBStock, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var awbStocks []*model.AWBStock
	for _, awbStock := range *r.list {
		awbStocks = append(awbStocks, &awbStock)
	}

	return awbStocks, nil
}
