package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vikash-parashar/goe_com/controllers"
)

func SetupProductRoutes(r *gin.Engine) {
	productGroup := r.Group("/product")
	{
		productGroup.GET("/get/list-all", controllers.GetProducts)
		productGroup.GET("/get/{id}", controllers.GetProductByProductId)
		productGroup.POST("/new", controllers.CreateNewProduct)
		productGroup.PATCH("/update/{id}", controllers.UpdateProduct)
		productGroup.DELETE("/delete/{id}", controllers.DeleteProduct)
		// Add other product routes
	}
}
