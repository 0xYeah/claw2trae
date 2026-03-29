package routes

import (
	"claw2trae/backend/handlers"

	"github.com/gin-gonic/gin"
)

func Setup(r *gin.Engine) {
	api := r.Group("/api")
	{
		api.GET("/ping", func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "pong"})
		})

		warehouses := api.Group("/warehouses")
		{
			warehouses.GET("", handlers.GetWarehouses)
			warehouses.GET("/:id", handlers.GetWarehouse)
			warehouses.POST("", handlers.CreateWarehouse)
			warehouses.PUT("/:id", handlers.UpdateWarehouse)
			warehouses.DELETE("/:id", handlers.DeleteWarehouse)
		}

		locations := api.Group("/locations")
		{
			locations.GET("", handlers.GetLocations)
			locations.GET("/:id", handlers.GetLocation)
			locations.POST("", handlers.CreateLocation)
			locations.PUT("/:id", handlers.UpdateLocation)
			locations.DELETE("/:id", handlers.DeleteLocation)
		}

		products := api.Group("/products")
		{
			products.GET("", handlers.GetProducts)
			products.GET("/:id", handlers.GetProduct)
			products.POST("", handlers.CreateProduct)
			products.PUT("/:id", handlers.UpdateProduct)
			products.DELETE("/:id", handlers.DeleteProduct)
		}

		inventory := api.Group("/inventory")
		{
			inventory.GET("", handlers.GetInventory)
			inventory.GET("/summary", handlers.GetInventorySummary)
		}

		inbounds := api.Group("/inbounds")
		{
			inbounds.GET("", handlers.GetInbounds)
			inbounds.POST("", handlers.CreateInbound)
		}

		outbounds := api.Group("/outbounds")
		{
			outbounds.GET("", handlers.GetOutbounds)
			outbounds.POST("", handlers.CreateOutbound)
		}
	}
}
