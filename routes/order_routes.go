package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vikash-parashar/goe_com/controllers"
)

func SetupOrderRoutes(r *gin.Engine) {
	orderGroup := r.Group("/order")
	{
		orderGroup.POST("/all", controllers.Register)
		orderGroup.POST("/{orderId}", controllers.Login)
	}
}
