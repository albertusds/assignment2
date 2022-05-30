package controllers

import (
	"assignment2/params"
	"assignment2/services"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type OrderController struct {
	orderService services.OrderService
	itemService  services.ItemService
}

func NewOrderController(serviceOrder *services.OrderService, serviceItem *services.ItemService) *OrderController {
	return &OrderController{
		orderService: *serviceOrder,
		itemService:  *serviceItem,
	}
}

func (o *OrderController) CreateOrder(c *gin.Context) {
	var req params.CreateOrder

	err := c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Message: "BAD REQUEST",
			Error:   err.Error(),
		})
		return
	}
	response, orderId, err := o.orderService.CreateOrder(req)
	if err != nil {
		c.JSON(response.Status, response)
		return
	}

	response = o.itemService.CreateItem(req.Items, orderId)
	c.JSON(response.Status, response)
}

func (o *OrderController) GetOrders(c *gin.Context) {

	_, orders, err := o.orderService.GetOrders()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, params.Response{
			Status:  http.StatusInternalServerError,
			Message: "INTERNAL SERVER ERROR",
			Error:   err.Error(),
		})
		return
	}

	response := o.itemService.GetItemByOrderId(orders)

	c.JSON(response.Status, response)
}

func (o *OrderController) DeleteOrders(c *gin.Context) {
	orderIdString := c.Param("orderId")
	orderId, err := strconv.Atoi(orderIdString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: err.Error(),
		})
	}

	_, err = o.orderService.DeleteOrders(uint(orderId))
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, params.Response{
			Status:  http.StatusInternalServerError,
			Message: "INTERNAL SERVER ERROR",
			Error:   err.Error(),
		})
		return
	}

	response := o.itemService.DeleteItemByOrderId(uint(orderId))

	c.JSON(response.Status, response)
}

func (o *OrderController) UpdateOrders(c *gin.Context) {
	orderIdString := c.Param("orderId")
	orderId, err := strconv.Atoi(orderIdString)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:         http.StatusBadRequest,
			Error:          "BAD REQUEST",
			AdditionalInfo: err.Error(),
		})
		return
	}

	var req params.UpdateOrder

	err = c.ShouldBindJSON(&req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, params.Response{
			Status:  http.StatusBadRequest,
			Message: "BAD REQUEST",
			Error:   err.Error(),
		})
		return
	}

	_, err = o.orderService.UpdateOrders(req)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, params.Response{
			Status:  http.StatusInternalServerError,
			Message: "INTERNAL SERVER ERROR",
			Error:   err.Error(),
		})
		return
	}

	response := o.itemService.UpdateItemByOID(&req, uint(orderId))

	c.JSON(response.Status, response)
}
