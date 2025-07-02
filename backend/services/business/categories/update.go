package categories

import (
	"bloghomnay-project/common"
	entityCategories "bloghomnay-project/services/entity/categories"
	"context"
	"net/http"
)

func (c *BusinessCategories) BusinessUpdateCategories(ctx context.Context, data *entityCategories.UpdateCategory, id int) *common.AppError {
	er := data.Validate()
	if er != nil {
		app := common.NewAppError(400, http.StatusText(400), er)
		return app
	}
	err := c.bz.UpdateCategory(ctx, data, id)
	if err != nil {
		app := common.NewAppError(404, http.StatusText(404), err)
		return app
	}
	return nil
}
