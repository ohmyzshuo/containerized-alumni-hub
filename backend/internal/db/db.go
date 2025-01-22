package db

import (
	"github.com/joho/godotenv"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func Init() {
//	dsn := "host=localhost user=shawn password=UMAlumniHub2024! dbname=alumni_hub port=5432 sslmode=disable"
	var err error
	err = godotenv.Load()
    if err != nil {
        log.Fatalf("Error loading .env file")
    }

    dsn := os.Getenv("DB_DSN")
    if dsn == "" {
        log.Fatalf("DB_DSN is not set in .env file")
    }

	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer
		logger.Config{
			SlowThreshold:             time.Second, // 慢 SQL 阈值
			LogLevel:                  logger.Info, // 日志级别
			IgnoreRecordNotFoundError: true,        // 忽略 ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,        // 是否启用彩色打印
		},
	)
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: newLogger,
	})
	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
}
