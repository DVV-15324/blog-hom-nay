package categories

import (
	"bloghomnay-project/common"
	entityCategories "bloghomnay-project/services/entity/categories"
	"context"
	"net/http"
)

func (c *BusinessCategories) BusinessGetALLCategories(ctx context.Context) ([]entityCategories.Categories, *common.AppError) {
	cates, err := c.bz.GetALLCategories(ctx)
	if err != nil {
		app := common.NewAppError(404, http.StatusText(404), err)
		return nil, app
	}
	return cates, nil
}
