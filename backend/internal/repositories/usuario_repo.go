package repositories

import (
	"CRUD-golang/internal/models"

	"github.com/jmoiron/sqlx"
)

type UsuarioRepo struct {
	db *sqlx.DB
}

func NewUsuarioRepo(db *sqlx.DB) *UsuarioRepo {
	return &UsuarioRepo{db: db}
}

func (r *UsuarioRepo) Create(usuario *models.Usuario) error {
	query := `
		INSERT INTO usuario (
			nome, 
			email,
			login,
			senha,
			role
		) 
		VALUES ($1, $2, $3, $4, $5) 
		RETURNING id, criado_em, atualizado_em;
	`
	return r.db.Get(&usuario, query, usuario.Nome, usuario.Email, usuario.Login, usuario.Senha, usuario.Role)
}

func (r *UsuarioRepo) GetByEmail(email string) (*models.Usuario, error) {
	query := `SELECT * FROM usuario WHERE email = $1`
	var usuario models.Usuario
	err := r.db.Get(&usuario, query, email)
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}

func (r *UsuarioRepo) GetByLogin(login string) (*models.Usuario, error) {
	query := `SELECT * FROM usuario WHERE login = $1`
	var usuario models.Usuario
	err := r.db.Get(&usuario, query, login)
	if err != nil {
		return nil, err
	}
	return &usuario, nil
}
