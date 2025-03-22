package entity

import "github.com/asaskevich/govalidator"

func init() {
	govalidator.TagMap["transactionType"] = govalidator.Validator(func(str string) bool {
		res := str == TRANSACTION_TYPE_INCOME.String()
		res = res || str == TRANSACTION_TYPE_EXPENSE.String()
		return res
	})

	govalidator.SetFieldsRequiredByDefault(true)
}

type TransactionType int

const (
	TRANSACTION_TYPE_INCOME TransactionType = iota
	TRANSACTION_TYPE_EXPENSE
)

func newTransactionType[T TransactionType | int](friendshipStatus T) *TransactionType {
	v := (TransactionType)(friendshipStatus)
	return &v
}

func (t TransactionType) String() string {
	switch t {
	case TRANSACTION_TYPE_INCOME:
		return "income"
	case TRANSACTION_TYPE_EXPENSE:
		return "expense"
	}
	return ""
}

func (e *TransactionType) isValid() error {
	_, err := govalidator.ValidateStruct(e)
	return err
}
