package repository

import (
	"sync"

	"github.com/audricimanuel/awb-stock-allocation/src/model"
	e "github.com/audricimanuel/awb-stock-allocation/utils/errors"
)

// TODO: the interface and implementation under here
type (
	OrderRepository interface {
		CreateOrder(order *model.Order) *model.Order
		UpdateOrder(id int64, status string) (*model.Order, error)
	}

	OrderRepositoryImpl struct {
		mu      sync.RWMutex
		list    *[]model.Order
		counter int64
	}
)

func NewOrderRepository(list *[]model.Order, counter int64) OrderRepository {
	return &OrderRepositoryImpl{
		list:    list,
		counter: counter,
	}
}

func (r *OrderRepositoryImpl) CreateOrder(order *model.Order) *model.Order {
	r.mu.Lock()
	defer r.mu.Unlock()

	r.counter++
	order.ID = r.counter
	*r.list = append(*r.list, *order)

	return &(*r.list)[len(*r.list)-1]
}

func (r *OrderRepositoryImpl) UpdateOrder(id int64, status string) (*model.Order, error) {
	r.mu.Lock()
	defer r.mu.Unlock()

	for i := range *r.list {
		if (*r.list)[i].ID == id {
			(*r.list)[i].Status = status
			return &(*r.list)[i], nil
		}
	}

	return nil, e.ErrOrderNotFound
}
