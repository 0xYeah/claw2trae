package handlers

import (
	"claw2trae/backend/database"
	"claw2trae/backend/models"
	"claw2trae/backend/utils"

	"github.com/gin-gonic/gin"
)

func GetInventory(c *gin.Context) {
	var query models.InventoryQuery
	if err := c.ShouldBindQuery(&query); err != nil {
		utils.Error(c, 400, err.Error())
		return
	}

	db := database.GetDB()
	var inventories []models.Inventory

	q := db.Preload("Warehouse").Preload("StorageLocation").Preload("Product")
	if query.WarehouseID != 0 {
		q = q.Where("warehouse_id = ?", query.WarehouseID)
	}
	if query.StorageLocationID != 0 {
		q = q.Where("storage_location_id = ?", query.StorageLocationID)
	}
	if query.ProductID != 0 {
		q = q.Where("product_id = ?", query.ProductID)
	}

	if err := q.Find(&inventories).Error; err != nil {
		utils.Error(c, 500, "Failed to fetch inventory")
		return
	}
	utils.Success(c, inventories)
}

func GetInventorySummary(c *gin.Context) {
	var summary models.InventorySummary

	db := database.GetDB()
	db.Model(&models.Warehouse{}).Count(&summary.TotalWarehouses)
	db.Model(&models.Product{}).Count(&summary.TotalProducts)
	db.Model(&models.InboundRecord{}).Count(&summary.TotalInbounds)
	db.Model(&models.OutboundRecord{}).Count(&summary.TotalOutbounds)

	var totalQty struct {
		Total float64
	}
	db.Model(&models.Inventory{}).Select("COALESCE(SUM(quantity), 0) as total").Scan(&totalQty)
	summary.TotalQuantity = totalQty.Total

	utils.Success(c, summary)
}
