package controllers

import (
	"errors"
	functions "go-ecommerce-app/internal/db.functions"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/internal/helper"
	"go-ecommerce-app/internal/schema"

	"github.com/google/uuid"
)

type UserContoller struct {
	DB functions.UserDBFunction
}

func (s UserContoller) SignUp(input dto.UserSignUp) (string, error) {

	hashedPassword, err := helper.HashPassword(input.Password)

	if err != nil {
		return "", errors.New("failed to hash password")
	}

	_, err = s.DB.SignUp(schema.User{
		Email:    input.Email,
		Password: hashedPassword,
		Phone:    input.Phone,
	})

	if err != nil {
		return "", err
	}

	return hashedPassword, nil
}

func (s UserContoller) FindUserByEmail(email string) (*schema.User, error) {

	user, err := s.DB.FindUserByEmail(email)

	if err != nil {
		return &schema.User{}, err
	}

	return &user, nil
}

func (s UserContoller) Login(email, password string) (string, error) {

	user, err := s.DB.FindUserByEmail(email)

	if err != nil {
		return "", errors.New("user not found")
	}

	valid := helper.CheckPassword(user.Password, password)

	if !valid {
		return "", errors.New("invalid credentials")
	}

	return "", nil
}

func (s UserContoller) GetVerificationCode(u *schema.User) (int, error) {
	return 0, nil
}

func (s UserContoller) VerifyCode(u *schema.User) error {
	return nil
}

func (s UserContoller) CreateProfile(id uuid.UUID, input any) error {
	return nil
}

func (s UserContoller) GetProfile(id uuid.UUID) (*schema.User, error) {
	return nil, nil
}

func (s UserContoller) UpdateProfile(id uuid.UUID, input any) error {
	return nil
}

func (s UserContoller) BecomeSeller(id uuid.UUID, input any) (string, error) {
	return "", nil
}

func (s UserContoller) FindCart(id uuid.UUID) (*schema.Cart, error) {
	return nil, nil
}

func (s UserContoller) CreateCart(input any, u *schema.User) (*schema.Cart, error) {
	return nil, nil
}

func (s UserContoller) CreateOrder(u *schema.User) (int, error) {
	return 0, nil
}

func (s UserContoller) GetOrders(u *schema.User) (*schema.Cart, error) {
	return nil, nil
}

func (s UserContoller) GetOrderById(id, UserId uuid.UUID) (*schema.Cart, error) {
	return nil, nil
}
