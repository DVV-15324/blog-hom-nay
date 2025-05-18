package categories

import (
	entityCategories "bloghomnay-project/services/entity/categories"
	"context"
)

type Reponsitorycategories interface {
	DeleteCategories(cxt context.Context, id int) error
	GetALLCategories(ctx context.Context) ([]entityCategories.Categories, error)
	CreateCategories(cxt context.Context, categories *entityCategories.CreateCategory) error
	UpdateCategory(cxt context.Context, categories *entityCategories.UpdateCategory, id int) error
}

type BusinessCategories struct {
	bz Reponsitorycategories
}

func NewBusinessCategories(bz Reponsitorycategories) *BusinessCategories {
	return &BusinessCategories{
		bz: bz,
	}
}
