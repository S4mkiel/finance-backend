package entity

import (
	"github.com/S4mkiel/finance-backend/utils"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
)

type Budget struct {
	Base     `json:",inline" valid:"-"`
	Category *TransactionCategory `json:"category,omitempty" valid:"-"`
	Limit    *float64             `json:"limit,omitempty" valid:"-"`
	UserId   *string              `json:"-" valid:"uuid"`
	User     *User                `json:"user" valid:"-"`
}

func NewBudget(id *string, category *TransactionCategory, limit *float64, user *User) (*Budget, error) {
	budget := Budget{
		Category: category,
		Limit:    limit,
		UserId:   user.ID,
		User:     user,
	}

	if id != nil {
		budget.ID = id
	} else {
		budget.ID = utils.PString(uuid.NewV4().String())
	}

	if err := budget.isValid(); err != nil {
		return nil, err
	}

	return &budget, nil
}

func (p Budget) isValid() error {
	_, err := govalidator.ValidateStruct(p)
	return err
}
