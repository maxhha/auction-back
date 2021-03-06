// This file was generated by gormdbopsgen. DO NOT EDIT.
package database

import (
	"auction-back/models"
	"auction-back/ports"
	"fmt"
)

type transactionDB struct{ *Database }

func (d *Database) Transaction() ports.TransactionDB { return &transactionDB{d} }

func (d *transactionDB) Create(transaction *models.Transaction) error {
	if transaction == nil {
		return ports.ErrTransactionIsNil
	}
	obj := Transaction{}
	obj.copy(transaction)
	if err := d.db.Create(&obj).Error; err != nil {
		return fmt.Errorf("create: %w", convertError(err))
	}

	*transaction = obj.into()
	return nil
}

func (d *transactionDB) Update(transaction *models.Transaction) error {
	if transaction == nil {
		return ports.ErrTransactionIsNil
	}

	obj := Transaction{}
	obj.copy(transaction)

	if err := d.db.Save(&obj).Error; err != nil {
		return fmt.Errorf("save: %w", convertError(err))
	}
	*transaction = obj.into()

	return nil
}

func (d *transactionDB) Pagination(first *int, after *string, filter *models.TransactionsFilter) (models.TransactionsConnection, error) {
	query := d.filter(d.db.Model(&Transaction{}), filter)
	query, err := paginationQueryByCreatedAtDesc(query, first, after)
	if err != nil {
		return models.TransactionsConnection{}, fmt.Errorf("pagination: %w", err)
	}

	var objs []Transaction
	if err := query.Find(&objs).Error; err != nil {
		return models.TransactionsConnection{}, fmt.Errorf("find: %w", err)
	}

	if len(objs) == 0 {
		return models.TransactionsConnection{
			PageInfo: &models.PageInfo{},
			Edges:    make([]*models.TransactionsConnectionEdge, 0),
		}, nil
	}

	hasNextPage := false

	if first != nil {
		hasNextPage = len(objs) > *first
		objs = objs[:len(objs)-1]
	}

	edges := make([]*models.TransactionsConnectionEdge, 0, len(objs))

	for _, obj := range objs {
		node := obj.into()

		edges = append(edges, &models.TransactionsConnectionEdge{
			Cursor: fmt.Sprintf("%v", node.ID),
			Node:   &node,
		})
	}

	startCursor := fmt.Sprintf("%v", objs[0].ID)
	endCursor := fmt.Sprintf("%v", objs[len(objs)-1].ID)

	return models.TransactionsConnection{
		PageInfo: &models.PageInfo{
			HasNextPage: hasNextPage,
			StartCursor: &startCursor,
			EndCursor:   &endCursor,
		},
		Edges: edges,
	}, nil
}
