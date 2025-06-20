package comment

import (
	"bloghomnay-project/common"
	entityComment "bloghomnay-project/services/entity/comment"
	"context"
)

type BusinessComment interface {
<<<<<<< HEAD
	BusinessCreateComment(ctx context.Context, data *entityComment.CreateComment) (*entityComment.Comment, *common.AppError)
	BusinessGetComment(ctx context.Context, postID int) ([]entityComment.Comment, *common.AppError)
	BusinessGetNotification(ctx context.Context, userid int, idPost int) ([]entityComment.Comment, *common.AppError)
=======
	BusinessCreateComment(ctx context.Context, data *entityComment.CreateComment) *common.AppError
	BusinessGetComment(ctx context.Context, postID int) ([]entityComment.Comment, *common.AppError)
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
}
type ApiComment struct {
	bz BusinessComment
}

func NewApiComment(bz BusinessComment) *ApiComment {
	return &ApiComment{
		bz: bz,
	}
}
