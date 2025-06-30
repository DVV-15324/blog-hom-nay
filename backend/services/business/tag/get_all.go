package tag

import (
	"bloghomnay-project/common"
	entityTag "bloghomnay-project/services/entity/tag"
	"context"
	"net/http"
)

func (c *BusinessTag) BusinessGetALLTag(ctx context.Context) ([]entityTag.Tag, *common.AppError) {
	tag, err := c.bzTag.GetALLTags(ctx)
	if err != nil {
		app := common.NewAppError(404, http.StatusText(404), err)
		return nil, app
	}

	return tag, nil
}
