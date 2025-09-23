package repository

import (
	"strings"
	"sync"

	"github.com/audricimanuel/awb-stock-allocation/src/model"
	e "github.com/audricimanuel/awb-stock-allocation/utils/errors"
)

// TODO: the interface and implementation under here
type (
	OrderRepository interface {
		CreateOrder(order *model.Order) *model.Order
		GetOrderById(id int64) (*model.Order, error)
		GetOrders(page int, awbNumber string) ([]*model.Order, bool)
		UpdateOrderStatus(id int64, status string) (*model.Order, error)
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

func (r *OrderRepositoryImpl) GetOrderById(id int64) (*model.Order, error) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	for i := range *r.list {
		if (*r.list)[i].ID == id {
			return &(*r.list)[i], nil
		}
	}

	return nil, e.ErrOrderNotFound
}

func (r *OrderRepositoryImpl) UpdateOrderStatus(id int64, status string) (*model.Order, error) {
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

func (r *OrderRepositoryImpl) GetOrders(page int, awbNumber string) ([]*model.Order, bool) {
	r.mu.RLock()
	defer r.mu.RUnlock()

	var filtered []*model.Order
	for _, order := range *r.list {
		if awbNumber == "" || strings.Contains(order.AWBNumber, awbNumber) {
			filtered = append(filtered, &order)
		}
	}

	pageSize := 5
	start := (page - 1) * pageSize
	if start >= len(filtered) {
		return nil, false
	}

	end := start + pageSize
	if end > len(filtered) {
		end = len(filtered)
	}

	hasNext := end < len(filtered)
	result := make([]*model.Order, end-start)
	copy(result, filtered[start:end])

	return result, hasNext
}
