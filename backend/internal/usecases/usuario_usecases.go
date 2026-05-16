package usecases

import (
	"CRUD-golang/internal/dto"
	"CRUD-golang/internal/models"
	"CRUD-golang/internal/repositories"
	apperror "CRUD-golang/pkg/utils/app_error"
	"database/sql"
)

type UsuarioUsecases struct {
	repo *repositories.UsuarioRepo
}

func NewUsuarioUsecases(repo *repositories.UsuarioRepo) *UsuarioUsecases {
	return &UsuarioUsecases{repo: repo}
}

func (u *UsuarioUsecases) Create(req dto.CreateUsuarioRequest) (*models.Usuario, error) {
	err := req.ValidarCampos()
	if err != nil {
		return nil, apperror.BadRequest(err.Error())
	}

	us, err := u.repo.GetByEmail(req.Email)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
	}
	if us != nil {
		return nil, apperror.Conflict("este email já está cadastrado")
	}

	us, err = u.repo.GetByLogin(req.Login)
	if err != nil {
		if err != sql.ErrNoRows {
			return nil, err
		}
	}
	if us != nil {
		return nil, apperror.Conflict("este login já está cadastrado")
	}

	usuario := models.Usuario{
		Nome: req.Nome,
		Email: req.Email,
		Login: req.Login,
		Senha: req.Senha,
		Role: "USER",
	}

	err = usuario.EncriptarSenha()
	if err != nil {
		return nil, err
	}

	err = u.repo.Create(&usuario)
	if err != nil {
		return nil, err
	}

	usuario.Senha = ""
	
	return &usuario, nil
}
