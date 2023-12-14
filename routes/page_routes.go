package routes

import (
	"github.com/gin-gonic/gin"
	// Import other necessary packages
)

func SetupDisplayRoutes(r *gin.Engine) {
	// Define routes for display pages
	pageGroup := r.Group("/page")
	{
		pageGroup.GET("/get/list-all", controllers.RenderIndexPage)

		// Add other product routes
	}
}
