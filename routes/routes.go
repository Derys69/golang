package routes

import (
	"golang/handlers"

	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	r := gin.Default()

	customer := r.Group("/customers")
	{
		customer.GET("/", handlers.GetCustomers)         // Ambil semua customer
		customer.GET("/:id", handlers.GetCustomerByID)   // Ambil customer berdasarkan ID
		customer.POST("/", handlers.CreateCustomer)      // Tambah customer baru
		customer.PUT("/:id", handlers.UpdateCustomer)    // Update data customer
		customer.DELETE("/:id", handlers.DeleteCustomer) // Hapus customer
	}

	return r
}
