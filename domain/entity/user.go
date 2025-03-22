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

type User struct {
	Base     `json:",inline" valid:"-"`
	Name     *string `json:"name,omitempty" valid:"required"`
	Email    *string `json:"email,omitempty" valid:"email,required"`
	Password *string `json:"password,omitempty" valid:"required"`
}

func NewUser(
	id *string,
	name *string,
	password *string,
	email *string,
) (*User, error) {
	newPass, err := utils.CryptPassword(password)
	if err != nil {
		return nil, err
	}

	user := User{
		Name:     name,
		Email:    email,
		Password: newPass,
	}

	if id != nil {
		user.ID = id
	} else {
		user.ID = utils.PString(uuid.NewV4().String())
	}

	user.CreatedAt = utils.PTime(time.Now())

	if err := user.isValid(); err != nil {
		return nil, err
	}

	return &user, nil
}

func (p *User) isValid() error {
	_, err := govalidator.ValidateStruct(p)
	return err
}
