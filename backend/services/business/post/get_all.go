package post

import (
	"bloghomnay-project/common"
	entityPosts "bloghomnay-project/services/entity/posts"
	"context"
	"net/http"
)

func (c *BusinessPost) BusinessGetALLPost(ctx context.Context) ([]entityPosts.Posts, *common.AppError) {
	post, err := c.bzPost.GetPosts(ctx)
	if err != nil {
		app := common.NewAppError(404, http.StatusText(404), err)
		return nil, app
	}

	return post, nil
}
