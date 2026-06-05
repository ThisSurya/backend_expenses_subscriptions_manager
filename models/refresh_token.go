package models

import "time"

type RefreshToken struct {
	Id        int        `json:"id" gorm:"primaryKey"`
	UserId    int        `json:"user_id"`
	Token     string     `json:"token"`
	ExpiresAt time.Time  `json:"expires_at" gorm:"type:timestamptz"`
	RevokedAt *time.Time `json:"revoked_at" gorm:"type:timestamptz"`
	CreatedAt time.Time  `gorm:"type:timestamptz"`
}
