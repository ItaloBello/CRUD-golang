package routes

import (
	"CRUD-golang/internal/controllers"

	"github.com/gin-gonic/gin"
)

func SetupUsuarioRoutes(a *gin.RouterGroup, ctrl *controllers.UsuarioController){
	usuario := a.Group("/usuario")
	{
		usuario.POST("/", ctrl.Create)
	}
}