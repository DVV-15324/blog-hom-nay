package post

import (
	"bloghomnay-project/common"
	entityPosts "bloghomnay-project/services/entity/posts"
	"context"
)

type BusinessPosts interface {
	BusinessCreatePost(ctx context.Context, data *entityPosts.CreatePost) *common.AppError
	BusinessDeletePost(ctx context.Context, id int) *common.AppError
	BusinessGetALLPost(ctx context.Context) ([]entityPosts.Posts, *common.AppError)
	BusinessGetPostByCategories(ctx context.Context, categoriesID int) ([]entityPosts.Posts, *common.AppError)
	BusinessGetPostByID(ctx context.Context, id int) (*entityPosts.Posts, *common.AppError)
	BusinessGetPostByTag(ctx context.Context, tagId []int) ([]entityPosts.Posts, *common.AppError)
	BusinessGetPostByTitles(ctx context.Context, title string) ([]entityPosts.Posts, *common.AppError)
	BusinessGetPostByUserId(ctx context.Context, id int) ([]entityPosts.Posts, *common.AppError)
	BusinessUpdatePostByID(ctx context.Context, data *entityPosts.UpdatePost, postId int) *common.AppError
}
type ApiPosts struct {
	bz BusinessPosts
}

func NewApiPosts(bz BusinessPosts) *ApiPosts {
	return &ApiPosts{
		bz: bz,
	}
}
