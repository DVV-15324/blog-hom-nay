package post

import (
	"bloghomnay-project/common"
	entityPosts "bloghomnay-project/services/entity/posts"
	"context"

	"net/http"
)

func (c *BusinessPost) BusinessGetPostByID(ctx context.Context, userId int, postId int) (*entityPosts.Posts, *common.AppError) {
	post, err := c.bzPost.GetPostsByID(ctx, postId)
	if err != nil {
		app := common.NewAppError(404, http.StatusText(404), err)
		return nil, app
	}
	//checkLike
	isLike, err_postLike := c.bzPostLike.BusinessGetPostLike(ctx, userId, postId)
	if err_postLike != nil {
		return nil, err_postLike
	}
	post.IsLike = isLike
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
	addComment, _ := c.bzComment.BusinessGetComment(ctx, postId)

	if addComment != nil {
		listComment := []common.CommentFormBase{}
		for i := 0; i < len(addComment); i++ {
			//convert tag
			comment := common.CommentFormBase{}
			comment.Id = addComment[i].Id
			comment.Content = addComment[i].Content
			comment.PostId = addComment[i].PostId
			comment.UserId = addComment[i].UserId
			comment.User = addComment[i].User
			comment.CreatedAt = addComment[i].CreatedAt
			comment.Mask()
			listComment = append(listComment, comment)
		}

		post.Comments = listComment
	}

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
