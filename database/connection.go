package database

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var DB *sql.DB

func InitDB() {
	var err error
	// Format: "username:password@tcp(host:port)/nama_database"
	// Ganti root dan password sesuai dengan database yang digunakan"
	dsn := "root:@Derysurya12345@tcp(127.0.0.1:3306)/db_uco"
	DB, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal("Gagal terhubung dengan database:", err)
	}
	// Mengecek apakah koneksi ke database berhasil
	if err = DB.Ping(); err != nil {
		log.Fatal("Database tidak merespons:", err)
	}
	log.Println("Koneksi berhasil")
}
