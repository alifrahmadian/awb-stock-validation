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
		GetAWBStockByAWBNumber(awbNumber string) *model.AWBStock
		UpdateAWBStatus(awbNumber string)
		CreateAWBStock(awbStock *model.AWBStock) *model.AWBStock
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

func (r *AWBStockRepositoryImpl) GetAWBStockByAWBNumber(awbNumber string) *model.AWBStock {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for i := range *r.list {
		if (*r.list)[i].AWBNumber == awbNumber {
			return &(*r.list)[i]
		}
	}

	return nil
}

func (r *AWBStockRepositoryImpl) UpdateAWBStatus(awbNumber string) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i := range *r.list {
		if (*r.list)[i].AWBNumber == awbNumber {
			(*r.list)[i].Status = constants.AWB_STATUS_IN_USE
		}
	}
}

func (r *AWBStockRepositoryImpl) CreateAWBStock(awbStock *model.AWBStock) *model.AWBStock {
	r.mu.Lock()
	defer r.mu.Unlock()

	*r.list = append(*r.list, *awbStock)

	return &(*r.list)[len(*r.list)-1]
}
