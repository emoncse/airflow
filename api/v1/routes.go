package v1

import (
	"github.com/emoncse/airflow/controllers"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes registers the API routes
func RegisterRoutes() *gin.Engine {
	router := gin.Default()
	v1 := router.Group("/api/v1")
	{
		v1.GET("/ping", controllers.Ping)
		v1.GET("/users", controllers.GetUsers)
		v1.GET("/products", controllers.PaymentController)
	}
	return router
}
