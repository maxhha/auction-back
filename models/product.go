package models

import (
	"database/sql"
	"time"
)

type Product struct {
	ID            string       `json:"id"`
	State         ProductState `json:"state"`
	Title         string       `json:"title"`
	Description   string       `json:"description"`
	CreatorID     string
	DeclainReason *string `json:"declainReason"`
	CreatedAt     time.Time
	UpdatedAt     time.Time
	DeletedAt     sql.NullTime
}

func (p *Product) IsEditable() bool {
	return p.State == ProductStateCreated || p.State == ProductStateDeclained
}
