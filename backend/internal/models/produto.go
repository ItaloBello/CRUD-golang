package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Produto struct {
	Id            int             `db:"id" json:"id"`
	Nome          string          `db:"nome" json:"nome"`
	Descricao     string          `db:"descricao" json:"descricao"`
	ValorUnitario decimal.Decimal `db:"valor_unitario" json:"valorUnitario"`
	UnidadeMedida string          `db:"unidade_medida" json:"unidadeMedida"`
	Quantidade    int             `db:"quantidade" json:"quantidade"`
	UsuarioId     int             `db:"usuario_id" json:"usuarioId"`
	CriadoEm      time.Time       `db:"criado_em" json:"criadoEm"`
	AtualizadoEm  time.Time       `db:"atualizado_em" json:"atualizadoEm"`
}
