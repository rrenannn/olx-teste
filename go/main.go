package main

import (
	"main/go/controllers"
	"main/go/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Middleware
	r.Use(controllers.BlockDesktopBrowsers)
	// Serve static files from the img directory
	r.Static("/img", "./img")

	// Serve the index.html file
	r.LoadHTMLFiles("./html/index.html", "./html/taxa.html", "./html/importante.html", "./html/inicial.html", "./html/dados.html", "./html/pagamento-taxa.html", "./html/inicial.html")

	routes.SetupRoutes(r)

	// Start the server
	r.Run(":8080")
}