package services

import (
	"assignment2/models"
	"assignment2/params"
	"assignment2/repositories"
	"net/http"
)

type ItemService struct {
	itemRepo repositories.ItemRepo
}

func NewItemService(repo repositories.ItemRepo) *ItemService {
	return &ItemService{
		itemRepo: repo,
	}
}

func (i *ItemService) CreateItem(request []params.CreateItem, orderId uint) *params.Response {

	for _, rq := range request {
		model := models.Item{
			ItemCode:    rq.ItemCode,
			Description: rq.Description,
			Quantity:    rq.Quantity,
			OrderID:     orderId,
		}

		err := i.itemRepo.CreateItem(&model)
		if err != nil {
			return &params.Response{
				Status:  http.StatusInternalServerError,
				Message: "INTERNAL SERVER ERROR",
				Error:   err.Error(),
			}
		}
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "SUCCESS",
		Payload: request,
	}
}

func (i *ItemService) GetItemByOrderId(request *[]models.Order) *params.Response {

	//create var to save list item per order id
	var response = []params.ReponseGetOrder{}

	for _, rq := range *request {

		listItem, err := i.itemRepo.GetItemByOrderId(rq.ID)
		if err != nil {
			return &params.Response{
				Status:  http.StatusInternalServerError,
				Message: "INTERNAL SERVER ERROR",
				Error:   err.Error(),
			}
		}

		responseItem := params.ReponseGetOrder{
			OrderId:      rq.ID,
			CustomerName: rq.CustomerName,
			OrderedAt:    rq.OrderedAt,
			Items:        *listItem,
		}

		response = append(response, responseItem)
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "SUCCESS",
		Payload: response,
	}
}

func (i *ItemService) DeleteItemByOrderId(orderId uint) *params.Response {

	var responseDelete params.ReponseDeleteOrder
	err := i.itemRepo.DeleteItemByOrderId(orderId)
	if err != nil {
		return &params.Response{
			Status:  http.StatusInternalServerError,
			Message: "INTERNAL SERVER ERROR",
			Error:   err.Error(),
		}
	}
	responseDelete = params.ReponseDeleteOrder{
		OrderId: orderId,
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "SUCCESS DELETE",
		Payload: responseDelete,
	}
}

func (i *ItemService) UpdateItemByOID(req *params.UpdateOrder, orderId uint) *params.Response {

	var responseUpdate []params.CreateItem

	for _, rq := range req.Items {
		var updateItemReq = models.Item{
			ItemCode:    rq.ItemCode,
			Description: rq.Description,
			Quantity:    rq.Quantity,
			OrderID:     orderId,
		}

		err := i.itemRepo.UpdateItemByOrderId(&updateItemReq)
		if err != nil {
			return &params.Response{
				Status:  http.StatusInternalServerError,
				Message: "INTERNAL SERVER ERROR",
				Error:   err.Error(),
			}
		}
		responseUpdate = append(responseUpdate, params.CreateItem{ItemCode: rq.ItemCode, Description: rq.Description, Quantity: rq.Quantity})
	}

	var responseUpdateAll = params.UpdateOrder{
		CustomerName: req.CustomerName,
		Items:        responseUpdate,
	}

	return &params.Response{
		Status:  http.StatusOK,
		Message: "SUCCESS UPDATE",
		Payload: responseUpdateAll,
	}
}
