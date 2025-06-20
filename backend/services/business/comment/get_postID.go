package comment

import (
	"bloghomnay-project/common"
	entityComment "bloghomnay-project/services/entity/comment"
	"context"

	"net/http"
)

func (c *BusinessComment) BusinessGetComment(ctx context.Context, postID int) ([]entityComment.Comment, *common.AppError) {

	comment, err := c.bzComment.GetCommentByPostId(ctx, postID)

	if err != nil {
		app := common.NewAppError(404, http.StatusText(404), err)
		return nil, app
	}
	for i := 0; i < len(comment); i++ {
		var user common.UserFormBase
		u, _ := c.bzUser.BzGetUsersById(ctx, comment[i].UserId)
		user.Id = u.Id
		user.Avatar = u.Avatar
		user.FirstName = u.FirstName
		user.LastName = u.LastName
		user.Phone = u.Phone
		user.Mask()
		comment[i].User = &user

	}
	return comment, nil
}
