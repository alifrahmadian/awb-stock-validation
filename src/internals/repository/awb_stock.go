package repository

import (
	"sync"

	"github.com/audricimanuel/awb-stock-allocation/src/model"
	"github.com/audricimanuel/awb-stock-allocation/utils/constants"
)

type (
	AWBStockRepository interface {
		// TODO: create the functions needed to implement here
		GetAWBStock() ([]*model.AWBStock, error)
		GetAWBStockByAWBNumber(AWBNumber string) *model.AWBStock
		UpdateAWBStatus(AWBNumber string)
		CreateAWBStock(AWBStock *model.AWBStock) *model.AWBStock
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

func (r *AWBStockRepositoryImpl) GetAWBStockByAWBNumber(AWBNumber string) *model.AWBStock {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for _, v := range *r.list {
		if v.AWBNumber == AWBNumber {
			return &v
		}
	}

	return nil
}

func (r *AWBStockRepositoryImpl) UpdateAWBStatus(AWBNumber string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i := range *r.list {
		if (*r.list)[i].AWBNumber == AWBNumber {
			(*r.list)[i].Status = constants.AWB_STATUS_IN_USE
		}
	}
}

func (r *AWBStockRepositoryImpl) CreateAWBStock(AWBStock *model.AWBStock) *model.AWBStock {
	r.mu.RLock()
	defer r.mu.RUnlock()

	*r.list = append(*r.list, *AWBStock)

	return AWBStock
}
