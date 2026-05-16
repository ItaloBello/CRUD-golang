package models

import (
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Usuario struct {
	Id           int       `db:"id" json:"id"`
	Nome         string    `db:"nome" json:"nome"`
	Email        string    `db:"email" json:"email"`
	Login        string    `db:"login" json:"login"`
	Senha        string    `db:"senha" json:"senha"`
	Role         string    `db:"role" json:"role"`
	CriadoEm     time.Time `db:"criado_em" json:"criadoEm"`
	AtualizadoEm time.Time `db:"atualizado_em" json:"atualizadoEm"`
}

func (u *Usuario) EncriptarSenha() error {
	senhaByte := []byte(u.Senha)
	hash, err := bcrypt.GenerateFromPassword(senhaByte, bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	u.Senha = string(hash)
	return nil
}

