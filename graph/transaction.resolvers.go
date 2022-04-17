package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"auction-back/graph/generated"
	"auction-back/models"
	"context"
	"fmt"
)

func (r *queryResolver) Transactions(ctx context.Context, first *int, after *string, filter *models.TransactionsFilter) (*models.TransactionsConnection, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *transactionResolver) AccountFrom(ctx context.Context, obj *models.Transaction) (*models.Account, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *transactionResolver) AccountTo(ctx context.Context, obj *models.Transaction) (*models.Account, error) {
	panic(fmt.Errorf("not implemented"))
}

func (r *transactionResolver) Offer(ctx context.Context, obj *models.Transaction) (*models.Offer, error) {
	panic(fmt.Errorf("not implemented"))
}

// Transaction returns generated.TransactionResolver implementation.
func (r *Resolver) Transaction() generated.TransactionResolver { return &transactionResolver{r} }

type transactionResolver struct{ *Resolver }
