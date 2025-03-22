package migration

import (
	"time"

	"github.com/go-gormigrate/gormigrate/v2"
	"gorm.io/gorm"
)

var M_2025032002200 *gormigrate.Migration = func() *gormigrate.Migration {
	type BaseTimestamps struct {
		CreatedAt *time.Time `gorm:"column:created_at;autoCreateTime"`
		UpdatedAt *time.Time `gorm:"column:updated_at;autoUpdateTime"`
		DeletedAt *time.Time `gorm:"column:deleted_at;index;autoDeleteTime"`
	}

	type BaseID struct {
		ID *string `gorm:"type:uuid;primaryKey"`
	}

	type Base struct {
		BaseID
		BaseTimestamps
	}
	type TransactionType int

	const (
		Income TransactionType = iota
		Expense
	)

	type TransactionCategory int

	const (
		Food TransactionCategory = iota
		Transport
		Entertainment
		Health
		Bills
		Education
		Shopping
		Investment
		Salary
		Other
	)

	type User struct {
		Base
		Name     *string `gorm:"type:varchar(255)"`
		Email    *string `gorm:"type:varchar(255);unique"`
		Password *string `gorm:"type:varchar(255)"`
	}

	type Transaction struct {
		Base
		Amount          *float64
		TransactionType *TransactionType
		Category        *TransactionCategory
		Date            *time.Time `gorm:"column:date;type:timestamp"`
		Notes           *string    `gorm:"type:text"`
		Currency        *string    `gorm:"type:varchar(255)"`
		UserID          *string    `gorm:"type:varchar(255)"`
		User            *User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	}

	type RecurringTransaction struct {
		Base
		Amount          *float64 `gorm:"type:decimal(20,8)"`
		TransactionType *TransactionType
		Category        *TransactionCategory
		Frequency       *string    `gorm:"type:varchar(20)"`
		NextDate        *time.Time `gorm:"column:next_date;type:timestamp"`
		UserID          *string    `gorm:"type:varchar(255)"`
		User            *User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	}

	type Goal struct {
		Base
		Title    *string    `json:"title" valid:"required"`
		Target   *float64   `json:"target" valid:"required"`
		Current  *float64   `gorm:"type:decimal(20,8)"`
		Deadline *time.Time `gorm:"column:dead_line;type:timestamp'"`
		UserID   *string    `gorm:"type:varchar(255)"`
		User     *User      `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	}

	type Budget struct {
		Base
		Category *TransactionCategory
		Limit    *float64 `gorm:"type:decimal(20,8)"`
		UserID   *string  `gorm:"type:varchar(255)"`
		User     *User    `gorm:"foreignKey:UserID;references:ID;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
	}

	return &gormigrate.Migration{
		ID: "2025032002200-init",
		Migrate: func(db *gorm.DB) error {
			return db.Transaction(
				func(tx *gorm.DB) error {
					if err := tx.AutoMigrate(
						&User{},
						&Transaction{},
						&RecurringTransaction{},
						&Budget{},
						&Goal{},
					); err != nil {
						return err
					}

					return nil
				},
			)
		},
		Rollback: func(db *gorm.DB) error {
			return db.Migrator().DropTable(
				&User{},
				&Transaction{},
				&RecurringTransaction{},
				&Budget{},
				&Goal{},
			)
		},
	}
}()
