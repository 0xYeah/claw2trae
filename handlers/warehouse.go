package handlers

import (
	"claw2trae/backend/database"
	"claw2trae/backend/models"
	"claw2trae/backend/utils"

	"github.com/gin-gonic/gin"
)

func GetWarehouses(c *gin.Context) {
	var warehouses []models.Warehouse
	if err := database.GetDB().Find(&warehouses).Error; err != nil {
		utils.Error(c, 500, "Failed to fetch warehouses")
		return
	}
	utils.Success(c, warehouses)
}

func GetWarehouse(c *gin.Context) {
	id := c.Param("id")
	var warehouse models.Warehouse
	if err := database.GetDB().First(&warehouse, id).Error; err != nil {
		utils.Error(c, 404, "Warehouse not found")
		return
	}
	utils.Success(c, warehouse)
}

func CreateWarehouse(c *gin.Context) {
	var input models.WarehouseCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, 400, err.Error())
		return
	}
	warehouse := models.Warehouse{
		Code:    input.Code,
		Name:    input.Name,
		Address: input.Address,
		Area:    input.Area,
		Manager: input.Manager,
		Status:  input.Status,
	}
	if warehouse.Status == 0 {
		warehouse.Status = 1
	}
	if err := database.GetDB().Create(&warehouse).Error; err != nil {
		utils.Error(c, 500, "Failed to create warehouse")
		return
	}
	utils.Success(c, warehouse)
}

func UpdateWarehouse(c *gin.Context) {
	id := c.Param("id")
	var input models.WarehouseUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, 400, err.Error())
		return
	}
	var warehouse models.Warehouse
	if err := database.GetDB().First(&warehouse, id).Error; err != nil {
		utils.Error(c, 404, "Warehouse not found")
		return
	}
	updates := map[string]interface{}{}
	if input.Name != "" {
		updates["name"] = input.Name
	}
	if input.Address != "" {
		updates["address"] = input.Address
	}
	if input.Area != 0 {
		updates["area"] = input.Area
	}
	if input.Manager != "" {
		updates["manager"] = input.Manager
	}
	if input.Status != 0 {
		updates["status"] = input.Status
	}
	if err := database.GetDB().Model(&warehouse).Updates(updates).Error; err != nil {
		utils.Error(c, 500, "Failed to update warehouse")
		return
	}
	utils.Success(c, warehouse)
}

func DeleteWarehouse(c *gin.Context) {
	id := c.Param("id")
	if err := database.GetDB().Delete(&models.Warehouse{}, id).Error; err != nil {
		utils.Error(c, 500, "Failed to delete warehouse")
		return
	}
	utils.SuccessWithMsg(c, "Warehouse deleted successfully", nil)
}
