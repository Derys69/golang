package handlers

import (
	"database/sql"
	"golang/database"
	"golang/units"
	"net/http"

	"github.com/gin-gonic/gin"
)

// Fungsi untuk mendapatkan data semua Customer
func GetCustomers(c *gin.Context) {
	rows, err := database.DB.Query("SELECT * FROM customer")
	if err != nil {
		// Jika query gagal, maka error 500
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	defer rows.Close()

	var customers []units.Customer
	for rows.Next() {
		var cust units.Customer
		if err := rows.Scan(&cust.KodeCust, &cust.Nama, &cust.Alamat, &cust.Kota, &cust.Telepon); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		customers = append(customers, cust)
	}
	c.JSON(http.StatusOK, customers)
}

// Fungsi untuk Mengambil data Customer berdasarkan userid yang dicari
func GetCustomerByID(c *gin.Context) {
	id := c.Param("id")
	var cust units.Customer
	err := database.DB.QueryRow("SELECT * FROM customer WHERE kodecust = ?", id).
		Scan(&cust.KodeCust, &cust.Nama, &cust.Alamat, &cust.Kota, &cust.Telepon)
	if err != nil {
		if err == sql.ErrNoRows {
			// Jika data tidak ditemukan, error 404
			c.JSON(http.StatusNotFound, gin.H{"error": "Customer tidak ditemukan"})
		} else {
			// Jika terjadi error lain, error 500
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		}
		return
	}
	c.JSON(http.StatusOK, cust)
}

// Fungsi untuk Membuat data baru Customer
func CreateCustomer(c *gin.Context) {
	var cust units.Customer
	// Bind JSON dari body ke struct
	if err := c.ShouldBindJSON(&cust); err != nil {
		// Jika JSON tidak valid, error 400
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Simpan data ke database
	_, err := database.DB.Exec("INSERT INTO customer VALUES (?, ?, ?, ?, ?)",
		cust.KodeCust, cust.Nama, cust.Alamat, cust.Kota, cust.Telepon)

	if err != nil {
		// Jika gagal menyimpan data, error 500
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// Berhasil menambahkan data
	c.JSON(http.StatusCreated, gin.H{"message": "Customer ditambahkan"})
}

// Fungsi untuk mengubah data Customer
func UpdateCustomer(c *gin.Context) {
	id := c.Param("id")
	var cust units.Customer
	// Bind data dari body JSON
	if err := c.ShouldBindJSON(&cust); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// Update data customer di database
	_, err := database.DB.Exec("UPDATE customer SET nama=?, alamat=?, kota=?, telepon=? WHERE kodecust=?",
		cust.Nama, cust.Alamat, cust.Kota, cust.Telepon, id)
	if err != nil {
		// Jika gagal, error 500
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// response jika berhasil
	c.JSON(http.StatusOK, gin.H{"message": "Customer diupdate"})
}

// Fungsi untuk menghapus data Customer
func DeleteCustomer(c *gin.Context) {
	id := c.Param("id")
	// Hapus data dari database
	_, err := database.DB.Exec("DELETE FROM customer WHERE kodecust = ?", id)

	if err != nil {
		// Jika ada masalah, error 500
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	// response jika berhasil
	c.JSON(http.StatusOK, gin.H{"message": "Customer dihapus"})
}
