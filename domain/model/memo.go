package model

import "time"

type Memo struct {
	ID        uint64    `gorm:"primaryKey" json:"id"`
	UserID    uint64    `json:"user_id"`
	Content   string    `json:"content"`
	CreatedAt *time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
