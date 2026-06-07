package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Expense struct {
	Id            int             `json:"id" gorm:"primaryKey"`
	UserId        int             `json:"user_id"`
	CategoryId    *int            `json:"category_id"`
	Amount        decimal.Decimal `json:"amount" gorm:"type:numeric(10,2);not null"`
	Description   *string         `json:"description"`
	ExpenseDate   time.Time       `json:"expense_date" gorm:"column:expenses_date;type:date;not null"`
	PaymentMethod *string         `json:"payment_method"`
	IsRecurring   bool            `json:"is_recurring"`
	CreatedAt     time.Time       `gorm:"type:timestamptz"`
	UpdatedAt     *time.Time      `gorm:"type:timestamptz"`
}
