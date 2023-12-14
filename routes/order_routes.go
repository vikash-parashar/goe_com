package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/ecommerce/controllers"
)

func SetupOrderRoutes(r *gin.Engine) {
	orderGroup := r.Group("/order")
	{
		orderGroup.POST("/all", controllers.Register)
		orderGroup.POST("/{orderId}", controllers.Login)
	}
}
