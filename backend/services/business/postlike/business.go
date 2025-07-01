package postlike

import (
	"context"
)

type ResponsitoryPostLike interface {
	CreatePostLikes(ctx context.Context, userId int, postId int) error
	DeletePostLikeByPost(ctx context.Context, userId int, postId int) error
	GetPostLike(ctx context.Context, userID int, postId int) (bool, error)
}

type BusinessPostLike struct {
	bzPostLike ResponsitoryPostLike
}

func NewBusinessPostLike(bzPostLike ResponsitoryPostLike) *BusinessPostLike {
	return &BusinessPostLike{
		bzPostLike: bzPostLike,
	}
}
