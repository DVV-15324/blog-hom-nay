package categories

import (
	"bloghomnay-project/common"
	entityCategories "bloghomnay-project/services/entity/categories"
	"context"
)

type BusinessCategories interface {
	BusinessCreateCategories(ctx context.Context, data *entityCategories.CreateCategory) *common.AppError
	BusinessDeleteCategories(ctx context.Context, id int) *common.AppError
	BusinessGetALLCategories(ctx context.Context) ([]entityCategories.Categories, *common.AppError)
	BusinessUpdateCategories(ctx context.Context, data *entityCategories.UpdateCategory, id int) *common.AppError
}
type ApiCategories struct {
	bz BusinessCategories
}

func NewApiCategories(bz BusinessCategories) *ApiCategories {
	return &ApiCategories{
		bz: bz,
	}
}
