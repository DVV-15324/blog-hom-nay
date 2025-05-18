package categories

import (
	"bloghomnay-project/common"
	"context"
	"net/http"
)

func (c *BusinessCategories) BusinessDeleteCategories(ctx context.Context, id int) *common.AppError {
	err := c.bz.DeleteCategories(ctx, id)
	if err != nil {
		app := common.NewAppError(404, http.StatusText(404), err)
		return app
	}
	return nil
}
