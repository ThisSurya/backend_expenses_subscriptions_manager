package main

import (
	"backend/config"
	"backend/controllers"
	db "backend/models/config"
	"backend/repository"
	"backend/routes"
	"backend/routes/middleware"
	"backend/services"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	cfg, err := config.Load()
	if err != nil {
		log.Fatal("Failed to Load Config", err)
	}
	gin.SetMode(cfg.Env)
	DB := db.InitPostgresDB(cfg.Database.Host, cfg.Database.Port, cfg.Database.User, cfg.Database.Password, cfg.Database.DbName)

	router := gin.New()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())

	// CORS middleware
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(204)
			return
		}
		c.Next()
	})

	public_api := router.Group("/v1/api")
	protected_api := router.Group("/v1/api")
	protected_api.Use(middleware.AuthMiddleware([]byte(cfg.JWT.SecretKey)))
	protected_api.Use(middleware.RateLimitMiddleware())

	// Repositories
	userRepo := repository.NewUserRepository(DB)
	expenseRepo := repository.NewExpenseRepository(DB)
	categoryRepo := repository.NewCategoryRepository(DB)
	subscriptionRepo := repository.NewSubscriptionController(DB)

	// Services
	userService := services.NewUserService(userRepo, []byte(cfg.JWT.SecretKey), cfg.JWT.TokenExp)
	expenseService := services.NewExpenseService(expenseRepo)
	categoryService := services.NewCategoryService(categoryRepo)
	subscriptionService := services.NewSubscriptionService(subscriptionRepo)

	// Controllers
	authController := controllers.NewAuthController(userService)
	expenseController := controllers.NewExpenseController(expenseService)
	categoryController := controllers.NewCategoryController(categoryService)
	subscriptionController := controllers.NewSubscriptionController(subscriptionService)

	public_api.POST("/register", authController.Register)
	public_api.POST("/login", authController.Login)

	routes.ExpenseRoutes(protected_api, expenseController)
	routes.CategoryRoutes(protected_api, categoryController)
	routes.SubscriptionRoutes(protected_api, subscriptionController)

	router.Run(":" + cfg.Server.Port)
}
