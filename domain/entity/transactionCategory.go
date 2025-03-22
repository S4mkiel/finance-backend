package entity

import (
	"errors"
	"github.com/asaskevich/govalidator"
)

func init() {
	// Criando um mapa de categorias válidas para reduzir repetições
	validCategories := map[string]bool{
		TRANSACTION_CATEGORY_FOOD.String():          true,
		TRANSACTION_CATEGORY_TRANSPORT.String():     true,
		TRANSACTION_CATEGORY_ENTERTAINMENT.String(): true,
		TRANSACTION_CATEGORY_HEALTH.String():        true,
		TRANSACTION_CATEGORY_BILLS.String():         true,
		TRANSACTION_CATEGORY_EDUCATION.String():     true,
		TRANSACTION_CATEGORY_SHOPPING.String():      true,
		TRANSACTION_CATEGORY_INVESTMENT.String():    true,
		TRANSACTION_CATEGORY_SALARY.String():        true,
		TRANSACTION_CATEGORY_OTHERS.String():        true,
	}

	govalidator.TagMap["transactionCategory"] = govalidator.Validator(func(str string) bool {
		return validCategories[str]
	})

	govalidator.SetFieldsRequiredByDefault(true)
}

type TransactionCategory int

const (
	TRANSACTION_CATEGORY_FOOD TransactionCategory = iota
	TRANSACTION_CATEGORY_TRANSPORT
	TRANSACTION_CATEGORY_ENTERTAINMENT
	TRANSACTION_CATEGORY_HEALTH
	TRANSACTION_CATEGORY_BILLS
	TRANSACTION_CATEGORY_EDUCATION
	TRANSACTION_CATEGORY_SHOPPING
	TRANSACTION_CATEGORY_INVESTMENT
	TRANSACTION_CATEGORY_SALARY
	TRANSACTION_CATEGORY_OTHERS
)

func (t TransactionCategory) String() string {
	switch t {
	case TRANSACTION_CATEGORY_FOOD:
		return "food"
	case TRANSACTION_CATEGORY_TRANSPORT:
		return "transport"
	case TRANSACTION_CATEGORY_ENTERTAINMENT:
		return "entertainment"
	case TRANSACTION_CATEGORY_HEALTH:
		return "health"
	case TRANSACTION_CATEGORY_BILLS:
		return "bills"
	case TRANSACTION_CATEGORY_EDUCATION:
		return "education"
	case TRANSACTION_CATEGORY_SHOPPING:
		return "shopping"
	case TRANSACTION_CATEGORY_INVESTMENT:
		return "investment"
	case TRANSACTION_CATEGORY_SALARY:
		return "salary"
	case TRANSACTION_CATEGORY_OTHERS:
		return "others"
	default:
		return ""
	}
}

func NewTransactionCategory[T int | string](value T) (*TransactionCategory, error) {
	var v TransactionCategory

	switch val := any(value).(type) {
	case int:
		if val < int(TRANSACTION_CATEGORY_FOOD) || val > int(TRANSACTION_CATEGORY_OTHERS) {
			return nil, errors.New("invalid transaction category")
		}
		v = TransactionCategory(val)
	case string:
		switch val {
		case "food":
			v = TRANSACTION_CATEGORY_FOOD
		case "transport":
			v = TRANSACTION_CATEGORY_TRANSPORT
		case "entertainment":
			v = TRANSACTION_CATEGORY_ENTERTAINMENT
		case "health":
			v = TRANSACTION_CATEGORY_HEALTH
		case "bills":
			v = TRANSACTION_CATEGORY_BILLS
		case "education":
			v = TRANSACTION_CATEGORY_EDUCATION
		case "shopping":
			v = TRANSACTION_CATEGORY_SHOPPING
		case "investment":
			v = TRANSACTION_CATEGORY_INVESTMENT
		case "salary":
			v = TRANSACTION_CATEGORY_SALARY
		case "others":
			v = TRANSACTION_CATEGORY_OTHERS
		default:
			return nil, errors.New("invalid transaction category string")
		}
	default:
		return nil, errors.New("unsupported transaction category format")
	}

	return &v, nil
}

func (t TransactionCategory) IsValid() error {
	if t < TRANSACTION_CATEGORY_FOOD || t > TRANSACTION_CATEGORY_OTHERS {
		return errors.New("invalid transaction category")
	}
	return nil
}
