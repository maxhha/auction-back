package database

import (
	"auction-back/models"
	"auction-back/ports"
	"database/sql"
	"fmt"
	"time"

	"gorm.io/gorm"
)

type productDB struct{ *Database }

func (d *Database) Product() ports.ProductDB { return &productDB{d} }

type Product struct {
	ID          string              `gorm:"default:generated();"`
	State       models.ProductState `gorm:"default:'CREATED';"`
	Title       string
	Description string
	CreatorID   string
	CreatedAt   time.Time
	UpdatedAt   time.Time
	DeletedAt   sql.NullTime
}

func (p *Product) into() models.Product {
	return models.Product{
		ID:          p.ID,
		State:       p.State,
		Title:       p.Title,
		Description: p.Description,
		CreatorID:   p.CreatorID,
		CreatedAt:   p.CreatedAt,
		UpdatedAt:   p.UpdatedAt,
		DeletedAt:   p.DeletedAt,
	}
}

func (p *Product) copy(product *models.Product) {
	if product == nil {
		return
	}
	p.ID = product.ID
	p.State = product.State
	p.Title = product.Title
	p.Description = product.Description
	p.CreatorID = product.CreatorID
	p.CreatedAt = product.CreatedAt
	p.UpdatedAt = product.UpdatedAt
	p.DeletedAt = product.DeletedAt
}

func (d *productDB) Create(product *models.Product) error {
	if product == nil {
		return fmt.Errorf("product is nil")
	}
	p := Product{}
	p.copy(product)
	if err := d.db.Create(&p).Error; err != nil {
		return fmt.Errorf("create: %w", err)
	}

	*product = p.into()
	return nil
}

func (d *productDB) filter(query *gorm.DB, config models.ProductsFilter) *gorm.DB {
	if len(config.OwnerIDs) > 0 {
		query = query.Joins(
			"JOIN ( ? ) ofd ON products.id = ofd.product_id AND ofd.owner_n = 1 AND ofd.owner_id IN (?) ",
			d.ownersNumberedQuery(query),
			config.OwnerIDs,
		)
	}
	return query
}

func (d *productDB) Pagination(config ports.ProductPaginationConfig) (models.ProductsConnection, error) {
	query := d.filter(d.db.Model(&Product{}), config.Filter)
	query, err := paginationQueryByCreatedAtDesc(query, config.First, config.After)

	if err != nil {
		return models.ProductsConnection{}, fmt.Errorf("pagination: %w", err)
	}

	var products []models.Product
	result := query.Find(&products)

	if result.Error != nil {
		return models.ProductsConnection{}, result.Error
	}

	if len(products) == 0 {
		return models.ProductsConnection{
			PageInfo: &models.PageInfo{},
			Edges:    make([]*models.ProductsConnectionEdge, 0),
		}, nil
	}

	hasNextPage := false

	if config.First != nil {
		hasNextPage = len(products) > *config.First
		products = products[:len(products)-1]
	}

	edges := make([]*models.ProductsConnectionEdge, 0, len(products))

	for _, node := range products {
		edges = append(edges, &models.ProductsConnectionEdge{
			Cursor: node.ID,
			Node:   &node,
		})
	}

	return models.ProductsConnection{
		PageInfo: &models.PageInfo{
			HasNextPage: hasNextPage,
			StartCursor: &products[0].ID,
			EndCursor:   &products[len(products)-1].ID,
		},
		Edges: edges,
	}, nil
}

func (d *productDB) GetCreator(p models.Product) (models.User, error) {
	return d.User().Get(p.CreatorID)
}

// Query all owners of queried products.
// Returns select query with fields: product_id, owner_id, from_date
func (d *productDB) ownersQuery(query *gorm.DB) *gorm.DB {
	creators := query.Session(&gorm.Session{Initialized: true}).
		Model(&Product{}).
		Select("id as product_id, creator_id as owner_id, created_at as from_date")

	buyers := query.Session(&gorm.Session{Initialized: true}).
		Model(&Product{}).
		Select("products.id as product_id, auctions.buyer_id as owner_id, auctions.finished_at as from_date").
		Joins("FULL JOIN auctions ON auctions.product_id = products.id")

	return d.db.Raw("? UNION ALL ?", creators, buyers)
}

// Query current owner of products
// Returns select query with fields: product_id, owner_id, from_date, owner_n
func (d *productDB) ownersNumberedQuery(query *gorm.DB) *gorm.DB {
	query = d.ownersQuery(query)

	ownersNumbered := d.db.Table("(?) as ofd", query)

	ownersNumbered = ownersNumbered.Select(
		"*, ROW_NUMBER() OVER(PARTITION BY ofd.product_id ORDER BY ofd.from_date DESC) as owner_n",
	)

	return ownersNumbered
}

func (d *productDB) GetOwner(p models.Product) (models.User, error) {
	query := d.db.Model(&Product{}).Where("id = ?", p.ID)
	query = d.ownersNumberedQuery(query)

	owner := User{}
	if err := query.Joins("JOIN users ON users.id = owner_id AND owner_n = 1").Take(&owner).Error; err != nil {
		return models.User{}, fmt.Errorf("take owner: %w", convertError(err))
	}

	return owner.into(), nil
}
