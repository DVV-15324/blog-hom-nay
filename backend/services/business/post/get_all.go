package post

import (
	"bloghomnay-project/common"
	entityPosts "bloghomnay-project/services/entity/posts"
	"context"
	"net/http"
)

func (c *BusinessPost) BusinessGetALLPost(ctx context.Context) ([]entityPosts.Posts, *common.AppError) {
	post, err := c.bzPost.GetPosts(ctx)
	if err != nil {
		app := common.NewAppError(404, http.StatusText(404), err)
		return nil, app
	}
<<<<<<< HEAD
	for i := 0; i < len(post); i++ {
		//gắn tag
		addTag, _ := c.bzTag.BusinessGetTagByPostId(ctx, post[i].Id)
		listTag := []common.TagFormBase{}
		for i := 0; i < len(addTag); i++ {
			//convert tag
			tag := common.TagFormBase{}
			tag.Id = addTag[i].Id
			tag.Name = addTag[i].Name
			tag.Mask()
			listTag = append(listTag, tag)
		}
		post[i].Tag = listTag
	}
	for i := 0; i < len(post); i++ {
		//gắn user
		addUser, _ := c.bzUser.BzGetUsersById(ctx, post[i].UserID)
		user := common.UserFormBase{}
		user.Id = addUser.Id
		user.FirstName = addUser.FirstName
		user.Avatar = addUser.Avatar
		user.LastName = addUser.LastName
		user.Phone = addUser.Phone

		user.Mask()
		post[i].User = &user

	}
	return post, nil

=======

	return post, nil
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
}
