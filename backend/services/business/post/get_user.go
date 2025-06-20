package post

import (
	"bloghomnay-project/common"
	entityPosts "bloghomnay-project/services/entity/posts"
	"context"
	"net/http"
)

func (c *BusinessPost) BusinessGetPostByUserId(ctx context.Context, id int) ([]entityPosts.Posts, *common.AppError) {
<<<<<<< HEAD
	post, err := c.bzPost.GetPostByUserId(ctx, id)
=======
	posts, err := c.bzPost.GetPostByUserId(ctx, id)
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
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

	return posts, nil
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
}
