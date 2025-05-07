package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	// Domain (only for type reference if needed, generally avoid direct use in main)

	// Infrastructure
	"github.com/ParthSharma272/GoStock/internal/infrastructure/config"                                // Replace your_project_path
	"github.com/ParthSharma272/GoStock/internal/infrastructure/database"                              // Replace your_project_path
	infraPersistence "github.com/ParthSharma272/GoStock/internal/infrastructure/persistence/postgres" // Replace your_project_path
	"github.com/ParthSharma272/GoStock/internal/infrastructure/web"                                   // Replace your_project_path
	"github.com/ParthSharma272/GoStock/internal/infrastructure/web/handler"                           // Replace your_project_path

	// Service
	"github.com/ParthSharma272/GoStock/internal/service" // Replace your_project_path
)

// @title Go E-commerce Backend API
// @version 1.0
// @description This is a sample server for a basic e-commerce backend.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1

// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
// @description Type "Bearer" followed by a space and JWT token.
func main() {
	// 1. Load Configuration
	// You can pass a path to LoadConfig if your .env is not in the root
	// e.g., config.LoadConfig("../../.env") if main.go is in cmd/api/
	config.LoadConfig() // Assumes .env is in the project root, or use path from executable
	cfg := config.AppConfig

	// 2. Initialize Database Connection
	db := database.NewDatabaseConnection(cfg)

	// 3. Initialize Repositories
	userRepo := infraPersistence.NewUserRepository()
	productRepo := infraPersistence.NewProductRepository()
	// categoryRepo := infraPersistence.NewCategoryRepository()
	// orderRepo := infraPersistence.NewOrderRepository()

	// 4. Initialize Services
	authSvc := service.NewAuthService(db, userRepo, cfg)
	productSvc := service.NewProductService(db, productRepo)
	// categorySvc := service.NewCategoryService(db, categoryRepo)
	// orderSvc := service.NewOrderService(db, orderRepo, productRepo) // Order service might need product repo for stock

	// 5. Initialize Handlers
	authHandler := handler.NewAuthHandler(authSvc)
	productHandler := handler.NewProductHandler(productSvc)
	// categoryHandler := handler.NewCategoryHandler(categorySvc)
	// orderHandler := handler.NewOrderHandler(orderSvc)

	// 6. Setup Router
	router := web.SetupRouter(cfg, authHandler, productHandler /*, categoryHandler, orderHandler */)

	// 7. Start Server with Graceful Shutdown
	srv := &http.Server{
		Addr:    ":" + cfg.Port,
		Handler: router,
	}

	go func() {
		log.Printf("Server listening on port %s\n", cfg.Port)
		log.Printf("API documentation available at http://localhost:%s/swagger/index.html\n", cfg.Port) // If you set up Swagger docs
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown:", err)
	}

	log.Println("Server exiting")
}
