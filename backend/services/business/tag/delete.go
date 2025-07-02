package tag

import (
	"bloghomnay-project/common"
	"context"
	"net/http"
)

func (c *BusinessTag) BusinessDeleteTag(ctx context.Context, id int) *common.AppError {
	err := c.bzTag.DeleteTagById(ctx, id)
	if err != nil {
		app := common.NewAppError(404, http.StatusText(404), err)
		return app
	}
	er := c.bzPosttag.DeletedPostTagByTag(ctx, id)
	if er != nil {
		app := common.NewAppError(404, http.StatusText(404), er)
		return app
	}
	return nil
}
