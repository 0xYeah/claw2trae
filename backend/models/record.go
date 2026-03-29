package models

import (
	"time"
)

type InboundRecord struct {
	ID                uint           `gorm:"primaryKey" json:"id"`
	InboundNo         string         `gorm:"uniqueIndex;size:50;not null" json:"inbound_no"`
	ProductID         uint           `gorm:"not null" json:"product_id"`
	WarehouseID       uint           `gorm:"not null" json:"warehouse_id"`
	StorageLocationID uint           `gorm:"not null" json:"storage_location_id"`
	Quantity          float64        `gorm:"type:decimal(10,2);not null" json:"quantity"`
	SourceInfo        string         `gorm:"size:255" json:"source_info"`
	Operator          string         `gorm:"size:100" json:"operator"`
	CreatedAt         time.Time      `json:"created_at"`
	Product           Product        `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Warehouse         Warehouse      `gorm:"foreignKey:WarehouseID" json:"warehouse,omitempty"`
	StorageLocation   StorageLocation `gorm:"foreignKey:StorageLocationID" json:"storage_location,omitempty"`
}

type InboundCreate struct {
	ProductID         uint    `json:"product_id" binding:"required"`
	WarehouseID       uint    `json:"warehouse_id" binding:"required"`
	StorageLocationID uint    `json:"storage_location_id" binding:"required"`
	Quantity          float64 `json:"quantity" binding:"required,gt=0"`
	SourceInfo        string  `json:"source_info"`
	Operator          string  `json:"operator"`
}

type OutboundRecord struct {
	ID                uint           `gorm:"primaryKey" json:"id"`
	OutboundNo        string         `gorm:"uniqueIndex;size:50;not null" json:"outbound_no"`
	ProductID         uint           `gorm:"not null" json:"product_id"`
	WarehouseID       uint           `gorm:"not null" json:"warehouse_id"`
	StorageLocationID uint           `gorm:"not null" json:"storage_location_id"`
	Quantity          float64        `gorm:"type:decimal(10,2);not null" json:"quantity"`
	DestinationInfo   string         `gorm:"size:255" json:"destination_info"`
	Operator          string         `gorm:"size:100" json:"operator"`
	CreatedAt         time.Time      `json:"created_at"`
	Product           Product        `gorm:"foreignKey:ProductID" json:"product,omitempty"`
	Warehouse         Warehouse      `gorm:"foreignKey:WarehouseID" json:"warehouse,omitempty"`
	StorageLocation   StorageLocation `gorm:"foreignKey:StorageLocationID" json:"storage_location,omitempty"`
}

type OutboundCreate struct {
	ProductID         uint    `json:"product_id" binding:"required"`
	WarehouseID       uint    `json:"warehouse_id" binding:"required"`
	StorageLocationID uint    `json:"storage_location_id" binding:"required"`
	Quantity          float64 `json:"quantity" binding:"required,gt=0"`
	DestinationInfo   string  `json:"destination_info"`
	Operator          string  `json:"operator"`
}
