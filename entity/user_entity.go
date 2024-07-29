package entity

import (
	"GoGinToDoList/helpers"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	ID          uuid.UUID `gorm:"type:uuid;primary_key;default:uuid_generate_v4()" json:"id"`
	Name        string    `gorm:"column:name" json:"name"`
	PhoneNumber string    `gorm:"column:phone_number" json:"phone_number"`
	Email       string    `gorm:"column:email" json:"email"`
	Password    string    `gorm:"column:password" json:"password"`
	Role        string    `gorm:"column:role" json:"role"`
	Timestamp
}

func (u *User) TableName() string {
	return "users"
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	defer func() {
		if r := recover(); r != nil {
			tx.Rollback()
		}
	}()

	var err error
	u.Password, err = helpers.HashPassword(u.Password)
	if err != nil {
		return err
	}

	return nil
}
