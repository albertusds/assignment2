package models

type Item struct {
	ID          uint   `gorm:"primaryKey" json:"item_id"`
	ItemCode    string `json:"item_code"`
	Description string `json:"description"`
	Quantity    int    `json:"quantity"`
	OrderID     uint   `json:"order_id"`
}
