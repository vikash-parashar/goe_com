package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/yourusername/ecommerce/middleware"
)

func SetupRoutes() *gin.Engine {
	r := gin.Default()

	// Middleware
	r.Use(middleware.LoggerMiddleware)
	r.Use(middleware.AuthMiddleware)

	// Setup routes from individual files
	SetupUserRoutes(r)
	SetupProductRoutes(r)
	SetupDisplayRoutes(r)

	return r
}
