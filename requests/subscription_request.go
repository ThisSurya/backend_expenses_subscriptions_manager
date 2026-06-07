package requests

import (
	"time"

	"github.com/shopspring/decimal"
)

type SubscriptionRequest struct {
	CategoryId      *int            `json:"category_id"`
	ExpenseId       *int            `json:"expense_id"`
	Name            string          `json:"name" binding:"required"`
	Amount          decimal.Decimal `json:"amount" binding:"required"`
	BillingCycle    string          `json:"billing_cycle" binding:"required"`
	NextBillingDate time.Time       `json:"next_billing_date" binding:"required"`
	IsActive        *bool           `json:"is_active"`
}
