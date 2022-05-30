package repositories

import (
	"assignment2/models"

	"gorm.io/gorm"
)

type OrderRepo interface {
	CreateOrder(order *models.Order) (uint, error)
	GetOrders() (*[]models.Order, error)
	GetOrdersById(orderId uint) (*models.Order, error)
	DeleteOrders(orderId uint) error
	UpdateOrders(payload *models.Order) error
}

type orderRepo struct {
	db *gorm.DB
}

func NewOrderRepo(db *gorm.DB) OrderRepo {
	return &orderRepo{db}
}

func (o *orderRepo) CreateOrder(order *models.Order) (uint, error) {
	err := o.db.Create(order).Error
	if err != nil {
		return 0, err
	}
	return order.ID, err
}

func (o *orderRepo) GetOrders() (*[]models.Order, error) {
	var orders []models.Order
	err := o.db.Find(&orders).Error
	return &orders, err
}

func (o *orderRepo) GetOrdersById(orderId uint) (*models.Order, error) {
	var orders models.Order
	err := o.db.First(&orders, "id=?", orderId).Error
	return &orders, err
}

func (o *orderRepo) DeleteOrders(orderId uint) error {
	var orders []models.Order
	return o.db.Delete(&orders, "id=?", orderId).Error
}

func (o *orderRepo) UpdateOrders(payload *models.Order) error {
	var orders models.Order
	return o.db.Model(&orders).Where("id=?", payload.ID).Updates(models.Order{CustomerName: payload.CustomerName}).Error
}
