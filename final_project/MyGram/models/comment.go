package models

import "time"

type Comments struct {
	IdComment uint `gorm:"primary_key;auto_increment:true" json:"idComment"`
	UserID    uint
	PhotoID   uint
	Message   string    `gorm:"not null" json:"message" valid:"required"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
}
