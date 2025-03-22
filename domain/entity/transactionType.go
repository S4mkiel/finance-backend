package entity

import (
	"errors"
	"github.com/asaskevich/govalidator"
)

func init() {
	govalidator.TagMap["transactionType"] = govalidator.Validator(func(str string) bool {
		return str == TRANSACTION_TYPE_INCOME.String() || str == TRANSACTION_TYPE_EXPENSE.String()
	})

	govalidator.SetFieldsRequiredByDefault(true)
}

type TransactionType int

const (
	TRANSACTION_TYPE_INCOME TransactionType = iota
	TRANSACTION_TYPE_EXPENSE
)

func (t TransactionType) String() string {
	switch t {
	case TRANSACTION_TYPE_INCOME:
		return "income"
	case TRANSACTION_TYPE_EXPENSE:
		return "expense"
	default:
		return ""
	}
}

func NewTransactionType[T int | string](value T) (*TransactionType, error) {
	var v TransactionType

	switch val := any(value).(type) {
	case int:
		if val != int(TRANSACTION_TYPE_INCOME) && val != int(TRANSACTION_TYPE_EXPENSE) {
			return nil, errors.New("invalid transaction type")
		}
		v = TransactionType(val)
	case string:
		switch val {
		case "income":
			v = TRANSACTION_TYPE_INCOME
		case "expense":
			v = TRANSACTION_TYPE_EXPENSE
		default:
			return nil, errors.New("invalid transaction type string")
		}
	default:
		return nil, errors.New("unsupported transaction type format")
	}

	return &v, nil
}

func (t TransactionType) IsValid() error {
	if t != TRANSACTION_TYPE_INCOME && t != TRANSACTION_TYPE_EXPENSE {
		return errors.New("invalid transaction type")
	}
	return nil
}
