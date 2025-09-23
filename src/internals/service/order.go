package service

import (
	"github.com/audricimanuel/awb-stock-allocation/src/internals/repository"
	"github.com/audricimanuel/awb-stock-allocation/src/internals/utils"
	"github.com/audricimanuel/awb-stock-allocation/src/model"
	"github.com/audricimanuel/awb-stock-allocation/utils/constants"
	e "github.com/audricimanuel/awb-stock-allocation/utils/errors"
)

// TODO: the interface and implementation under here
type (
	OrderService interface {
		CreateOrder(order *model.Order) (*model.Order, error)
		UpdateOrderStatus(id int64, status int) (*model.Order, error)
		GetOrderById(id int64) (*model.Order, error)
		GetOrders(page int, awbNumber string) ([]*model.Order, bool)
	}

	OrderServiceImpl struct {
		orderRepository    repository.OrderRepository
		awbStockRepository repository.AWBStockRepository
	}
)

func NewOrderService(
	orderRepository repository.OrderRepository,
	awbStockRepository repository.AWBStockRepository,
) OrderService {
	return &OrderServiceImpl{
		orderRepository:    orderRepository,
		awbStockRepository: awbStockRepository,
	}
}

func (s *OrderServiceImpl) CreateOrder(order *model.Order) (*model.Order, error) {

	// calculate total price
	totalPrice := utils.CalculateTotalPrice(order.TotalWeight)

	order = &model.Order{
		AWBNumber:   order.AWBNumber,
		Sender:      order.Sender,
		Receiver:    order.Receiver,
		TotalWeight: order.TotalWeight,
		TotalPrice:  totalPrice,
		Status:      constants.ORDER_STATUS_PENDING,
	}

	awbStock := s.awbStockRepository.GetAWBStockByAWBNumber(order.AWBNumber)
	if awbStock != nil {
		if awbStock.Status == constants.AWB_STATUS_IN_USE {
			return nil, e.ErrAWBHasBeenUsed
		} else {
			orderModel := s.orderRepository.CreateOrder(order)
			s.awbStockRepository.UpdateAWBStatus(orderModel.AWBNumber, constants.AWB_STATUS_IN_USE)

			return orderModel, nil
		}
	}

	isAWBValid := utils.ValidateAWBNumber(order.AWBNumber)
	if !isAWBValid {
		return nil, e.ErrAWBInvalid
	}

	awbModel := &model.AWBStock{
		AWBNumber: order.AWBNumber,
		Status:    constants.AWB_STATUS_IN_USE,
	}

	s.awbStockRepository.CreateAWBStock(awbModel)
	orderModel := s.orderRepository.CreateOrder(order)

	return orderModel, nil
}

func (s *OrderServiceImpl) UpdateOrderStatus(id int64, status int) (*model.Order, error) {
	orderStatus, err := utils.MapInputStatusToString(status)
	if err != nil {
		return nil, err
	}

	order, err := s.orderRepository.GetOrderById(id)
	if err != nil {
		return nil, err
	}

	switch order.Status {
	case constants.ORDER_STATUS_PENDING:
		if orderStatus != constants.ORDER_STATUS_CONFIRM && orderStatus != constants.ORDER_STATUS_CANCELLED {
			return nil, e.ErrOrderStatusPendingValidation
		}
	case constants.ORDER_STATUS_CONFIRM:
		if orderStatus != constants.ORDER_STATUS_SHIPPED && orderStatus != constants.ORDER_STATUS_CANCELLED {
			return nil, e.ErrOrderStatusConfirmValidation
		}
	case constants.ORDER_STATUS_SHIPPED:
		if orderStatus != constants.ORDER_STATUS_COMPLETED {
			return nil, e.ErrOrderStatusShippedValidation
		}
	case constants.ORDER_STATUS_COMPLETED, constants.ORDER_STATUS_CANCELLED:
		return nil, e.ErrOrderStatusFinal
	}

	updatedOrder, err := s.orderRepository.UpdateOrderStatus(id, orderStatus)
	if err != nil {
		return nil, err
	}

	if updatedOrder.Status == constants.ORDER_STATUS_CANCELLED {
		s.awbStockRepository.UpdateAWBStatus(updatedOrder.AWBNumber, constants.AWB_STATUS_NOT_IN_USE)
	}

	return updatedOrder, nil
}

func (s *OrderServiceImpl) GetOrderById(id int64) (*model.Order, error) {
	order, err := s.orderRepository.GetOrderById(id)
	if err != nil {
		return nil, err
	}

	return order, nil
}

func (s *OrderServiceImpl) GetOrders(page int, awbNumber string) ([]*model.Order, bool) {
	orders, hasNext := s.orderRepository.GetOrders(page, awbNumber)
	return orders, hasNext
}
