package entity

import (
	"gorm.io/gorm"
	"time"
)

type Timestamp struct {
	CreatedAt time.Time `gorm:"type:timestamp with time zone" json:"created_at"`
	UpdateAt  time.Time `gorm:"type:timestamp with time zone" json:"update_at"`
	DeletedAt gorm.DeletedAt
}

type Authorization struct {
	Token string `json:"token"`
	Role  string `json:"role"`
}
