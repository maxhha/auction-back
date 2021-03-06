// This file was generated by portmocksgen. DO NOT EDIT.
package ports

import (
	"auction-back/models"
	"github.com/stretchr/testify/mock"
)

type AccountDBMock struct{ mock.Mock }

type AuctionDBMock struct{ mock.Mock }

type BankDBMock struct{ mock.Mock }

type DBMock struct {
	AccountMock        AccountDBMock
	AuctionMock        AuctionDBMock
	BankMock           BankDBMock
	NominalAccountMock NominalAccountDBMock
	OfferMock          OfferDBMock
	ProductMock        ProductDBMock
	RoleMock           RoleDBMock
	TokenMock          TokenDBMock
	TransactionMock    TransactionDBMock
	TxMock             TXDBMock
	UserMock           UserDBMock
	UserFormMock       UserFormDBMock
}

type NominalAccountDBMock struct{ mock.Mock }

type OfferDBMock struct{ mock.Mock }

type ProductDBMock struct{ mock.Mock }

type RoleDBMock struct{ mock.Mock }

type TXDBMock struct{ mock.Mock }

type TokenDBMock struct{ mock.Mock }

type TransactionDBMock struct{ mock.Mock }

type UserDBMock struct{ mock.Mock }

type UserFormDBMock struct{ mock.Mock }

func (m *AccountDBMock) Create(account *models.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *AccountDBMock) Get(id string) (models.Account, error) {
	args := m.Called(id)
	return args.Get(0).(models.Account), args.Error(1)
}

func (m *AccountDBMock) GetAvailableMoney(account models.Account) (map[models.CurrencyEnum]models.Money, error) {
	args := m.Called(account)
	return args.Get(0).(map[models.CurrencyEnum]models.Money), args.Error(1)
}

func (m *AccountDBMock) LockFull(account *models.Account) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *AccountDBMock) Pagination(first *int, after *string, filter *models.AccountsFilter) (models.AccountsConnection, error) {
	args := m.Called(first, after, filter)
	return args.Get(0).(models.AccountsConnection), args.Error(1)
}

func (m *AccountDBMock) Take(config AccountTakeConfig) (models.Account, error) {
	args := m.Called(config)
	return args.Get(0).(models.Account), args.Error(1)
}

func (m *AuctionDBMock) Create(auction *models.Auction) error {
	args := m.Called(auction)
	return args.Error(0)
}

func (m *AuctionDBMock) Get(id string) (models.Auction, error) {
	args := m.Called(id)
	return args.Get(0).(models.Auction), args.Error(1)
}

func (m *AuctionDBMock) LockShare(auction *models.Auction) error {
	args := m.Called(auction)
	return args.Error(0)
}

func (m *AuctionDBMock) Pagination(first *int, after *string, filter *models.AuctionsFilter) (models.AuctionsConnection, error) {
	args := m.Called(first, after, filter)
	return args.Get(0).(models.AuctionsConnection), args.Error(1)
}

func (m *AuctionDBMock) Take(config AuctionTakeConfig) (models.Auction, error) {
	args := m.Called(config)
	return args.Get(0).(models.Auction), args.Error(1)
}

func (m *AuctionDBMock) Update(auction *models.Auction) error {
	args := m.Called(auction)
	return args.Error(0)
}

func (m *BankDBMock) Create(bank *models.Bank) error {
	args := m.Called(bank)
	return args.Error(0)
}

func (m *BankDBMock) Get(id string) (models.Bank, error) {
	args := m.Called(id)
	return args.Get(0).(models.Bank), args.Error(1)
}

func (m *BankDBMock) Pagination(first *int, after *string, filter *models.BanksFilter) (models.BanksConnection, error) {
	args := m.Called(first, after, filter)
	return args.Get(0).(models.BanksConnection), args.Error(1)
}

func (m *BankDBMock) Take(config BankTakeConfig) (models.Bank, error) {
	args := m.Called(config)
	return args.Get(0).(models.Bank), args.Error(1)
}

func (m *BankDBMock) Update(bank *models.Bank) error {
	args := m.Called(bank)
	return args.Error(0)
}

func (m *DBMock) Account() AccountDB {
	return &m.AccountMock
}

func (m *DBMock) Auction() AuctionDB {
	return &m.AuctionMock
}

func (m *DBMock) Bank() BankDB {
	return &m.BankMock
}

func (m *DBMock) NominalAccount() NominalAccountDB {
	return &m.NominalAccountMock
}

func (m *DBMock) Offer() OfferDB {
	return &m.OfferMock
}

func (m *DBMock) Product() ProductDB {
	return &m.ProductMock
}

func (m *DBMock) Role() RoleDB {
	return &m.RoleMock
}

func (m *DBMock) Token() TokenDB {
	return &m.TokenMock
}

func (m *DBMock) Transaction() TransactionDB {
	return &m.TransactionMock
}

func (m *DBMock) Tx() TXDB {
	return &m.TxMock
}

func (m *DBMock) User() UserDB {
	return &m.UserMock
}

func (m *DBMock) UserForm() UserFormDB {
	return &m.UserFormMock
}

func (m *NominalAccountDBMock) Create(account *models.NominalAccount) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *NominalAccountDBMock) Get(id string) (models.NominalAccount, error) {
	args := m.Called(id)
	return args.Get(0).(models.NominalAccount), args.Error(1)
}

func (m *NominalAccountDBMock) Pagination(first *int, after *string, filter *models.NominalAccountsFilter) (models.NominalAccountsConnection, error) {
	args := m.Called(first, after, filter)
	return args.Get(0).(models.NominalAccountsConnection), args.Error(1)
}

func (m *NominalAccountDBMock) Take(config NominalAccountTakeConfig) (models.NominalAccount, error) {
	args := m.Called(config)
	return args.Get(0).(models.NominalAccount), args.Error(1)
}

func (m *NominalAccountDBMock) Update(account *models.NominalAccount) error {
	args := m.Called(account)
	return args.Error(0)
}

func (m *OfferDBMock) Create(offer *models.Offer) error {
	args := m.Called(offer)
	return args.Error(0)
}

func (m *OfferDBMock) Get(id string) (models.Offer, error) {
	args := m.Called(id)
	return args.Get(0).(models.Offer), args.Error(1)
}

func (m *OfferDBMock) Pagination(first *int, after *string, filter *models.OffersFilter) (models.OffersConnection, error) {
	args := m.Called(first, after, filter)
	return args.Get(0).(models.OffersConnection), args.Error(1)
}

func (m *OfferDBMock) Take(config OfferTakeConfig) (models.Offer, error) {
	args := m.Called(config)
	return args.Get(0).(models.Offer), args.Error(1)
}

func (m *OfferDBMock) Update(offer *models.Offer) error {
	args := m.Called(offer)
	return args.Error(0)
}

func (m *ProductDBMock) Create(product *models.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *ProductDBMock) Get(id string) (models.Product, error) {
	args := m.Called(id)
	return args.Get(0).(models.Product), args.Error(1)
}

func (m *ProductDBMock) GetCreator(product models.Product) (models.User, error) {
	args := m.Called(product)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *ProductDBMock) GetOwner(product models.Product) (models.User, error) {
	args := m.Called(product)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *ProductDBMock) Pagination(first *int, after *string, filter *models.ProductsFilter) (models.ProductsConnection, error) {
	args := m.Called(first, after, filter)
	return args.Get(0).(models.ProductsConnection), args.Error(1)
}

func (m *ProductDBMock) Update(product *models.Product) error {
	args := m.Called(product)
	return args.Error(0)
}

func (m *RoleDBMock) Find(config RoleFindConfig) ([]models.Role, error) {
	args := m.Called(config)
	return args.Get(0).([]models.Role), args.Error(1)
}

func (m *TXDBMock) Commit() error {
	args := m.Called()
	return args.Error(0)
}

func (m *TXDBMock) DB() DB {
	args := m.Called()
	return args.Get(0).(DB)
}

func (m *TXDBMock) Rollback() {
	m.Called()
	return
}

func (m *TokenDBMock) Create(token *models.Token) error {
	args := m.Called(token)
	return args.Error(0)
}

func (m *TokenDBMock) GetUser(token models.Token) (models.User, error) {
	args := m.Called(token)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *TokenDBMock) Take(config TokenTakeConfig) (models.Token, error) {
	args := m.Called(config)
	return args.Get(0).(models.Token), args.Error(1)
}

func (m *TokenDBMock) Update(token *models.Token) error {
	args := m.Called(token)
	return args.Error(0)
}

func (m *TransactionDBMock) Create(tr *models.Transaction) error {
	args := m.Called(tr)
	return args.Error(0)
}

func (m *TransactionDBMock) Find(config TransactionFindConfig) ([]models.Transaction, error) {
	args := m.Called(config)
	return args.Get(0).([]models.Transaction), args.Error(1)
}

func (m *TransactionDBMock) Get(id int) (models.Transaction, error) {
	args := m.Called(id)
	return args.Get(0).(models.Transaction), args.Error(1)
}

func (m *TransactionDBMock) Pagination(first *int, after *string, filter *models.TransactionsFilter) (models.TransactionsConnection, error) {
	args := m.Called(first, after, filter)
	return args.Get(0).(models.TransactionsConnection), args.Error(1)
}

func (m *TransactionDBMock) Take(config TransactionTakeConfig) (models.Transaction, error) {
	args := m.Called(config)
	return args.Get(0).(models.Transaction), args.Error(1)
}

func (m *TransactionDBMock) Update(tr *models.Transaction) error {
	args := m.Called(tr)
	return args.Error(0)
}

func (m *UserDBMock) Create(user *models.User) error {
	args := m.Called(user)
	return args.Error(0)
}

func (m *UserDBMock) Get(id string) (models.User, error) {
	args := m.Called(id)
	return args.Get(0).(models.User), args.Error(1)
}

func (m *UserDBMock) LastApprovedUserForm(user models.User) (models.UserForm, error) {
	args := m.Called(user)
	return args.Get(0).(models.UserForm), args.Error(1)
}

func (m *UserDBMock) MostRelevantUserForm(user models.User) (models.UserForm, error) {
	args := m.Called(user)
	return args.Get(0).(models.UserForm), args.Error(1)
}

func (m *UserDBMock) Pagination(first *int, after *string, filter *models.UsersFilter) (models.UsersConnection, error) {
	args := m.Called(first, after, filter)
	return args.Get(0).(models.UsersConnection), args.Error(1)
}

func (m *UserFormDBMock) Create(form *models.UserForm) error {
	args := m.Called(form)
	return args.Error(0)
}

func (m *UserFormDBMock) Get(id string) (models.UserForm, error) {
	args := m.Called(id)
	return args.Get(0).(models.UserForm), args.Error(1)
}

func (m *UserFormDBMock) GetLoginForm(input models.LoginInput) (models.UserForm, error) {
	args := m.Called(input)
	return args.Get(0).(models.UserForm), args.Error(1)
}

func (m *UserFormDBMock) Pagination(first *int, after *string, filter *models.UserFormsFilter) (models.UserFormsConnection, error) {
	args := m.Called(first, after, filter)
	return args.Get(0).(models.UserFormsConnection), args.Error(1)
}

func (m *UserFormDBMock) Take(config UserFormTakeConfig) (models.UserForm, error) {
	args := m.Called(config)
	return args.Get(0).(models.UserForm), args.Error(1)
}

func (m *UserFormDBMock) Update(form *models.UserForm) error {
	args := m.Called(form)
	return args.Error(0)
}
