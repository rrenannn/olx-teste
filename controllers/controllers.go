package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"strings"
	"fmt"
)

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

func BlockDesktopBrowsers(c *gin.Context) {
	fmt.Println("Middleware chamado")
    userAgent := c.Request.UserAgent()
	if userAgent != "" {
    	fmt.Println(userAgent)
	} else {
    fmt.Println("User-Agent não está presente ou está vazio")
	}
    if strings.Contains(userAgent, "Windows NT") || strings.Contains(userAgent, "Macintosh") || strings.Contains(userAgent, "Linux") || strings.Contains(userAgent, "X11") {
        c.AbortWithStatus(403)
        return
    }
    c.Next()
}