package controllers

import (
	"errors"
	"go-ecommerce-app/configs"
	functions "go-ecommerce-app/internal/db.functions"
	"go-ecommerce-app/internal/dto"
	"go-ecommerce-app/internal/helper"
	"go-ecommerce-app/internal/schema"

	"github.com/google/uuid"
)

type CatalogContoller struct {
	CatalogDB functions.CatalogDBFunction
	UserDB    functions.UserDBFunction
	Auth      helper.Auth
	Config    configs.AppConfig
}

func (r CatalogContoller) CreateCategory(id uuid.UUID, input dto.AddCategory) (dto.CategoryResponse, error) {

	seller, err := r.UserDB.FindUserById(id)

	if err != nil {
		return dto.CategoryResponse{}, errors.New("cannot find user")
	}

	category := schema.Category{
		Name:     input.Name,
		Owner:    seller.ID,
		ImageUrl: input.ImageUrl,
	}

	data, err := r.CatalogDB.CreateCategory(category)

	if err != nil {
		return dto.CategoryResponse{}, errors.New("failed to create category")
	}

	return dto.ToCategoryResponse(data), nil
}
