package repository

import (
	"sync"

	"github.com/audricimanuel/awb-stock-allocation/src/model"
	e "github.com/audricimanuel/awb-stock-allocation/utils/errors"
)

type (
	AWBStockRepository interface {
		// TODO: create the functions needed to implement here
		GetAWBStock() ([]*model.AWBStock, error)
		GetAWBStockByAWBNumber(AWBNumber string) (*model.AWBStock, error)
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

func (r *AWBStockRepositoryImpl) GetAWBStockByAWBNumber(AWBNumber string) (*model.AWBStock, error) {
	var err error

	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, v := range *r.list {
		if v.AWBNumber == AWBNumber {
			return &v, nil
		} else {
			err = e.ErrAWBNotFound
		}
	}

	return nil, err
}
