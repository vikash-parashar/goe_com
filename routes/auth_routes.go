package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/ecommerce/controllers"
)

func SetupUserRoutes(r *gin.Engine) {
	authGroup := r.Group("/auth")
	{
		authGroup.POST("/register", controllers.Register)
		authGroup.POST("/login", controllers.Login)
	}
}
