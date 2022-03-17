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

type ActivateTokenInput struct {
	Token string `json:"token"`
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

type CreateTokenInput struct {
	Action TokenActionEnum        `json:"action"`
	Data   map[string]interface{} `json:"data"`
}

type CreateUserResult struct {
	Token string `json:"token"`
}

type IncreaseBalanceInput struct {
	UserID string  `json:"userId"`
	Amount float64 `json:"amount"`
}

type IncreaseBalanceResult struct {
	User *User `json:"user"`
}

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
	ID          string               `json:"id"`
	Date        *time.Time           `json:"date"`
	State       TransactionStateEnum `json:"state"`
	Type        TransactionTypeEnum  `json:"type"`
	Currency    CurrencyEnum         `json:"currency"`
	Amount      float64              `json:"amount"`
	Error       *string              `json:"error"`
	Offer       *Offer               `json:"offer"`
	AccountFrom Account              `json:"accountFrom"`
	AccountTo   Account              `json:"accountTo"`
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

type TokenActionEnum string

const (
	TokenActionEnumApproveUserEmail TokenActionEnum = "APPROVE_USER_EMAIL"
	TokenActionEnumApproveUserPhone TokenActionEnum = "APPROVE_USER_PHONE"
)

var AllTokenActionEnum = []TokenActionEnum{
	TokenActionEnumApproveUserEmail,
	TokenActionEnumApproveUserPhone,
}

func (e TokenActionEnum) IsValid() bool {
	switch e {
	case TokenActionEnumApproveUserEmail, TokenActionEnumApproveUserPhone:
		return true
	}
	return false
}

func (e TokenActionEnum) String() string {
	return string(e)
}

func (e *TokenActionEnum) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = TokenActionEnum(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid TokenActionEnum", str)
	}
	return nil
}

func (e TokenActionEnum) MarshalGQL(w io.Writer) {
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
