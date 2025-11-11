package functions

import (
	"go-ecommerce-app/internal/schema"

	"gorm.io/gorm"
)

type CatalogDBFunction interface {
	CreateCategory(details schema.Category) (schema.Category, error)
}

type catalogDBFunction struct {
	db *gorm.DB
}

func InitializeCatalogDBFunction(db *gorm.DB) CatalogDBFunction {
	return catalogDBFunction{
		db: db,
	}
}

func (r catalogDBFunction) CreateCategory(details schema.Category) (schema.Category, error) {
    if err := r.db.Create(&details).Error; err != nil {
        return schema.Category{}, err
    }
    return details, nil
}
