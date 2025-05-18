package post

import (
	"bloghomnay-project/common"
	entityPosts "bloghomnay-project/services/entity/posts"
	"context"
	"net/http"
)

func (c *BusinessPost) BusinessGetPostByTitles(ctx context.Context, title string) ([]entityPosts.Posts, *common.AppError) {
	posts, err := c.bzPost.GetPostByTitles(ctx, title)
	if err != nil {
		app := common.NewAppError(404, http.StatusText(404), err)
		return nil, app
	}

	return posts, nil
}
