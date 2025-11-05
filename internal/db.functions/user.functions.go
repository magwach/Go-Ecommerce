package functions

import (
	"errors"
	"go-ecommerce-app/internal/schema"
	"log"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type UserDBFunction interface {
	SignUp(user schema.User) (schema.User, error)
	FindUser(email string) (schema.User, error)
	FindUserById(id uuid.UUID) (schema.User, error)
	UpdateUser(id uuid.UUID, user schema.User) (schema.User, error)
}

type userDBFunction struct {
	db *gorm.DB
}

func InitializeUserDBFunction(db *gorm.DB) UserDBFunction {
	return userDBFunction{
		db: db,
	}
}

func (r userDBFunction) SignUp(user schema.User) (schema.User, error) {

	err := r.db.Create(&user).Error

	if err != nil {
		log.Printf("error in creating user: %v", err)
		return schema.User{}, errors.New("error while creating user")
	}

	return user, nil
}

func (r userDBFunction) FindUser(email string) (schema.User, error) {
	return schema.User{}, nil
}

func (r userDBFunction) FindUserById(id uuid.UUID) (schema.User, error) {
	return schema.User{}, nil
}

func (r userDBFunction) UpdateUser(id uuid.UUID, user schema.User) (schema.User, error) {
	return schema.User{}, nil
}
