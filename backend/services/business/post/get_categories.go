package post

import (
	"bloghomnay-project/common"
	entityPosts "bloghomnay-project/services/entity/posts"
	"context"
	"net/http"
)

func (c *BusinessPost) BusinessGetPostByCategories(ctx context.Context, categoriesID int) ([]entityPosts.Posts, *common.AppError) {
	posts, err := c.bzPost.GetPostByCategories(ctx, categoriesID)
	if err != nil {
		app := common.NewAppError(404, http.StatusText(404), err)
		return nil, app
	}

	return posts, nil
}
