package model

import "time"

type User struct {
	ID        uint64     `gorm:"primaryKey" json:"id"`
	Name      string     `json:"name"`
	Email     string     `json:"email"`
	CreatedAt *time.Time `gorm:"autoCreateTime" json:"created_at"`
	UpdatedAt *time.Time `gorm:"autoUpdateTime" json:"updated_at"`
}
