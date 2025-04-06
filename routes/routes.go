package routes

import (
	"golang/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	customer := r.Group("/customers")
	{
		customer.GET("/", handlers.GetCustomers)
		customer.GET("/:id", handlers.GetCustomerByID)
		customer.POST("/", handlers.CreateCustomer)
		customer.PUT("/:id", handlers.UpdateCustomer)
		customer.DELETE("/:id", handlers.DeleteCustomer)
	}

	return r
}
