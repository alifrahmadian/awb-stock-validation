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
			s.awbStockRepository.UpdateAWBStatus(orderModel.AWBNumber)

			return orderModel, nil
		}
	}

	// if awbStock is available in DB
	// do the awb validation based on convention
	// if valid, create the AWB and set the status "in_use"

	// validate awb
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
