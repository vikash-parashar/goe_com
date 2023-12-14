package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/vikash-parashar/goe_com/middleware"
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
