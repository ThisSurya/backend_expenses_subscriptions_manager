package models

import "time"

type Category struct {
	Id        int        `json:"id" gorm:"primaryKey"`
	UserId    int        `json:"user_id"`
	Name      string     `json:"name"`
	Color     string     `json:"color"`
	IconUrl   *string    `json:"icon_url"`
	CreatedAt time.Time  `gorm:"type:timestamptz"`
	UpdatedAt *time.Time `gorm:"type:timestamptz"`
}
