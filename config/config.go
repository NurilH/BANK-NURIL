package config

import (
	"test-b/models"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

// inisialisasi database
func InitDB() {

	// config := "root:nuril123@tcp(127.0.0.1:3306)/test_bank?charset=utf8mb4&parseTime=True&loc=Local"
	config := "user:root123@tcp(bankDB)/test_bank?charset=utf8mb4&parseTime=True&loc=Local"

	var e error

	DB, e = gorm.Open(mysql.Open(config), &gorm.Config{})
	if e != nil {
		panic(e)
	}
	InitMigrate()
}

// auto migrate -> untuk membuat tabel otomatis jika tabel tidak terdapat pada database
func InitMigrate() {
	DB.AutoMigrate(&models.Nasabah{})
}
