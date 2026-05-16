package controllers

import (
	"CRUD-golang/internal/dto"
	"CRUD-golang/internal/usecases"
	apperror "CRUD-golang/pkg/utils/app_error"
	"errors"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UsuarioController struct {
	uc *usecases.UsuarioUsecases
}

func NewUsuarioController(uc *usecases.UsuarioUsecases) UsuarioController {
	return UsuarioController{uc: uc}
}

func (c *UsuarioController) Create(ctx *gin.Context) {
	var req dto.CreateUsuarioRequest
	if err := ctx.ShouldBindJSON(&req); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"erro":     true,
			"mensagem": "body inválido",
		})
		return
	}

	usuario, err := c.uc.Create(req)
	if err != nil {
		var appErr *apperror.AppError
		if errors.As(err, &appErr) {
			appErr = err.(*apperror.AppError)
			ctx.JSON(appErr.StatusCode, gin.H{
				"erro":     true,
				"mensagem": err.Error(),
			})
			return
		}
		ctx.JSON(http.StatusInternalServerError, gin.H{
			"erro":     true,
			"mensagem": err.Error(),
		})
		return
	}

	ctx.JSON(http.StatusCreated, gin.H{
		"erro":    false,
		"usuario": usuario,
	})
}
