package entity

import "github.com/asaskevich/govalidator"

func init() {
	govalidator.TagMap["transactionCategory"] = govalidator.Validator(func(str string) bool {
		res := str == TRANSACTION_CATEGORY_FOOD.String()
		res = res || str == TRANSACTION_CATEGORY_TRANSPORT.String()
		res = res || str == TRANSACTION_CATEGORY_ENTERTAINMENT.String()
		res = res || str == TRANSACTION_CATEGORY_HEALTH.String()
		res = res || str == TRANSACTION_CATEGORY_BILLS.String()
		res = res || str == TRANSACTION_CATEGORY_EDUCATION.String()
		res = res || str == TRANSACTION_CATEGORY_SHOPPING.String()
		res = res || str == TRANSACTION_CATEGORY_INVESTMENT.String()
		res = res || str == TRANSACTION_CATEGORY_SALARY.String()
		res = res || str == TRANSACTION_CATEGORY_OTHERS.String()
		return res
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

func newTransactionCategory[T TransactionCategory | int](transactionCategory T) *TransactionCategory {
	v := (TransactionCategory)(transactionCategory)
	return &v
}

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
	case TRANSACTION_CATEGORY_OTHERS:
		return "others"
	}
	return ""
}

func (e *TransactionCategory) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}
