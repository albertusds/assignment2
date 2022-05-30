package repositories

import (
	"assignment2/models"

	"gorm.io/gorm"
)

type ItemRepo interface {
	CreateItem(order *models.Item) error
	GetItemByOrderId(orderId uint) (*[]models.Item, error)
	DeleteItemByOrderId(orderId uint) error
	UpdateItemByOrderId(payload *models.Item) error
}

type itemRepo struct {
	db *gorm.DB
}

func NewItemRepo(db *gorm.DB) ItemRepo {
	return &itemRepo{db}
}

func (i *itemRepo) CreateItem(item *models.Item) error {
	return i.db.Create(item).Error
}

func (i *itemRepo) GetItemByOrderId(orderId uint) (*[]models.Item, error) {
	var listItem []models.Item

	err := i.db.Where("order_id=?", orderId).Find(&listItem).Error
	if err != nil {
		return nil, err
	}

	return &listItem, err
}

func (i *itemRepo) DeleteItemByOrderId(orderId uint) error {
	var deleteItem models.Item

	err := i.db.Where("order_id=?", orderId).Delete(&deleteItem).Error
	return err
}

func (i *itemRepo) UpdateItemByOrderId(payload *models.Item) error {
	var updateItem models.Item
	err := i.db.Model(&updateItem).Where("order_id=?", payload.OrderID).Where("item_code=?", payload.ItemCode).Updates(models.Item{Description: payload.Description, Quantity: payload.Quantity}).Error
	return err
}
