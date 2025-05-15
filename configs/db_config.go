package configs

import (
	"fmt"
	"log"
	"nft-service/models"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func InitDB(cfg *Config) *gorm.DB {
	dsn := fmt.Sprintf("host=%s user=username password=password dbname=nft_service port=5432 sslmode=disable TimeZone=Asia/Ho_Chi_Minh", cfg.DbHost)
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatalf("failed to connect to database: %v", err)
	}

	// Tự động migrate bảng `tokens`
	if err := db.AutoMigrate(&models.Token{}); err != nil {
		log.Fatalf("failed to migrate database: %v", err)
	}

	return db
}
