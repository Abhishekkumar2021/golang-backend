package main

import (
	"os"

	"github.com/Abhishekkumar2021/golang-backend/middlewares"
	"github.com/Abhishekkumar2021/golang-backend/routes"
	"github.com/Abhishekkumar2021/golang-backend/utils"
	"github.com/gin-gonic/gin"
)

// This functions runs at the very beginning of the application
func init() {
	utils.LoadEnv()
}

func main() {
	// stop default console logging
	gin.SetMode(gin.ReleaseMode)

	// create a new gin router
	server := gin.Default()

	// apply middlewares
	server.Use(middlewares.Logger)

	// define routes
	routes.HealthCheckRouter(server)
	routes.UserRouter(server)
	routes.StreamRouter(server)

	// Serve static files
	server.Static("/public", "./public")

	// start the server
	PORT := string(os.Getenv("PORT"))
	server.Run(":" + PORT)
}
