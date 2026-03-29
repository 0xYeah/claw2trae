package models

import (
	"time"

	"gorm.io/gorm"
)

type StorageLocation struct {
	ID           uint           `gorm:"primaryKey" json:"id"`
	WarehouseID  uint           `gorm:"not null" json:"warehouse_id"`
	Code         string         `gorm:"size:50;not null" json:"code"`
	Capacity     float64        `gorm:"type:decimal(10,2)" json:"capacity"`
	Status       int            `gorm:"default:1" json:"status"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	Warehouse    Warehouse      `gorm:"foreignKey:WarehouseID" json:"warehouse,omitempty"`
}

type LocationCreate struct {
	WarehouseID uint    `json:"warehouse_id" binding:"required"`
	Code       string  `json:"code" binding:"required"`
	Capacity   float64 `json:"capacity"`
	Status     int     `json:"status"`
}

type LocationUpdate struct {
	WarehouseID uint    `json:"warehouse_id"`
	Code       string  `json:"code"`
	Capacity   float64 `json:"capacity"`
	Status     int     `json:"status"`
}
