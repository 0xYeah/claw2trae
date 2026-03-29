package handlers

import (
	"claw2trae/backend/database"
	"claw2trae/backend/models"
	"claw2trae/backend/utils"

	"github.com/gin-gonic/gin"
)

func GetLocations(c *gin.Context) {
	var locations []models.StorageLocation
	query := database.GetDB()
	if wid := c.Query("warehouse_id"); wid != "" {
		query = query.Where("warehouse_id = ?", wid)
	}
	if err := query.Preload("Warehouse").Find(&locations).Error; err != nil {
		utils.Error(c, 500, "Failed to fetch locations")
		return
	}
	utils.Success(c, locations)
}

func GetLocation(c *gin.Context) {
	id := c.Param("id")
	var location models.StorageLocation
	if err := database.GetDB().Preload("Warehouse").First(&location, id).Error; err != nil {
		utils.Error(c, 404, "Location not found")
		return
	}
	utils.Success(c, location)
}

func CreateLocation(c *gin.Context) {
	var input models.LocationCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, 400, err.Error())
		return
	}
	location := models.StorageLocation{
		WarehouseID: input.WarehouseID,
		Code:       input.Code,
		Capacity:   input.Capacity,
		Status:     input.Status,
	}
	if location.Status == 0 {
		location.Status = 1
	}
	if err := database.GetDB().Create(&location).Error; err != nil {
		utils.Error(c, 500, "Failed to create location")
		return
	}
	database.GetDB().Preload("Warehouse").First(&location, location.ID)
	utils.Success(c, location)
}

func UpdateLocation(c *gin.Context) {
	id := c.Param("id")
	var input models.LocationUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, 400, err.Error())
		return
	}
	var location models.StorageLocation
	if err := database.GetDB().First(&location, id).Error; err != nil {
		utils.Error(c, 404, "Location not found")
		return
	}
	updates := map[string]interface{}{}
	if input.WarehouseID != 0 {
		updates["warehouse_id"] = input.WarehouseID
	}
	if input.Code != "" {
		updates["code"] = input.Code
	}
	if input.Capacity != 0 {
		updates["capacity"] = input.Capacity
	}
	if input.Status != 0 {
		updates["status"] = input.Status
	}
	if err := database.GetDB().Model(&location).Updates(updates).Error; err != nil {
		utils.Error(c, 500, "Failed to update location")
		return
	}
	database.GetDB().Preload("Warehouse").First(&location, location.ID)
	utils.Success(c, location)
}

func DeleteLocation(c *gin.Context) {
	id := c.Param("id")
	if err := database.GetDB().Delete(&models.StorageLocation{}, id).Error; err != nil {
		utils.Error(c, 500, "Failed to delete location")
		return
	}
	utils.SuccessWithMsg(c, "Location deleted successfully", nil)
}
