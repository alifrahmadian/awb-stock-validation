package controller

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"

	"github.com/audricimanuel/awb-stock-allocation/src/internals/dto"
	"github.com/audricimanuel/awb-stock-allocation/src/internals/service"
	"github.com/audricimanuel/awb-stock-allocation/src/model"
	e "github.com/audricimanuel/awb-stock-allocation/utils/errors"
	"github.com/audricimanuel/awb-stock-allocation/utils/httputils"
	"github.com/go-chi/chi/v5"
	"github.com/sirupsen/logrus"
)

// TODO: the interface and implementation under here
type (
	OrderController interface {
		CreateOrder(w http.ResponseWriter, r *http.Request)
		UpdateOrderStatus(w http.ResponseWriter, r *http.Request)
	}

	OrderControllerImpl struct {
		orderService service.OrderService
		logger       *logrus.Logger
	}
)

func NewOrderController(orderService service.OrderService, logger *logrus.Logger) OrderController {
	return &OrderControllerImpl{
		orderService: orderService,
		logger:       logger,
	}
}

// @Tags Order
// @Summary Create Order
// @Description "Create an Order"
// @Accept json
// @Produce json
// @Success 200 {object} httputils.BaseResponse
// @Router /order [post]
func (a *OrderControllerImpl) CreateOrder(w http.ResponseWriter, r *http.Request) {
	var req *dto.CreateOrderRequest

	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		a.logger.WithError(err).Warn("failed to decode request body")
		httputils.MapBaseResponse(w, r, nil, err, nil)
		return
	}

	if req.AWBNumber == "" {
		err = e.ErrOrderAWBRequired

		a.logger.WithError(err).Warn(err.Error())
		httputils.MapBaseResponse(w, r, nil, err, nil)
		return
	}

	if req.Sender == "" {
		err = e.ErrOrderSenderRequired

		a.logger.WithError(err).Warn(err.Error())
		httputils.MapBaseResponse(w, r, nil, err, nil)
		return
	}

	if req.Receiver == "" {
		err = e.ErrOrderReceiverRequired

		a.logger.WithError(err).Warn(err.Error())
		httputils.MapBaseResponse(w, r, nil, err, nil)
		return
	}

	if req.TotalWeight == 0 {
		err = e.ErrOrderTotalWeightRequired

		a.logger.WithError(err).Warn(err.Error())
		httputils.MapBaseResponse(w, r, nil, err, nil)
		return
	}

	order := &model.Order{
		AWBNumber:   req.AWBNumber,
		Sender:      req.Sender,
		Receiver:    req.Receiver,
		TotalWeight: req.TotalWeight,
	}

	a.logger.WithFields(logrus.Fields{
		"awb_number":   order.AWBNumber,
		"sender":       order.Sender,
		"receiver":     order.Receiver,
		"total_weight": order.TotalWeight,
	}).Info("creating order")

	orderModel, err := a.orderService.CreateOrder(order)
	if err != nil {
		a.logger.WithError(err).Warn(err.Error())
		httputils.MapBaseResponse(w, r, nil, err, nil)
		return
	}

	resp := &dto.CreateOrderResponse{
		ID:          orderModel.ID,
		AWBNumber:   orderModel.AWBNumber,
		Sender:      orderModel.Sender,
		Receiver:    orderModel.Receiver,
		TotalWeight: orderModel.TotalWeight,
		TotalPrice:  orderModel.TotalPrice,
		Status:      orderModel.Status,
	}

	a.logger.WithField("awb_number", orderModel.ID).Info("order created successfully")

	meta := httputils.SetBaseMeta(1, 10, 100)

	httputils.MapBaseResponse(w, r, resp, err, &meta)
}

func (a *OrderControllerImpl) UpdateOrderStatus(w http.ResponseWriter, r *http.Request) {
	var req *dto.UpdateOrderStatusRequest

	idStr := chi.URLParam(r, "id")
	id, err := strconv.ParseInt(idStr, 10, 64)
	if err != nil {
		err = errors.New("invalid order id")

		a.logger.WithError(err).Error(err.Error())
		httputils.MapBaseResponse(w, r, nil, err, nil)
		return
	}

	err = json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		a.logger.WithError(err).Error(err.Error())
		httputils.MapBaseResponse(w, r, nil, err, nil)
		return
	}

	a.logger.WithFields(logrus.Fields{
		"id":     id,
		"status": req.Status,
	}).Info("updating order status")

	order, err := a.orderService.UpdateOrderStatus(id, req.Status)
	if err != nil {
		a.logger.WithError(err).Error(err.Error())
		httputils.MapBaseResponse(w, r, nil, err, nil)
		return
	}

	resp := &dto.CreateOrderResponse{
		ID:          order.ID,
		AWBNumber:   order.AWBNumber,
		Sender:      order.Sender,
		Receiver:    order.Receiver,
		TotalWeight: order.TotalWeight,
		TotalPrice:  order.TotalPrice,
		Status:      order.Status,
	}

	a.logger.WithField("awb_number", order.ID).Info("order status updated successfully")
	httputils.MapBaseResponse(w, r, resp, nil, nil)
}
