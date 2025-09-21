package repository

import (
	"sync"

	"github.com/audricimanuel/awb-stock-allocation/src/model"
)

// TODO: the interface and implementation under here
type (
	OrderRepository interface {
		CreateOrder(order *model.Order) (*model.Order, error)
	}

	OrderRepositoryImpl struct {
		mu   sync.RWMutex
		list *[]model.Order
	}
)

func NewOrderRepository(list *[]model.Order) OrderRepository {
	return &OrderRepositoryImpl{
		list: list,
	}
}

func (r *OrderRepositoryImpl) CreateOrder(order *model.Order) (*model.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	*r.list = append(*r.list, *order)

	return order, nil
}
