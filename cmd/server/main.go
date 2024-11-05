package main

import (
	"log"
	"user-management-mysql/config"
	"user-management-mysql/controllers"
	"user-management-mysql/models"
	"user-management-mysql/repositories"
	"user-management-mysql/routes"
	seeder "user-management-mysql/seeders"
	"user-management-mysql/services"

	"github.com/gin-gonic/gin"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func main() {
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to PostgreSQL
	dsn := "host=" + cfg.PostgresHost + " user=" + cfg.PostgresUser + " password=" + cfg.PostgresPassword + " dbname=" + cfg.PostgresDB + " port=" + cfg.PostgresPort + " sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to database:", err)
	}
	db.AutoMigrate(&models.User{}, &models.Role{})
	if err := seeder.SeedRoles(db); err != nil {
		log.Fatalf("Could not seed roles: %v", err)
	}

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)
	roleRepo := repositories.NewRoleRepository(db)
	productRepo := repositories.NewProductRepository(db)

	// Initialize services
	userService := services.NewUserService(userRepo, roleRepo)
	roleService := services.NewRoleService(roleRepo)
	productService := services.NewProductService(productRepo)

	// Initialize controllers
	userController := controllers.NewUserController(userService)
	roleController := controllers.NewRoleController(roleService)
	productController := controllers.NewProductController(productService)

	// Setup router
	router := gin.Default()

	// Register API routes
	routes.SetupAPIRoutes(router, userController, roleController, productController)

	// Start server
	log.Println("Starting server on port 8080")
	router.Run(":8080")
}
