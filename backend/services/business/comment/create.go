package comment

import (
	"bloghomnay-project/common"
	entityComment "bloghomnay-project/services/entity/comment"
	"context"
	"net/http"
)

<<<<<<< HEAD
func (c *BusinessComment) BusinessCreateComment(ctx context.Context, data *entityComment.CreateComment) (*entityComment.Comment, *common.AppError) {
=======
<<<<<<< HEAD
func (c *BusinessComment) BusinessCreateComment(ctx context.Context, data *entityComment.CreateComment) (*entityComment.Comment, *common.AppError) {
	er := data.Validate()
	if er != nil {
		app := common.NewAppError(400, http.StatusText(400), er)
		return nil, app
	}
	comment, err := c.bzComment.CreateComment(ctx, data)
	if err != nil {
		app := common.NewAppError(500, http.StatusText(500), err)
		return nil, app
	}

	var user common.UserFormBase
	u, _ := c.bzUser.BzGetUsersById(ctx, data.UserID)
	user.Id = u.Id
	user.Avatar = u.Avatar
	user.FirstName = u.FirstName
	user.LastName = u.LastName
	user.Phone = u.Phone
	user.Mask()
	comment.User = &user

	return comment, nil
=======
func (c *BusinessComment) BusinessCreateComment(ctx context.Context, data *entityComment.CreateComment) *common.AppError {
>>>>>>> 46bb8061e5da0877aec93433ec83d5f5d8b0e033
	er := data.Validate()
	if er != nil {
		app := common.NewAppError(400, http.StatusText(400), er)
		return nil, app
	}
	comment, err := c.bzComment.CreateComment(ctx, data)
	if err != nil {
		app := common.NewAppError(500, http.StatusText(500), err)
		return nil, app
	}
<<<<<<< HEAD

	var user common.UserFormBase
	u, _ := c.bzUser.BzGetUsersById(ctx, data.UserID)
	user.Id = u.Id
	user.Avatar = u.Avatar
	user.FirstName = u.FirstName
	user.LastName = u.LastName
	user.Phone = u.Phone
	user.Mask()
	comment.User = &user

	return comment, nil
=======
	return nil
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
>>>>>>> 46bb8061e5da0877aec93433ec83d5f5d8b0e033
}
