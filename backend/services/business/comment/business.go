package comment

import (
	"bloghomnay-project/common"
	entityComment "bloghomnay-project/services/entity/comment"
	entityUser "bloghomnay-project/services/entity/user"
	"context"
)

type ReponsitoryComment interface {
	CreateComment(cxt context.Context, comment *entityComment.CreateComment) error
	GetCommentByPostId(ctx context.Context, postID int) ([]entityComment.Comment, error)
}

type BzUser interface {
	BzGetUsersById(ctx context.Context, id int) (*entityUser.Users, *common.AppError)
}

type BusinessComment struct {
	bzComment ReponsitoryComment
	bzUser    BzUser
}

func NewBusinessComment(bzC ReponsitoryComment, bzU BzUser) *BusinessComment {
	return &BusinessComment{
		bzComment: bzC,
		bzUser:    bzU,
	}
}
