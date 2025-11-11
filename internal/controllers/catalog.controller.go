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
		Name:     *input.Name,
		Owner:    seller.ID,
		ImageUrl: *input.ImageUrl,
	}

	data, err := r.CatalogDB.CreateCategory(category)

	if err != nil {
		return dto.CategoryResponse{}, err
	}

	return dto.ToCategoryResponse(data), nil
}

func (r CatalogContoller) FindCategories() ([]dto.CategoryResponse, error) {

	categories, err := r.CatalogDB.FindCategories()

	if err != nil {
		return []dto.CategoryResponse{}, errors.New("failed to find categories")
	}

	mashalledCategories := []dto.CategoryResponse{}

	for _, category := range categories {
		mashalledCategories = append(mashalledCategories, dto.ToCategoryResponse(*category))
	}

	return mashalledCategories, nil
}

func (r CatalogContoller) FindCategoryById(id uuid.UUID) (dto.CategoryResponse, error) {

	category, err := r.CatalogDB.FindCategoryById(id)

	if err != nil {
		return dto.CategoryResponse{}, errors.New("failed to find category")
	}

	return dto.ToCategoryResponse(category), nil
}

func (r CatalogContoller) EditCategory(id uuid.UUID, input dto.AddCategory) (dto.CategoryResponse, error) {

	category, err := r.CatalogDB.FindCategoryById(id)

	if err != nil {
		return dto.CategoryResponse{}, errors.New("failed to find category")
	}

	if input.Name != nil {
		category.Name = *input.Name
	}
	if input.ImageUrl != nil {
		category.ImageUrl = *input.ImageUrl
	}

	data, err := r.CatalogDB.EditCategory(id, category)

	if err != nil {
		return dto.CategoryResponse{}, errors.New("failed to edit category")
	}

	return dto.ToCategoryResponse(data), nil
}

func (r CatalogContoller) DeleteCategory(id uuid.UUID) error {
	if err := r.CatalogDB.DeleteCategory(id); err != nil {
		return errors.New("failed to delete category")
	}
	return nil
}
