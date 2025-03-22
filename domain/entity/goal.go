package entity

import (
	"github.com/S4mkiel/finance-backend/utils"
	"github.com/asaskevich/govalidator"
	uuid "github.com/satori/go.uuid"
	"time"
)

type Goal struct {
	Base     `json:",inline" valid:"-"`
	Title    *string    `json:"title,omitempty" valid:"-"`
	Target   *float64   `json:"target,omitempty" valid:"-"`
	Current  *float64   `json:"current,omitempty" valid:"-"`
	Deadline *time.Time `json:"dead_line,omitempty" valid:"-"`
	UserId   *string    `json:"-" valid:"uuid"`
	User     *User      `json:"user" valid:"-"`
}

func NewGoal(id, title *string, target, current *float64, deadLine *time.Time, user *User) (*Goal, error) {
	goal := Goal{
		Title:    title,
		Target:   target,
		Current:  current,
		Deadline: deadLine,
		UserId:   user.ID,
		User:     user,
	}

	if id != nil {
		goal.ID = id
	} else {
		goal.ID = utils.PString(uuid.NewV4().String())
	}

	if err := goal.isValid(); err != nil {
		return nil, err
	}

	return &goal, nil
}

func (p *Goal) isValid() error {
	_, err := govalidator.ValidateStruct(p)
	return err
}
