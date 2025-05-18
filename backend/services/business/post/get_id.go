package post

import (
	"bloghomnay-project/common"
	entityPosts "bloghomnay-project/services/entity/posts"
	"context"
	"fmt"
	"net/http"
)

func (c *BusinessPost) BusinessGetPostByID(ctx context.Context, id int) (*entityPosts.Posts, *common.AppError) {
	post, err := c.bzPost.GetPostsByID(ctx, id)
	if err != nil {
		app := common.NewAppError(404, http.StatusText(404), err)
		return nil, app
	}
	//gắn tag
	addTag, _ := c.bzTag.BusinessGetTagByPostId(ctx, post.Id)
	listTag := []common.TagFormBase{}
	for i := 0; i < len(addTag); i++ {
		//convert tag
		tag := common.TagFormBase{}
		tag.Id = addTag[i].Id
		tag.Name = addTag[i].Name
		tag.Mask()
		listTag = append(listTag, tag)
	}
	fmt.Println(listTag)
	//gắn user
	addUser, _ := c.bzUser.BzGetUsersById(ctx, post.UserID)
	user := common.UserFormBase{}
	user.Id = addUser.Id
	user.FirstName = addUser.FirstName
	user.Avatar = addUser.Avatar
	user.LastName = addUser.LastName
	user.Phone = addUser.Phone
	user.Mask()
	post.User = &user

	post.Tag = listTag
	return post, nil
}
