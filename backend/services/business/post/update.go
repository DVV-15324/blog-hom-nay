package post

import (
	"bloghomnay-project/common"
	entityPosts "bloghomnay-project/services/entity/posts"
	"context"
	"net/http"
	"strings"
)

func (c *BusinessPost) BusinessUpdatePostByID(ctx context.Context, data *entityPosts.UpdatePost, postId int) *common.AppError {
	// get post
	post, err := c.bzPost.GetPostsByID(ctx, postId)
	if err != nil {
		return common.NewAppError(404, http.StatusText(404), err)
	}

	addTag, removetag := filterTags(post.Tag, data.Tag)

	for _, tag := range addTag {
		if tag.Id == 0 {
			println("⚠️ Bỏ qua tag có ID = 0 (không hợp lệ)")
			continue
		}
		if err := c.bzPostTag.CreatePostTag(ctx, postId, tag.Id); err != nil {
			if strings.Contains(err.Error(), "UNIQUE KEY constraint") {
				println("Tag đã tồn tại:", tag.Id, "postId:", postId)
				continue
			}
			return common.NewAppError(500, http.StatusText(500), err)
		}
	}

	println("Tạo tag xong")

	for _, tag := range removetag {
		if err := c.bzPostTag.DeletePostTagByPostAndTag(ctx, postId, tag.Id); err != nil {
			return common.NewAppError(500, http.StatusText(500), err)
		}
	}

	if err := c.bzPost.UpdatePosts(ctx, data, postId); err != nil {
		return common.NewAppError(500, http.StatusText(500), err)
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
