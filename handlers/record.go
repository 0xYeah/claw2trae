package handlers

import (
	"fmt"
	"time"

	"claw2trae/backend/database"
	"claw2trae/backend/models"
	"claw2trae/backend/utils"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetInbounds(c *gin.Context) {
	var inbounds []models.InboundRecord
	query := database.GetDB().Preload("Product").Preload("Warehouse").Preload("StorageLocation")

	if wid := c.Query("warehouse_id"); wid != "" {
		query = query.Where("warehouse_id = ?", wid)
	}
	if pid := c.Query("product_id"); pid != "" {
		query = query.Where("product_id = ?", pid)
	}

	if err := query.Order("created_at DESC").Find(&inbounds).Error; err != nil {
		utils.Error(c, 500, "Failed to fetch inbound records")
		return
	}
	utils.Success(c, inbounds)
}

func CreateInbound(c *gin.Context) {
	var input models.InboundCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, 400, err.Error())
		return
	}

	inbound := models.InboundRecord{
		InboundNo:         fmt.Sprintf("IN-%s", uuid.New().String()[:8]),
		ProductID:         input.ProductID,
		WarehouseID:       input.WarehouseID,
		StorageLocationID: input.StorageLocationID,
		Quantity:          input.Quantity,
		SourceInfo:        input.SourceInfo,
		Operator:          input.Operator,
		CreatedAt:         time.Now(),
	}

	tx := database.GetDB().Begin()

	if err := tx.Create(&inbound).Error; err != nil {
		tx.Rollback()
		utils.Error(c, 500, "Failed to create inbound record")
		return
	}

	var inventory models.Inventory
	err := tx.Where("warehouse_id = ? AND storage_location_id = ? AND product_id = ?",
		input.WarehouseID, input.StorageLocationID, input.ProductID).First(&inventory).Error

	if err != nil {
		inventory = models.Inventory{
			WarehouseID:       input.WarehouseID,
			StorageLocationID: input.StorageLocationID,
			ProductID:         input.ProductID,
			Quantity:          input.Quantity,
		}
		if err := tx.Create(&inventory).Error; err != nil {
			tx.Rollback()
			utils.Error(c, 500, "Failed to update inventory")
			return
		}
	} else {
		if err := tx.Model(&inventory).Update("quantity", inventory.Quantity+input.Quantity).Error; err != nil {
			tx.Rollback()
			utils.Error(c, 500, "Failed to update inventory")
			return
		}
	}

	tx.Commit()

	database.GetDB().Preload("Product").Preload("Warehouse").Preload("StorageLocation").First(&inbound, inbound.ID)
	utils.Success(c, inbound)
}

func GetOutbounds(c *gin.Context) {
	var outbounds []models.OutboundRecord
	query := database.GetDB().Preload("Product").Preload("Warehouse").Preload("StorageLocation")

	if wid := c.Query("warehouse_id"); wid != "" {
		query = query.Where("warehouse_id = ?", wid)
	}
	if pid := c.Query("product_id"); pid != "" {
		query = query.Where("product_id = ?", pid)
	}

	if err := query.Order("created_at DESC").Find(&outbounds).Error; err != nil {
		utils.Error(c, 500, "Failed to fetch outbound records")
		return
	}
	utils.Success(c, outbounds)
}

func CreateOutbound(c *gin.Context) {
	var input models.OutboundCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, 400, err.Error())
		return
	}

	tx := database.GetDB().Begin()

	var inventory models.Inventory
	err := tx.Where("warehouse_id = ? AND storage_location_id = ? AND product_id = ?",
		input.WarehouseID, input.StorageLocationID, input.ProductID).First(&inventory).Error

	if err != nil {
		tx.Rollback()
		utils.Error(c, 404, "Inventory not found for this location")
		return
	}

	if inventory.Quantity < input.Quantity {
		tx.Rollback()
		utils.Error(c, 400, "Insufficient inventory")
		return
	}

	outbound := models.OutboundRecord{
		OutboundNo:        fmt.Sprintf("OUT-%s", uuid.New().String()[:8]),
		ProductID:         input.ProductID,
		WarehouseID:       input.WarehouseID,
		StorageLocationID: input.StorageLocationID,
		Quantity:          input.Quantity,
		DestinationInfo:   input.DestinationInfo,
		Operator:          input.Operator,
		CreatedAt:         time.Now(),
	}

	if err := tx.Create(&outbound).Error; err != nil {
		tx.Rollback()
		utils.Error(c, 500, "Failed to create outbound record")
		return
	}

	if err := tx.Model(&inventory).Update("quantity", inventory.Quantity-input.Quantity).Error; err != nil {
		tx.Rollback()
		utils.Error(c, 500, "Failed to update inventory")
		return
	}

	tx.Commit()

	database.GetDB().Preload("Product").Preload("Warehouse").Preload("StorageLocation").First(&outbound, outbound.ID)
	utils.Success(c, outbound)
}
