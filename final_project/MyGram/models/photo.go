package models

import "time"

type Photo struct {
	IdPhoto   uint   `gorm:"primary_key;auto_increment:true" json:"idPhoto"`
	Title     string `gorm:"not null" json:"title" valid:"required"`
	Caption   string
	PhotoURL  string `gorm:"not null" json:"photo_url" valid:"required"`
	UserID    uint
	CreatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"createdAt"`
	UpdatedAt time.Time  `gorm:"default:CURRENT_TIMESTAMP" json:"updatedAt"`
	Comments  []Comments `grom:"constraint:OnUpdate:CASCADE,OnDelete: SET NULL" json:"comments"`
}
