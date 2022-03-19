// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"auction-back/db"
	"fmt"
	"io"
	"strconv"
	"time"
)

// Nominal account
type Account interface {
	IsAccount()
}

type AccountsConnection struct {
	PageInfo *PageInfo                 `json:"pageInfo"`
	Edges    []*AccountsConnectionEdge `json:"edges"`
}

type AccountsConnectionEdge struct {
	Cursor string  `json:"cursor"`
	Node   Account `json:"node"`
}

// Bank that is cooperated with platform
type Bank struct {
	ID string `json:"id"`
	// Name of bank
	Name string `json:"name"`
	// Special account of that bank
	Account *BankAccount `json:"account"`
}

// Special account for banks. Amount on this account is always nonpositve
type BankAccount struct {
	ID string `json:"id"`
	// Owner of account. Each bank have one special account
	Bank *Bank `json:"bank"`
	// All transactions in which the account is involved
	Transactions *TransactionsConnection `json:"transactions"`
}

func (BankAccount) IsAccount() {}

type CreateOfferInput struct {
	ProductID string  `json:"productId"`
	Amount    float64 `json:"amount"`
}

type CreateOfferResult struct {
	Offer *Offer `json:"offer"`
}

type CreateProductInput struct {
	Name        string  `json:"name"`
	Description *string `json:"description"`
}

type CreateProductResult struct {
	Product *Product `json:"product"`
}

type LoginInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

// Мoney in a specific currency
type Money struct {
	Amount   float64      `json:"amount"`
	Currency CurrencyEnum `json:"currency"`
}

type OfferProductInput struct {
	ProductID string `json:"productId"`
}

type OfferProductResult struct {
	Product *Product `json:"product"`
}

type OffersConnection struct {
	PageInfo *PageInfo               `json:"pageInfo"`
	Edges    []*OffersConnectionEdge `json:"edges"`
}

type OffersConnectionEdge struct {
	Cursor string `json:"cursor"`
	Node   *Offer `json:"node"`
}

type PageInfo struct {
	HasNextPage     bool    `json:"hasNextPage"`
	HasPreviousPage bool    `json:"hasPreviousPage"`
	StartCursor     *string `json:"startCursor"`
	EndCursor       *string `json:"endCursor"`
}

// Product image
type ProductImage struct {
	ID       string `json:"id"`
	Filename string `json:"filename"`
	Path     string `json:"path"`
}

type ProductsConnection struct {
	PageInfo *PageInfo                 `json:"pageInfo"`
	Edges    []*ProductsConnectionEdge `json:"edges"`
}

type ProductsConnectionEdge struct {
	Cursor string   `json:"cursor"`
	Node   *Product `json:"node"`
}

type RemoveOfferInput struct {
	OfferID string `json:"offerId"`
}

type RemoveOfferResult struct {
	Status string `json:"status"`
}

type RequestSetUserEmailInput struct {
	Email string `json:"email"`
}

type RequestSetUserPhoneInput struct {
	Phone string `json:"phone"`
}

type SellProductInput struct {
	ProductID string `json:"productId"`
}

type SellProductResult struct {
	Product *Product `json:"product"`
}

type TakeOffProductInput struct {
	ProductID string `json:"productId"`
}

type TakeOffProductResult struct {
	Product *Product `json:"product"`
}

// Used for actions activation
type TokenInput struct {
	Token string `json:"token"`
}

// Used for login and registration
type TokenResult struct {
	Token string `json:"token"`
}

type Transaction struct {
	ID string `json:"id"`
	// Time of apply this transaction
	Date *time.Time `json:"date"`
	// Current state
	State TransactionStateEnum `json:"state"`
	// Transaction type
	Type TransactionTypeEnum `json:"type"`
	// Transaction currency
	Currency CurrencyEnum `json:"currency"`
	// Transaction amount
	Amount float64 `json:"amount"`
	// Error message for state = ERROR or FAILED
	Error *string `json:"error"`
	// Offer for type = BUY
	Offer *Offer `json:"offer"`
	// From account
	AccountFrom Account `json:"accountFrom"`
	// To account
	AccountTo Account `json:"accountTo"`
}

type TransactionsConnection struct {
	PageInfo *PageInfo                     `json:"pageInfo"`
	Edges    []*TransactionsConnectionEdge `json:"edges"`
}

type TransactionsConnectionEdge struct {
	Cursor string       `json:"cursor"`
	Node   *Transaction `json:"node"`
}

type UpdateUserPasswordInput struct {
	OldPassword *string `json:"oldPassword"`
	Password    string  `json:"password"`
}

// Nominal account that was created for client
type UserAccount struct {
	ID string `json:"id"`
	// Bank in which the account was created
	Bank *Bank `json:"bank"`
	// All transactions in which the account is involved
	Transactions *TransactionsConnection `json:"transactions"`
	// Owner of account
	User *db.User `json:"user"`
}

func (UserAccount) IsAccount() {}

type UserAccountsConnection struct {
	PageInfo *PageInfo                     `json:"pageInfo"`
	Edges    []*UserAccountsConnectionEdge `json:"edges"`
}

// Connection with UserAccount only
type UserAccountsConnectionEdge struct {
	Cursor string       `json:"cursor"`
	Node   *UserAccount `json:"node"`
}

// UserFrom with all required fields filled in
type UserFormFilled struct {
	// User email
	Email string `json:"email"`
	// User phone
	Phone string `json:"phone"`
	// User name
	Name string `json:"name"`
}

type UserFormsConnection struct {
	PageInfo *PageInfo                  `json:"pageInfo"`
	Edges    []*UserFormsConnectionEdge `json:"edges"`
}

type UserFormsConnectionEdge struct {
	Cursor string       `json:"cursor"`
	Node   *db.UserForm `json:"node"`
}

type UserResult struct {
	User *db.User `json:"user"`
}

type UsersConnection struct {
	PageInfo *PageInfo              `json:"pageInfo"`
	Edges    []*UsersConnectionEdge `json:"edges"`
}

type UsersConnectionEdge struct {
	Cursor string   `json:"cursor"`
	Node   *db.User `json:"node"`
}

type CurrencyEnum string

const (
	CurrencyEnumRub CurrencyEnum = "RUB"
	CurrencyEnumEur CurrencyEnum = "EUR"
	CurrencyEnumUsd CurrencyEnum = "USD"
)

var AllCurrencyEnum = []CurrencyEnum{
	CurrencyEnumRub,
	CurrencyEnumEur,
	CurrencyEnumUsd,
}

func (e CurrencyEnum) IsValid() bool {
	switch e {
	case CurrencyEnumRub, CurrencyEnumEur, CurrencyEnumUsd:
		return true
	}
	return false
}

func (e CurrencyEnum) String() string {
	return string(e)
}

func (e *CurrencyEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = CurrencyEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid CurrencyEnum", str)
	}
	return nil
}

func (e CurrencyEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type OfferStateEnum string

const (
	OfferStateEnumCreated               OfferStateEnum = "CREATED"
	OfferStateEnumCancelled             OfferStateEnum = "CANCELLED"
	OfferStateEnumTransferringMoney     OfferStateEnum = "TRANSFERRING_MONEY"
	OfferStateEnumTransferMoneyFailed   OfferStateEnum = "TRANSFER_MONEY_FAILED"
	OfferStateEnumTransferringProduct   OfferStateEnum = "TRANSFERRING_PRODUCT"
	OfferStateEnumTransferProductFailed OfferStateEnum = "TRANSFER_PRODUCT_FAILED"
	OfferStateEnumSucceeded             OfferStateEnum = "SUCCEEDED"
	OfferStateEnumReturningMoney        OfferStateEnum = "RETURNING_MONEY"
	OfferStateEnumReturnMoneyFailed     OfferStateEnum = "RETURN_MONEY_FAILED"
	OfferStateEnumMoneyReturned         OfferStateEnum = "MONEY_RETURNED"
)

var AllOfferStateEnum = []OfferStateEnum{
	OfferStateEnumCreated,
	OfferStateEnumCancelled,
	OfferStateEnumTransferringMoney,
	OfferStateEnumTransferMoneyFailed,
	OfferStateEnumTransferringProduct,
	OfferStateEnumTransferProductFailed,
	OfferStateEnumSucceeded,
	OfferStateEnumReturningMoney,
	OfferStateEnumReturnMoneyFailed,
	OfferStateEnumMoneyReturned,
}

func (e OfferStateEnum) IsValid() bool {
	switch e {
	case OfferStateEnumCreated, OfferStateEnumCancelled, OfferStateEnumTransferringMoney, OfferStateEnumTransferMoneyFailed, OfferStateEnumTransferringProduct, OfferStateEnumTransferProductFailed, OfferStateEnumSucceeded, OfferStateEnumReturningMoney, OfferStateEnumReturnMoneyFailed, OfferStateEnumMoneyReturned:
		return true
	}
	return false
}

func (e OfferStateEnum) String() string {
	return string(e)
}

func (e *OfferStateEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OfferStateEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OfferStateEnum", str)
	}
	return nil
}

func (e OfferStateEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TransactionStateEnum string

const (
	TransactionStateEnumCreated    TransactionStateEnum = "CREATED"
	TransactionStateEnumCancelled  TransactionStateEnum = "CANCELLED"
	TransactionStateEnumProcessing TransactionStateEnum = "PROCESSING"
	TransactionStateEnumError      TransactionStateEnum = "ERROR"
	TransactionStateEnumSucceeded  TransactionStateEnum = "SUCCEEDED"
	TransactionStateEnumFailed     TransactionStateEnum = "FAILED"
)

var AllTransactionStateEnum = []TransactionStateEnum{
	TransactionStateEnumCreated,
	TransactionStateEnumCancelled,
	TransactionStateEnumProcessing,
	TransactionStateEnumError,
	TransactionStateEnumSucceeded,
	TransactionStateEnumFailed,
}

func (e TransactionStateEnum) IsValid() bool {
	switch e {
	case TransactionStateEnumCreated, TransactionStateEnumCancelled, TransactionStateEnumProcessing, TransactionStateEnumError, TransactionStateEnumSucceeded, TransactionStateEnumFailed:
		return true
	}
	return false
}

func (e TransactionStateEnum) String() string {
	return string(e)
}

func (e *TransactionStateEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TransactionStateEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TransactionStateEnum", str)
	}
	return nil
}

func (e TransactionStateEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type TransactionTypeEnum string

const (
	TransactionTypeEnumDeposit    TransactionTypeEnum = "DEPOSIT"
	TransactionTypeEnumBuy        TransactionTypeEnum = "BUY"
	TransactionTypeEnumFee        TransactionTypeEnum = "FEE"
	TransactionTypeEnumWithdrawal TransactionTypeEnum = "WITHDRAWAL"
)

var AllTransactionTypeEnum = []TransactionTypeEnum{
	TransactionTypeEnumDeposit,
	TransactionTypeEnumBuy,
	TransactionTypeEnumFee,
	TransactionTypeEnumWithdrawal,
}

func (e TransactionTypeEnum) IsValid() bool {
	switch e {
	case TransactionTypeEnumDeposit, TransactionTypeEnumBuy, TransactionTypeEnumFee, TransactionTypeEnumWithdrawal:
		return true
	}
	return false
}

func (e TransactionTypeEnum) String() string {
	return string(e)
}

func (e *TransactionTypeEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TransactionTypeEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TransactionTypeEnum", str)
	}
	return nil
}

func (e TransactionTypeEnum) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
