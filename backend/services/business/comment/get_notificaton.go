package comment

import (
	"bloghomnay-project/common"
	entityComment "bloghomnay-project/services/entity/comment"
	"context"
	"net/http"
)

func (c *BusinessComment) BusinessGetNotification(ctx context.Context, userid int, idPost int) ([]entityComment.Comment, *common.AppError) {

	data, err := c.bzComment.GetNotification(ctx, userid)

	if err != nil {
		app := common.NewAppError(404, http.StatusText(404), err)
		return nil, app

	}
	for i := 0; i < len(data); i++ {
		postBase := common.PostFormBase{}
		post, err_post := c.bzPost.BusinessGetPostByID(ctx, userid, idPost)
		if err_post != nil {
			return nil, err_post
		}
		postBase.Title = post.Title
		postBase.Id = post.Id
		postBase.UserBase = post.User
		postBase.Mask()
		data[i].Post = &postBase
		userBase := common.UserFormBase{}
		user, err_user := c.bzUser.BzGetUsersById(ctx, userid)
		if err_user != nil {
			return nil, err_user
		}
		userBase.Avatar = user.Avatar
		userBase.FirstName = user.FirstName
		userBase.LastName = user.LastName
		userBase.Phone = user.Phone
		userBase.Id = user.Id
		userBase.Mask()
		data[i].User = &userBase
	}

	return data, nil
}
