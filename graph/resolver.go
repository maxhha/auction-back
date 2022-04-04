package graph

import (
	"auction-back/db"
	"auction-back/ports/bank"
	"auction-back/ports/role"
	"auction-back/ports/token"
	"sync"

	"gorm.io/gorm"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	DB         *gorm.DB
	Market     map[string]chan *db.Product
	MarketLock sync.Mutex
	TokenPort  token.Interface
	BankPort   bank.Interface
	RolePort   role.Interface
}

func New(DB *gorm.DB, token token.Interface, bank bank.Interface, role role.Interface) *Resolver {
	r := Resolver{
		DB:         DB,
		Market:     make(map[string]chan *db.Product),
		MarketLock: sync.Mutex{},
		TokenPort:  token,
		BankPort:   bank,
		RolePort:   role,
	}

	return &r
}
