package post

import (
	"bloghomnay-project/common"
	entityPosts "bloghomnay-project/services/entity/posts"
	"context"
	"net/http"
)

func (c *BusinessPost) BusinessUpdatePostByID(ctx context.Context, data *entityPosts.UpdatePost, postId int) *common.AppError {
	//get post
	post, err := c.bzPost.GetPostsByID(ctx, postId)
	if err != nil {
		app := common.NewAppError(404, http.StatusText(404), err)
		return app
	}
	addTag, removetag := filterTags(post.Tag, data.Tag)
	for _, tag := range addTag {
		if err := c.bzPostTag.UpdatePostTags(ctx, postId, tag.Id); err != nil {
			return common.NewAppError(500, http.StatusText(500), err)
		}
	}
	for _, tag := range removetag {
		if err := c.bzPostTag.DeletePostTagByPostAndTag(ctx, postId, tag.Id); err != nil {
			return common.NewAppError(500, http.StatusText(500), err)
		}
	}

	err_u := c.bzPost.UpdatePosts(ctx, data, postId)
	if err_u != nil {
		app := common.NewAppError(500, http.StatusText(500), err)
		return app
	}
	return nil
}
func filterTags(oldTags, newTags []common.TagFormBase) (toAdd, toRemove []common.TagFormBase) {
	oldMap := make(map[int]common.TagFormBase)
	newMap := make(map[int]common.TagFormBase)

	for _, t := range oldTags {
		oldMap[t.Id] = t
	}
	for _, t := range newTags {
		newMap[t.Id] = t
	}

	// Tìm tag cần thêm (có trong new mà không có trong old)
	for id, tag := range newMap {
		if _, exists := oldMap[id]; !exists {
			toAdd = append(toAdd, tag)
		}
	}

	// Tìm tag cần xóa (có trong old mà không có trong new)
	for id, tag := range oldMap {
		if _, exists := newMap[id]; !exists {
			toRemove = append(toRemove, tag)
		}
	}

	return
}
