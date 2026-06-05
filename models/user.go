package models

import "time"

type UserRole string

const (
	UserRoleBasic   UserRole = "basic"
	UserRolePremium UserRole = "premium"
	UserRoleAdmin   UserRole = "admin"
)

type User struct {
	Id              int        `json:"id" gorm:"primaryKey"`
	GoogleId        *string    `json:"google_id"`
	Username        *string    `json:"username"`
	Email           string     `json:"email"`
	Password        string     `json:"-"`
	AvatarUrl       *string    `json:"avatar_url"`
	ReminderEnabled bool       `json:"reminder_enabled"`
	ReminderDays    int        `json:"reminder_days"`
	Role            UserRole   `json:"role" gorm:"type:user_role;default:'basic'"`
	Timezone        string     `json:"timezone" gorm:"default:'UTC'"`
	CreatedAt       time.Time  `gorm:"type:timestamptz"`
	UpdatedAt       *time.Time `gorm:"type:timestamptz"`
}
