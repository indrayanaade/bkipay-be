package database

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}

func ConnectDB() (*gorm.DB, error) {
	// Load .env file
	_ = godotenv.Load(".env") // abaikan error kalau tidak ada

	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Jakarta",
		getEnv("DB_HOST", "bkipay-db"),
		getEnv("DB_USER", "bkipay"),
		getEnv("DB_PASS", "adMin!1234"),
		getEnv("DB_NAME", "bkipaydb"),
		getEnv("DB_PORT", "5432"),
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}