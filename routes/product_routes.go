package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/ecommerce/controllers"
)

func SetupProductRoutes(r *gin.Engine) {
	productGroup := r.Group("/product")
	{
		productGroup.GET("/list", controllers.GetProducts)
		// Add other product routes
	}
}
