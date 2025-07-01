package tag

import (
	"bloghomnay-project/common"
	entityTag "bloghomnay-project/services/entity/tag"
	"context"
	"net/http"
)

func (c *BusinessTag) BusinessCreateTag(ctx context.Context, data *entityTag.CreateTagForm) *common.AppError {
	er := data.Validate()
	if er != nil {
		app := common.NewAppError(400, http.StatusText(400), er)
		return app
	}
	err := c.bzTag.CreateTag(ctx, data)
	if err != nil {
		app := common.NewAppError(500, http.StatusText(500), err)
		return app
	}
	return nil
}
