package handlers

import (
	"claw2trae/backend/database"
	"claw2trae/backend/models"
	"claw2trae/backend/utils"

	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	var products []models.Product
	query := database.GetDB()
	if category := c.Query("category"); category != "" {
		query = query.Where("category = ?", category)
	}
	if err := query.Find(&products).Error; err != nil {
		utils.Error(c, 500, "Failed to fetch products")
		return
	}
	utils.Success(c, products)
}

func GetProduct(c *gin.Context) {
	id := c.Param("id")
	var product models.Product
	if err := database.GetDB().First(&product, id).Error; err != nil {
		utils.Error(c, 404, "Product not found")
		return
	}
	utils.Success(c, product)
}

func CreateProduct(c *gin.Context) {
	var input models.ProductCreate
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, 400, err.Error())
		return
	}
	product := models.Product{
		Code:     input.Code,
		Name:     input.Name,
		Spec:     input.Spec,
		Unit:     input.Unit,
		Category: input.Category,
	}
	if err := database.GetDB().Create(&product).Error; err != nil {
		utils.Error(c, 500, "Failed to create product")
		return
	}
	utils.Success(c, product)
}

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var input models.ProductUpdate
	if err := c.ShouldBindJSON(&input); err != nil {
		utils.Error(c, 400, err.Error())
		return
	}
	var product models.Product
	if err := database.GetDB().First(&product, id).Error; err != nil {
		utils.Error(c, 404, "Product not found")
		return
	}
	updates := map[string]interface{}{}
	if input.Name != "" {
		updates["name"] = input.Name
	}
	if input.Spec != "" {
		updates["spec"] = input.Spec
	}
	if input.Unit != "" {
		updates["unit"] = input.Unit
	}
	if input.Category != "" {
		updates["category"] = input.Category
	}
	if err := database.GetDB().Model(&product).Updates(updates).Error; err != nil {
		utils.Error(c, 500, "Failed to update product")
		return
	}
	utils.Success(c, product)
}

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	if err := database.GetDB().Delete(&models.Product{}, id).Error; err != nil {
		utils.Error(c, 500, "Failed to delete product")
		return
	}
	utils.SuccessWithMsg(c, "Product deleted successfully", nil)
}
