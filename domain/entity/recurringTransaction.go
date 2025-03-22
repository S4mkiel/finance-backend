package entity

import (
	"github.com/S4mkiel/finance-backend/utils"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type RecurringTransaction struct {
	Base            `json:",inline" valid:"-"`
	Amount          *float64             `json:"amount,omitempty" valid:"-"`
	TransactionType *TransactionType     `json:"transaction_type,omitempty" valid:"-"`
	Category        *TransactionCategory `json:"category,omitempty" valid:"-"`
	Frequency       *string              `json:"frequency,omitempty" valid:"-"`
	NextDate        *time.Time           `json:"next_date,omitempty" valid:"-"`
	UserID          *string              `json:"-" valid:"-"`
	User            *User                `json:"user" valid:"-"`
}

func NewRecurringTransaction(
	id *string,
	amount *float64,
	transactionType *TransactionType,
	category *TransactionCategory,
	frequency *string,
	nextDate *time.Time,
	user *User,
) (*RecurringTransaction, error) {
	rTransaction := RecurringTransaction{
		Amount:          amount,
		TransactionType: transactionType,
		Category:        category,
		Frequency:       frequency,
		NextDate:        nextDate,
		UserID:          user.ID,
		User:            user,
	}

	if id != nil {
		rTransaction.ID = id
	} else {
		rTransaction.ID = utils.PString(uuid.NewV4().String())
	}

	rTransaction.CreatedAt = utils.PTime(time.Now())

	if err := rTransaction.isValid(); err != nil {
		return nil, err
	}
	return &rTransaction, nil
}

func (p *RecurringTransaction) isValid() error {
	_, err := govalidator.ValidateStruct(p)
	return err
}
