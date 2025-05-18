package comment

import (
	"bloghomnay-project/common"
	entityComment "bloghomnay-project/services/entity/comment"
	"context"
	"net/http"
)

func (c *BusinessComment) BusinessCreateComment(ctx context.Context, data *entityComment.CreateComment) *common.AppError {
	er := data.Validate()
	if er != nil {
		app := common.NewAppError(400, http.StatusText(400), er)
		return app
	}
	err := c.bzComment.CreateComment(ctx, data)
	if err != nil {
		app := common.NewAppError(500, http.StatusText(500), err)
		return app
	}
	return nil
}
