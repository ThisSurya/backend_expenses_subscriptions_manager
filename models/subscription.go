package models

import (
	"time"

	"github.com/shopspring/decimal"
)

type Subscription struct {
	Id              int             `json:"id" gorm:"primaryKey"`
	UserId          int             `json:"user_id"`
	CategoryId      *int            `json:"category_id"`
	ExpenseId       *int            `json:"expense_id"`
	Name            string          `json:"name"`
	Amount          decimal.Decimal `json:"amount" gorm:"type:numeric(10,2);not null"`
	BillingCycle    string          `json:"billing_cycle"`
	NextBillingDate time.Time       `json:"next_billing_date" gorm:"type:date;not null"`
	IsActive        *bool           `json:"is_active" gorm:"default:true"`
	CreatedAt       time.Time       `gorm:"type:timestamptz"`
	UpdatedAt       *time.Time      `gorm:"type:timestamptz"`
}
