package requests

import (
	"time"

	"github.com/shopspring/decimal"
)

type ExpenseRequest struct {
	UserId        int             `json:"user_id"`
	CategoryId    *int            `json:"category_id"`
	Amount        decimal.Decimal `json:"amount" binding:"required"`
	Description   *string         `json:"description"`
	ExpenseDate   time.Time       `json:"expense_date" binding:"required"`
	PaymentMethod *string         `json:"payment_method" binding:"required"`
	IsRecurring   bool            `json:"is_recurring" binding:"required"`
}
