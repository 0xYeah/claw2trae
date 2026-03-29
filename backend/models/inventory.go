package models

import (
	"time"
)

type Inventory struct {
	ID               uint           `gorm:"primaryKey" json:"id"`
	WarehouseID      uint           `gorm:"not null" json:"warehouse_id"`
	StorageLocationID uint           `gorm:"not null" json:"storage_location_id"`
	ProductID        uint           `gorm:"not null" json:"product_id"`
	Quantity         float64        `gorm:"type:decimal(10,2);default:0" json:"quantity"`
	UpdatedAt        time.Time      `json:"updated_at"`
	Warehouse        Warehouse      `gorm:"foreignKey:WarehouseID" json:"warehouse,omitempty"`
	StorageLocation  StorageLocation `gorm:"foreignKey:StorageLocationID" json:"storage_location,omitempty"`
	Product          Product        `gorm:"foreignKey:ProductID" json:"product,omitempty"`
}

type InventoryQuery struct {
	WarehouseID      uint `form:"warehouse_id"`
	StorageLocationID uint `form:"storage_location_id"`
	ProductID        uint `form:"product_id"`
}

type InventorySummary struct {
	TotalWarehouses  int64   `json:"total_warehouses"`
	TotalProducts    int64   `json:"total_products"`
	TotalQuantity    float64 `json:"total_quantity"`
	TotalInbounds   int64   `json:"total_inbounds"`
	TotalOutbounds  int64   `json:"total_outbounds"`
}
