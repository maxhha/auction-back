package database

import (
	"auction-back/models"
	"auction-back/ports"
	"database/sql"
	"time"

	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

//go:generate go run ../../codegen/gormdbops/main.go --out account_gen.go --model Account --methods Get,Take,Create,Update,Pagination

type Account struct {
	ID               string `gorm:"default:generated();"`
	Number           string
	UserID           string
	NominalAccountID string
	CreatedAt        time.Time
	DeletedAt        gorm.DeletedAt
}

func (a *Account) into() models.Account {
	return models.Account{
		ID:               a.ID,
		Number:           a.Number,
		UserID:           a.UserID,
		NominalAccountID: a.NominalAccountID,
		CreatedAt:        a.CreatedAt,
		DeletedAt:        sql.NullTime(a.DeletedAt),
	}
}

func (a *Account) copy(account *models.Account) {
	if account == nil {
		return
	}
	a.ID = account.ID
	a.Number = account.Number
	a.UserID = account.UserID
	a.NominalAccountID = account.NominalAccountID
	a.CreatedAt = account.CreatedAt
	a.DeletedAt = gorm.DeletedAt(account.DeletedAt)
}

func (d *accountDB) filter(query *gorm.DB, config *models.AccountsFilter) *gorm.DB {
	if config != nil {
		return query
	}

	if len(config.UserIDs) > 0 {
		query = query.Where("user_id IN ?", config.UserIDs)
	}

	if config.AvailableFrom != nil {
		panic("unimplimented!")
	}

	return query
}

func (d *accountDB) LockFull(account *models.Account) error {
	if account == nil {
		return ports.ErrAccountIsNil
	}
	obj := Account{}
	err := d.db.Clauses(clause.Locking{
		Strength: "UPDATE",
		Table:    clause.Table{Name: clause.CurrentTable},
	}).
		Take(&obj, "id = ?", account.ID).
		Error
	if err != nil {
		return convertError(err)
	}

	*account = obj.into()
	return nil
}

// var countingTransactionStates = []models.TransactionState{
// 	models.TransactionStateSucceeded,
// 	models.TransactionStateProcessing,
// 	models.TransactionStateError,
// }

// func (d *accountDB) availableMoneyQuery(query *gorm.DB) *gorm.DB {
// 	toTrs := d.db.Model(&Transaction{}).
// 		Select("currency, account_to_id as account_id, amount").
// 		Joins(
// 			"JOIN ( ? ) a ON account_to_id = a.id AND transaction.state IN ?",
// 			query.Session(&gorm.Session{Initialized: true}).Model(&Account{}),
// 			countingTransactionStates,
// 		)

// 	fromTrs := d.db.Model(&Transaction{}).
// 		Select("currency, account_from_id as account_id, -amount").
// 		Joins(
// 			"JOIN ( ? ) a ON account_from_id = a.id AND ( state IN ? OR ( type IN ? AND state IN ? ) )",
// 			query.Session(&gorm.Session{Initialized: true}).Model(&Account{}),
// 			countingTransactionStates,
// 			[]models.TransactionType{
// 				models.TransactionTypeBuy,
// 				// TODO: Add models.TransactionTypeBuyFee
// 			},
// 			[]models.TransactionState{
// 				models.TransactionStateCreated,
// 			},
// 		)

// 	allTrs := d.db.Raw("? UNION ALL ?", fromTrs, toTrs)

// 	moneyQuery := d.db.
// 		Select("a.currency, a.account_id, SUM(a.amount) as amount").
// 		Table("? a", allTrs).
// 		Group("a.currency, a.account_id")

// 	return moneyQuery
// }

func (d *accountDB) GetAvailableMoney(account models.Account) (map[models.CurrencyEnum]models.Money, error) {
	// TODO: Test this
	// query := d.availableMoneyQuery(d.db.Model(&Account{}).Where("id = ?", account.ID))

	// var moneys []models.Money
	// if err := query.Scan(&moneys).Error; err != nil {
	// 	return nil, fmt.Errorf("scan: %w", err)
	// }

	// moneysMap := make(map[models.CurrencyEnum]models.Money, len(moneys))
	// for _, m := range moneys {
	// 	moneysMap[m.Currency] = m
	// }

	// return moneysMap, nil
	moneysMap := make(map[models.CurrencyEnum]models.Money, 1)
	moneysMap[models.CurrencyEnumRub] = models.Money{
		Currency: models.CurrencyEnumRub,
		Amount:   decimal.NewFromFloatWithExponent(100, -2),
	}

	return moneysMap, nil
}
