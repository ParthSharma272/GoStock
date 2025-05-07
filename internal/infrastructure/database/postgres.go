package database

import (
	"fmt"
	productDomain "github.com/ParthSharma272/GoStock/internal/domain/product" // Replace your_project_path
	userDomain "github.com/ParthSharma272/GoStock/internal/domain/user"       // Replace your_project_path
	"github.com/ParthSharma272/GoStock/internal/infrastructure/config"        // Replace your_project_path
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewDatabaseConnection(cfg *config.Config) *gorm.DB {
	var err error
	var db *gorm.DB

	logLevel := logger.Silent
	// if gin.Mode() == gin.DebugMode { // If using Gin directly for mode
	// 	logLevel = logger.Info
	// }

	switch cfg.DBDriver {
	case "postgres":
		db, err = gorm.Open(postgres.Open(cfg.DBUrl), &gorm.Config{Logger: logger.Default.LogMode(logLevel)})
	case "mysql":
		db, err = gorm.Open(mysql.Open(cfg.DBUrl), &gorm.Config{Logger: logger.Default.LogMode(logLevel)})
	default:
		log.Fatalf("Unsupported database driver: %s", cfg.DBDriver)
	}

	if err != nil {
		log.Fatalf("Failed to connect to database: %v", err)
	}
	fmt.Println("Database connection successful!")

	// Auto-migrate schema (for development, consider a proper migration tool for production)
	// Add other models here as they are created (e.g., Category, Order, OrderItem)
	err = db.AutoMigrate(
		&userDomain.User{},
		&productDomain.Product{},
		// &categoryDomain.Category{},
		// &orderDomain.Order{},
		// &orderDomain.OrderItem{},
	)
	if err != nil {
		log.Fatalf("Failed to migrate database: %v", err)
	}
	fmt.Println("Database migration successful!")
	return db
}
