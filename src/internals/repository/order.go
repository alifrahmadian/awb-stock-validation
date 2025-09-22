package repository

import (
	"sync"

	"github.com/audricimanuel/awb-stock-allocation/src/model"
)

// TODO: the interface and implementation under here
type (
	OrderRepository interface {
		CreateOrder(order *model.Order) *model.Order
		// UpdateOrder(awbNumber string) *model.Order
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
	r.mu.RLock()
	defer r.mu.RUnlock()

	r.counter++
	order.ID = r.counter
	*r.list = append(*r.list, *order)

	return order
}

// func (r *OrderRepositoryImpl) UpdateOrder(awbNumber string) *model.Order {
// 	r.mu.Lock()
// 	defer r.mu.RUnlock()
// }
