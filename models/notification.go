package models

import "time"

type NotificationType string

type NotificationStatus string

const (
	NotificationTypeEmail NotificationType = "email"
)

const (
	NotificationStatusPending NotificationStatus = "pending"
	NotificationStatusSent    NotificationStatus = "sent"
	NotificationStatusFailed  NotificationStatus = "failed"
)

type Notification struct {
	Id             int                `json:"id" gorm:"primaryKey"`
	UserId         int                `json:"user_id"`
	SubscriptionId *int               `json:"subscription_id"`
	Type           NotificationType   `json:"type" gorm:"type:notification_type"`
	Subject        string             `json:"subject"`
	Status         NotificationStatus `json:"status" gorm:"type:notification_status"`
	ScheduledAt    time.Time          `json:"scheduled_at" gorm:"type:timestamptz"`
	SentAt         *time.Time         `json:"sent_at" gorm:"type:timestamptz"`
	NumTries       int                `json:"num_tries"`
	MaxRetries     int                `json:"max_retries"`
	ErrorMessage   *string            `json:"error_message"`
	CreatedAt      time.Time          `gorm:"type:timestamptz"`
}
