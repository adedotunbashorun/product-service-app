package main

import (
	"fmt"
	"log"
	"product-service-app/config"
	"product-service-app/controllers"
	"product-service-app/docs"
	"product-service-app/models"
	"product-service-app/repositories"
	"product-service-app/routes"
	"product-service-app/seeder"
	"product-service-app/services"
	"time"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func connectToDB() *gorm.DB {
	// var err error
	// Load configuration
	cfg := config.LoadConfig()

	// Connect to PostgreSQL
	dsn := "host=" + cfg.PostgresHost + " user=" + cfg.PostgresUser + " password=" + cfg.PostgresPassword + " dbname=" + cfg.PostgresDB + " port=" + cfg.PostgresPort + " sslmode=disable"

	for retries := 5; retries > 0; retries-- {
		db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
		if err == nil {
			fmt.Println("Database connected successfully!")
			return db
		}
		fmt.Println("Failed to connect to database. Retrying in 5 seconds...")
		time.Sleep(5 * time.Second)
	}
	fmt.Println("Could not connect to the database. Exiting.")
	return nil
}

// @title Product Management API
// @version 1.0
// @description API documentation for Product Management Service.
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	db := connectToDB()
	if db == nil {
		fmt.Println("Exiting application due to database connection failure.")
		return
	}
	db.AutoMigrate(&models.User{}, &models.Role{}, &models.Product{}, &models.Order{}, &models.OrderItem{})
	if err := seeder.SeedRoles(db); err != nil {
		log.Fatalf("Could not seed roles: %v", err)
	}
	if err := seeder.SeedUsers(db); err != nil {
		log.Fatalf("Could not seed users: %v", err)
	}
	if err := seeder.SeedProducts(db); err != nil {
		log.Fatalf("Could not seed products: %v", err)
	}
	if err := seeder.SeedOrders(db); err != nil {
		log.Fatalf("Could not seed orders: %v", err)
	}

	// Initialize repositories
	userRepo := repositories.NewUserRepository(db)
	roleRepo := repositories.NewRoleRepository(db)
	productRepo := repositories.NewProductRepository(db)
	orderRepo := repositories.NewOrderRepository(db)

	// Initialize services
	userService := services.NewUserService(userRepo, roleRepo)
	roleService := services.NewRoleService(roleRepo)
	productService := services.NewProductService(productRepo)
	orderService := services.NewOrderService(orderRepo)

	// Initialize controllers
	userController := controllers.NewUserController(userService)
	roleController := controllers.NewRoleController(roleService)
	productController := controllers.NewProductController(productService)
	orderController := controllers.NewOrderController(orderService)

	// Setup router
	router := gin.Default()

	docs.SwaggerInfo.BasePath = "/api"
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Register API routes
	routes.SetupAPIRoutes(router, userController, roleController, productController, orderController)

	// Start server
	log.Println("Starting server on port 8080")
	router.Run(":8080")
}
