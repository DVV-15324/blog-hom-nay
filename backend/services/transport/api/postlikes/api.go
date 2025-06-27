package postlikes

import (
	"bloghomnay-project/common"
	entityPostLike "bloghomnay-project/services/entity/postlikes"
	"context"
)

type BusinessPostLike interface {
	BusinessCreatePostLike(ctx context.Context, data *entityPostLike.CreatePostLikes) *common.AppError
	BusinessDeletePostLike(ctx context.Context, userId int, postId int) *common.AppError
}
type ApiPostLike struct {
	bz BusinessPostLike
}

func NewApiPostLikes(bz BusinessPostLike) *ApiPostLike {
	return &ApiPostLike{
		bz: bz,
	}
}
