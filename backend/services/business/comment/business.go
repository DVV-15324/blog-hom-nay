package comment

import (
	"bloghomnay-project/common"
	entityComment "bloghomnay-project/services/entity/comment"
<<<<<<< HEAD
	entityPost "bloghomnay-project/services/entity/posts"
=======
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
	entityUser "bloghomnay-project/services/entity/user"
	"context"
)

type ReponsitoryComment interface {
<<<<<<< HEAD
	CreateComment(cxt context.Context, comment *entityComment.CreateComment) (*entityComment.Comment, error)
	GetCommentByPostId(ctx context.Context, postID int) ([]entityComment.Comment, error)
	GetNotification(ctx context.Context, userID int) ([]entityComment.Comment, error)
=======
	CreateComment(cxt context.Context, comment *entityComment.CreateComment) error
	GetCommentByPostId(ctx context.Context, postID int) ([]entityComment.Comment, error)
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
}

type BzUser interface {
	BzGetUsersById(ctx context.Context, id int) (*entityUser.Users, *common.AppError)
}

<<<<<<< HEAD
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
=======
type BusinessComment struct {
	bzComment ReponsitoryComment
	bzUser    BzUser
}

func NewBusinessComment(bzC ReponsitoryComment, bzU BzUser) *BusinessComment {
	return &BusinessComment{
		bzComment: bzC,
		bzUser:    bzU,
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
	}
}
