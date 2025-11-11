package helper

import (
	"go-ecommerce-app/internal/schema"

	"gorm.io/gorm"
)

func RunMigrations(db *gorm.DB) error {
	err := db.AutoMigrate(&schema.User{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&schema.BankAccount{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&schema.Category{})
	if err != nil {
		return err
	}
	err = db.AutoMigrate(&schema.Product{})
	if err != nil {
		return err
	}
	return nil
}
