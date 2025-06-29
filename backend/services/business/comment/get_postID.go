package comment

import (
	"bloghomnay-project/common"
	entityComment "bloghomnay-project/services/entity/comment"
	"context"
<<<<<<< HEAD

=======
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
	"net/http"
)

func (c *BusinessComment) BusinessGetComment(ctx context.Context, postID int) ([]entityComment.Comment, *common.AppError) {
<<<<<<< HEAD

	comment, err := c.bzComment.GetCommentByPostId(ctx, postID)

=======
<<<<<<< HEAD

	comment, err := c.bzComment.GetCommentByPostId(ctx, postID)

=======
	comment, err := c.bzComment.GetCommentByPostId(ctx, postID)
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
>>>>>>> 46bb8061e5da0877aec93433ec83d5f5d8b0e033
	if err != nil {
		app := common.NewAppError(404, http.StatusText(404), err)
		return nil, app
	}
	for i := 0; i < len(comment); i++ {
<<<<<<< HEAD
=======
<<<<<<< HEAD
>>>>>>> 46bb8061e5da0877aec93433ec83d5f5d8b0e033
		var user common.UserFormBase
		u, _ := c.bzUser.BzGetUsersById(ctx, comment[i].UserId)
		user.Id = u.Id
		user.Avatar = u.Avatar
		user.FirstName = u.FirstName
		user.LastName = u.LastName
		user.Phone = u.Phone
		user.Mask()
		comment[i].User = &user

<<<<<<< HEAD
=======
=======
		user, _ := c.bzUser.BzGetUsersById(ctx, comment[i].UserId)
		comment[i].User.Avatar = user.Avatar
		comment[i].User.FirstName = user.FirstName
		comment[i].User.LastName = user.LastName
		comment[i].User.Phone = user.Phone
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
>>>>>>> 46bb8061e5da0877aec93433ec83d5f5d8b0e033
	}
	return comment, nil
}
