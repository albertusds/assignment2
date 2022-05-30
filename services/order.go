package services

import (
	"assignment2/models"
	"assignment2/params"
	"assignment2/repositories"
	"net/http"
	"time"
)

type OrderService struct {
	orderRepo repositories.OrderRepo
}

func NewOrderService(repo repositories.OrderRepo) *OrderService {
	return &OrderService{
		orderRepo: repo,
	}
}

func (o *OrderService) CreateOrder(request params.CreateOrder) (*params.Response, uint, error) {
	model := models.Order{
		CustomerName: request.CustomerName,
		OrderedAt:    time.Now(),
	}

	orderId, err := o.orderRepo.CreateOrder(&model)
	if err != nil {
		return &params.Response{
			Status:  http.StatusBadRequest,
			Message: "BAD REQUEST",
			Error:   err.Error(),
		}, orderId, err
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "SUCCESS",
		Payload: request,
	}, orderId, err
}

func (o *OrderService) GetOrders() (*params.Response, *[]models.Order, error) {
	orders, err := o.orderRepo.GetOrders()
	if err != nil {
		return &params.Response{
			Status:  http.StatusBadRequest,
			Message: "BAD REQUEST",
			Error:   err.Error(),
		}, nil, err
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "SUCCESS",
		Payload: orders,
	}, orders, err
}

func (o *OrderService) DeleteOrders(orderId uint) (*params.Response, error) {
	err := o.orderRepo.DeleteOrders(orderId)
	if err != nil {
		return &params.Response{
			Status:  http.StatusBadRequest,
			Message: "BAD REQUEST",
			Error:   err.Error(),
		}, err
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "SUCCESS",
		Payload: params.ReponseDeleteOrder{
			OrderId: orderId,
		},
	}, err
}

func (o *OrderService) UpdateOrders(request params.UpdateOrder) (*params.Response, error) {
	var updateOrderReq = models.Order{
		ID:           request.ID,
		CustomerName: request.CustomerName,
	}

	err := o.orderRepo.UpdateOrders(&updateOrderReq)
	if err != nil {
		return &params.Response{
			Status:  http.StatusBadRequest,
			Message: "BAD REQUEST",
			Error:   err.Error(),
		}, err
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "SUCCESS",
		Payload: params.UpdateOrder{
			CustomerName: request.CustomerName,
		},
	}, err
}
