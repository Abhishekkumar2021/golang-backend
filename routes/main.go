package routes

import (
	"github.com/Abhishekkumar2021/golang-backend/controllers"
	"github.com/gin-gonic/gin"
)

// Healthcheck routes
func HealthCheckRouter(server *gin.Engine) {
	server.GET("/healthcheck", controllers.HealthCheck)
}

// User routes
func UserRouter(server *gin.Engine) {
	server.POST("/user", controllers.AddUser)
	server.GET("/user/:id", controllers.GetUserByID)
	server.GET("/user", controllers.GetAllUsers)
	server.GET("/user/email/:email", controllers.GetUserByEmail)
	server.GET("/user/username/:username", controllers.GetUserByUsername)
	server.PATCH("/user/:id", controllers.UpdateUser)
	server.DELETE("/user/:id", controllers.DeleteUser)
}

func StreamRouter(server *gin.Engine) {
	server.GET("/stream", controllers.Stream)
}