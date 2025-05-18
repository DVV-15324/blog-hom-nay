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
		user, _ := c.bzUser.BzGetUsersById(ctx, comment[i].UserId)
		comment[i].User.Avatar = user.Avatar
		comment[i].User.FirstName = user.FirstName
		comment[i].User.LastName = user.LastName
		comment[i].User.Phone = user.Phone
	}
	return comment, nil
}
