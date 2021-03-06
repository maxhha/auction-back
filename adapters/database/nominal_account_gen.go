// This file was generated by gormdbopsgen. DO NOT EDIT.
package database

import (
	"auction-back/models"
	"auction-back/ports"
	"fmt"
)

type nominalAccountDB struct{ *Database }

func (d *Database) NominalAccount() ports.NominalAccountDB { return &nominalAccountDB{d} }

func (d *nominalAccountDB) Get(id string) (models.NominalAccount, error) {
	obj := NominalAccount{}
	if err := d.db.Take(&obj, "id = ?", id).Error; err != nil {
		return models.NominalAccount{}, fmt.Errorf("take: %w", convertError(err))
	}

	return obj.into(), nil
}

func (d *nominalAccountDB) Take(config ports.NominalAccountTakeConfig) (models.NominalAccount, error) {
	query := d.filter(d.db, config.Filter)

	obj := NominalAccount{}
	if err := query.Take(&obj).Error; err != nil {
		return models.NominalAccount{}, fmt.Errorf("take: %w", convertError(err))
	}

	return obj.into(), nil
}

func (d *nominalAccountDB) Create(nominalAccount *models.NominalAccount) error {
	if nominalAccount == nil {
		return ports.ErrNominalAccountIsNil
	}
	obj := NominalAccount{}
	obj.copy(nominalAccount)
	if err := d.db.Create(&obj).Error; err != nil {
		return fmt.Errorf("create: %w", convertError(err))
	}

	*nominalAccount = obj.into()
	return nil
}

func (d *nominalAccountDB) Update(nominalAccount *models.NominalAccount) error {
	if nominalAccount == nil {
		return ports.ErrNominalAccountIsNil
	}

	obj := NominalAccount{}
	obj.copy(nominalAccount)

	if err := d.db.Save(&obj).Error; err != nil {
		return fmt.Errorf("save: %w", convertError(err))
	}
	*nominalAccount = obj.into()

	return nil
}

func (d *nominalAccountDB) Pagination(first *int, after *string, filter *models.NominalAccountsFilter) (models.NominalAccountsConnection, error) {
	query := d.filter(d.db.Model(&NominalAccount{}), filter)
	query, err := paginationQueryByCreatedAtDesc(query, first, after)
	if err != nil {
		return models.NominalAccountsConnection{}, fmt.Errorf("pagination: %w", err)
	}

	var objs []NominalAccount
	if err := query.Find(&objs).Error; err != nil {
		return models.NominalAccountsConnection{}, fmt.Errorf("find: %w", err)
	}

	if len(objs) == 0 {
		return models.NominalAccountsConnection{
			PageInfo: &models.PageInfo{},
			Edges:    make([]*models.NominalAccountsConnectionEdge, 0),
		}, nil
	}

	hasNextPage := false

	if first != nil {
		hasNextPage = len(objs) > *first
		objs = objs[:len(objs)-1]
	}

	edges := make([]*models.NominalAccountsConnectionEdge, 0, len(objs))

	for _, obj := range objs {
		node := obj.into()

		edges = append(edges, &models.NominalAccountsConnectionEdge{
			Cursor: fmt.Sprintf("%v", node.ID),
			Node:   &node,
		})
	}

	startCursor := fmt.Sprintf("%v", objs[0].ID)
	endCursor := fmt.Sprintf("%v", objs[len(objs)-1].ID)

	return models.NominalAccountsConnection{
		PageInfo: &models.PageInfo{
			HasNextPage: hasNextPage,
			StartCursor: &startCursor,
			EndCursor:   &endCursor,
		},
		Edges: edges,
	}, nil
}
