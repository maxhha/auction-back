// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package model

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

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

type Bank struct {
	ID      string       `json:"id"`
	Name    string       `json:"name"`
	Account *BankAccount `json:"account"`
}

type BankAccount struct {
	ID           string                  `json:"id"`
	Bank         *Bank                   `json:"bank"`
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

type IncreaseBalanceInput struct {
	UserID string  `json:"userId"`
	Amount float64 `json:"amount"`
}

type IncreaseBalanceResult struct {
	User *User `json:"user"`
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

type RegisterResult struct {
	Token string `json:"token"`
}

type RemoveOfferInput struct {
	OfferID string `json:"offerId"`
}

type RemoveOfferResult struct {
	Status string `json:"status"`
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

type Transaction struct {
	ID          string                `json:"id"`
	Date        time.Time             `json:"date"`
	Status      TransactionStatusEnum `json:"status"`
	Type        TransactionTypeEnum   `json:"type"`
	Currency    CurrencyEnum          `json:"currency"`
	Amount      float64               `json:"amount"`
	Error       *string               `json:"error"`
	AccountFrom Account               `json:"accountFrom"`
	AccountTo   Account               `json:"accountTo"`
}

type TransactionsConnection struct {
	PageInfo *PageInfo                     `json:"pageInfo"`
	Edges    []*TransactionsConnectionEdge `json:"edges"`
}

type TransactionsConnectionEdge struct {
	Cursor string       `json:"cursor"`
	Node   *Transaction `json:"node"`
}

type UserAccount struct {
	ID           string                  `json:"id"`
	Bank         *Bank                   `json:"bank"`
	Transactions *TransactionsConnection `json:"transactions"`
	User         *User                   `json:"user"`
}

func (UserAccount) IsAccount() {}

type UserAccountsConnection struct {
	PageInfo *PageInfo                     `json:"pageInfo"`
	Edges    []*UserAccountsConnectionEdge `json:"edges"`
}

type UserAccountsConnectionEdge struct {
	Cursor string       `json:"cursor"`
	Node   *UserAccount `json:"node"`
}

type UsersConnection struct {
	PageInfo *PageInfo              `json:"pageInfo"`
	Edges    []*UsersConnectionEdge `json:"edges"`
}

type UsersConnectionEdge struct {
	Cursor string `json:"cursor"`
	Node   *User  `json:"node"`
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

type TransactionStatusEnum string

const (
	TransactionStatusEnumOffering   TransactionStatusEnum = "OFFERING"
	TransactionStatusEnumProcessing TransactionStatusEnum = "PROCESSING"
	TransactionStatusEnumError      TransactionStatusEnum = "ERROR"
	TransactionStatusEnumOk         TransactionStatusEnum = "OK"
)

var AllTransactionStatusEnum = []TransactionStatusEnum{
	TransactionStatusEnumOffering,
	TransactionStatusEnumProcessing,
	TransactionStatusEnumError,
	TransactionStatusEnumOk,
}

func (e TransactionStatusEnum) IsValid() bool {
	switch e {
	case TransactionStatusEnumOffering, TransactionStatusEnumProcessing, TransactionStatusEnumError, TransactionStatusEnumOk:
		return true
	}
	return false
}

func (e TransactionStatusEnum) String() string {
	return string(e)
}

func (e *TransactionStatusEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TransactionStatusEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TransactionStatusEnum", str)
	}
	return nil
}

func (e TransactionStatusEnum) MarshalGQL(w io.Writer) {
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
