package postlike

import (
	"bloghomnay-project/common"
	"context"

	"net/http"
)

func (c *BusinessPostLike) BusinessGetPostLike(ctx context.Context, userId int, postID int) (bool, *common.AppError) {

	isLike, err := c.bzPostLike.GetPostLike(ctx, userId, postID)
	if err != nil {
		app := common.NewAppError(404, http.StatusText(404), err)
		return false, app
	}
	return isLike, nil
}
