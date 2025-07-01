package tag

import (
	"bloghomnay-project/common"
	entityTag "bloghomnay-project/services/entity/tag"
	"context"
	"net/http"
)

func (c *BusinessTag) BusinessGetTagByPostId(ctx context.Context, postID int) ([]entityTag.Tag, *common.AppError) {
	tag, err := c.bzTag.GetTagByPostId(ctx, postID)
	if err != nil {
		app := common.NewAppError(404, http.StatusText(404), err)
		return nil, app
	}

	return tag, nil
}
