package post

<<<<<<< HEAD
/*
=======
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
import (
	"bloghomnay-project/common"
	entityPosts "bloghomnay-project/services/entity/posts"
	"context"
	"net/http"
)

<<<<<<< HEAD
func (c *BusinessPost) BusinessGetPostByTagName(ctx context.Context, tagId []int) ([]entityPosts.Posts, *common.AppError) {
=======
func (c *BusinessPost) BusinessGetPostByTag(ctx context.Context, tagId []int) ([]entityPosts.Posts, *common.AppError) {
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
	posts, err := c.bzPost.GetsPostByTagsID(ctx, tagId)
	if err != nil {
		app := common.NewAppError(404, http.StatusText(404), err)
		return nil, app
	}

	return posts, nil
}
<<<<<<< HEAD
*/
=======
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
