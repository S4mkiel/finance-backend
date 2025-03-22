package entity

import (
	"github.com/S4mkiel/finance-backend/utils"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

func init() {
	govalidator.SetFieldsRequiredByDefault(true)
}

type Transaction struct {
	Base            `json:",inline" valid:"-"`
	Amount          *float64             `json:"amount,omitempty" valid:"-"`
	TransactionType *TransactionType     `json:"transaction_type,omitempty" valid:"-"`
	Category        *TransactionCategory `json:"category,omitempty" valid:"-"`
	Date            *time.Time           `json:"date,omitempty" valid:"-"`
	Notes           *string              `json:"notes,omitempty" valid:"-"`
	Currency        *string              `json:"currency,omitempty" valid:"-"`
	UserID          *string              `json:"-" valid:"uuid"`
	User            *User                `json:"user" valid:"-"`
}

func NewTransaction(
	id *string,
	amount *float64,
	transactionType *TransactionType,
	category *TransactionCategory,
	date *time.Time,
	notes *string,
	currency *string,
	user *User,
) (*Transaction, error) {
	transaction := Transaction{
		Amount:          amount,
		TransactionType: transactionType,
		Category:        category,
		Date:            date,
		Notes:           notes,
		Currency:        currency,
		UserID:          user.ID,
		User:            user,
	}

	if id != nil {
		transaction.ID = id
	} else {
		transaction.ID = utils.PString(uuid.NewV4().String())
	}

	transaction.CreatedAt = utils.PTime(time.Now())

	if err := transaction.isValid(); err != nil {
		return nil, err
	}

	return &transaction, nil
}

func (p *Transaction) isValid() error {
	_, err := govalidator.ValidateStruct(p)
	return err
}
