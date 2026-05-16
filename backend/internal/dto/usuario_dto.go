package dto

import (
	"fmt"
	"strings"
)

type CreateUsuarioRequest struct {
	Nome  string `json:"nome"`
	Email string `json:"email"`
	Login string `json:"login"`
	Senha string `json:"senha"`
}

func (u *CreateUsuarioRequest) ValidarCampos() error {
	if u.Nome == "" || u.Email == "" || u.Login == "" || u.Senha == "" {
		return fmt.Errorf("todos os campos são obrigatórios")
	}
	if !strings.Contains(u.Email, "@") || !strings.Contains(u.Email, ".") {
		return fmt.Errorf("o email deve ser válido")
	}
	return nil
}
