package models

import "time"

type Item struct {
	ItemID      int    `json:"item_id" gorm:"primary_key;auto_increment:true"`
	ItemCode    string `json:"item_code" gorm:"not null;type:varchar(150)"`
	Description string `json:"description" gorm:"not null; type:text"`
	Quantity    int    `json:"quantity"`
	OrderID     int    `json:"order_id" grom:"index"`
	Items       []Item `json:"item" gorm:"foreignKey:order_id"`
}

type Order struct {
	OrderID      int       `json:"order_id" gorm:"primary_key;auto_increment:true"`
	CustomerName string    `json:"customer_name"`
	OrderedAt    time.Time `json:"ordered_at"`
}
