package migrations

import (
	"GoGinToDoList/entity"
	"fmt"
	"gorm.io/gorm"
)

func Migrate(db *gorm.DB) error {
	if err := db.AutoMigrate(
		&entity.User{},
	); err != nil {
		return err
	}
	fmt.Print("Migrate Done")

	return nil
}
