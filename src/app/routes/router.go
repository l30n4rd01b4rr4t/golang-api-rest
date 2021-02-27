package routes

import (
	"github.com/gin-gonic/gin"

	"github.com/l30n4rd01b4rr4t/golang-api-rest/src/app/handler"
)

func MapRoutes() *gin.Engine {

	// Create services.

	router := gin.Default()
	router.GET("ping", handler.PingHandler)

	// Grouping API v1
	v1 := router.Group("/match/sites/:site_id")
	{
		v1.GET("/items/:item_id/step/init", handler.FormInit)
	}
	return router
}
