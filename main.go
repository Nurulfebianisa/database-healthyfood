// main.go

package main

import (
	"net/http"
	"tusk/config"
	"tusk/controllers"
	"tusk/models"

	"github.com/gin-gonic/gin"
)

func main() {
	// Database
	db := config.DatabaseConnection()
	db.AutoMigrate(&models.User{}, &models.Task{})
	config.CreateOwnerAccount(db)

	// Controller
	userController := controllers.UserController{DB: db}
	taskController := controllers.TaskController{DB: db}

	// Router
	router := gin.Default()

	// Middleware CORS
	router.Use(func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*") // Atur origin sesuai dengan aplikasi Flutter Anda
		c.Writer.Header().Set("Access-Control-Allow-Methods", "GET, POST, DELETE, PATCH, OPTIONS")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK)
			return
		}
		c.Next()
	})

	// Endpoint
	router.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, "Welcome to Daily Healthy Food Management App")
	})

	router.POST("/users/login", userController.Login)
	router.POST("/users", userController.CreateAccount)
	router.DELETE("/users/:id", userController.Delete)
	router.GET("/users/Employee", userController.GetEmployee)

	router.POST("/tasks", taskController.Create)
	router.DELETE("/tasks/:id", taskController.Delete)
	router.PATCH("/tasks/:id/submit", taskController.Submit)
	router.PATCH("/tasks/:id/reject", taskController.Reject)
	router.PATCH("/tasks/:id/fix", taskController.Fix)
	router.PATCH("/tasks/:id/approve", taskController.Approve)
	router.GET("/tasks/:id", taskController.FindById)
	router.GET("/tasks/review/asc", taskController.NeedToBeReview)
	router.GET("/tasks/progress/:userId", taskController.ProgressTasks)
	router.GET("/tasks/stat/:userId", taskController.Statistic)
	router.GET("/tasks/user/:userId/:status", taskController.FindByUserAndStatus)

	router.Static("/attachments", "./attachments")
	router.Run("192.168.209.160:8080")
}
