package models

import (
	"time"

	"gorm.io/gorm"
)

type Warehouse struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Code      string         `gorm:"uniqueIndex;size:50;not null" json:"code"`
	Name      string         `gorm:"size:100;not null" json:"name"`
	Address   string         `gorm:"size:255" json:"address"`
	Area      float64        `gorm:"type:decimal(10,2)" json:"area"`
	Manager   string         `gorm:"size:100" json:"manager"`
	Status    int            `gorm:"default:1" json:"status"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type WarehouseCreate struct {
	Code    string  `json:"code" binding:"required"`
	Name    string  `json:"name" binding:"required"`
	Address string  `json:"address"`
	Area    float64 `json:"area"`
	Manager string  `json:"manager"`
	Status  int     `json:"status"`
}

type WarehouseUpdate struct {
	Name    string  `json:"name"`
	Address string  `json:"address"`
	Area    float64 `json:"area"`
	Manager string  `json:"manager"`
	Status  int     `json:"status"`
}
