package routes

import (
	"main/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRoutes(r *gin.Engine) {
	r.GET("/", controllers.Index)
	r.POST("/login", controllers.SalvarUsuario)
	r.GET("/taxa", controllers.Tax)
	r.GET("/importante", controllers.Important)
	r.GET("/dados", controllers.Info)
	r.GET("/pagamento-taxa", controllers.Payment)
	r.GET("/inicial", controllers.Inicial)
	r.GET("/admin", controllers.AdminIndex)
	r.GET("/usuarios/:id", controllers.UsuarioShow)
}