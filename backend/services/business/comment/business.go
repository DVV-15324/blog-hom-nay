package comment

import (
	"bloghomnay-project/common"
	entityComment "bloghomnay-project/services/entity/comment"
	entityPost "bloghomnay-project/services/entity/posts"
	entityUser "bloghomnay-project/services/entity/user"
	"context"
)

type ReponsitoryComment interface {
	CreateComment(cxt context.Context, comment *entityComment.CreateComment) (*entityComment.Comment, error)
	GetCommentByPostId(ctx context.Context, postID int) ([]entityComment.Comment, error)
	GetNotification(ctx context.Context, userID int) ([]entityComment.Comment, error)
}

type BzUser interface {
	BzGetUsersById(ctx context.Context, id int) (*entityUser.Users, *common.AppError)
}

type BzPost interface {
	BusinessGetPostByID(ctx context.Context, userId int, postId int) (*entityPost.Posts, *common.AppError)
}

type BusinessComment struct {
	bzComment ReponsitoryComment
	bzUser    BzUser
	bzPost    BzPost
}

func NewBusinessComment(bzC ReponsitoryComment, bzU BzUser, bzPost BzPost) *BusinessComment {
	return &BusinessComment{
		bzComment: bzC,
		bzUser:    bzU,
		bzPost:    bzPost,
	}
}
