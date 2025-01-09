package main

import (
	"main/controllers"
	"main/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Middleware
	r.Use(controllers.BlockDesktopBrowsers)
	// Serve static files from the img directory
	r.Static("/img", "./img")

	// Serve the index.html file
	r.LoadHTMLFiles("./public/index.html", "./public/taxa.html", "./public/importante.html", "./public/inicial.html", "./public/dados.html", "./public/pagamento-taxa.html", "./public/inicial.html")

	routes.SetupRoutes(r)

	// Start the server
	r.Run(":8080")
}