package comment

import (
	"bloghomnay-project/common"
	entityComment "bloghomnay-project/services/entity/comment"
	"context"
)

type BusinessComment interface {
	BusinessCreateComment(ctx context.Context, data *entityComment.CreateComment) *common.AppError
	BusinessGetComment(ctx context.Context, postID int) ([]entityComment.Comment, *common.AppError)
}
type ApiComment struct {
	bz BusinessComment
}

func NewApiComment(bz BusinessComment) *ApiComment {
	return &ApiComment{
		bz: bz,
	}
}
