package config

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectDB() {
	host := App.DB_HOST
	user := App.DB_USER
	pass := App.DB_PASS
	dbname := App.DB_NAME
	port := App.DB_PORT
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Ho_Chi_Minh", host, user, pass, dbname, port)
	log.Println(dsn)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("❌ Lỗi kết nối DB: %v", err)
	}
	DB = db
	log.Println("✅ Đã kết nối Database thành công")
}