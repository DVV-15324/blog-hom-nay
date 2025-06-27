package post

import (
	"bloghomnay-project/common"
	entityPosts "bloghomnay-project/services/entity/posts"
	"context"
	"fmt"
	"net/http"
	"strings"
	"unicode"
)

func (c *BusinessPost) attachTagsAndUsers(ctx context.Context, posts []entityPosts.Posts) {
	for i := range posts {
		if addTag, err := c.bzTag.BusinessGetTagByPostId(ctx, posts[i].Id); err == nil {
			var listTag []common.TagFormBase
			for _, t := range addTag {
				tag := common.TagFormBase{
					Id:   t.Id,
					Name: t.Name,
				}
				tag.Mask()
				listTag = append(listTag, tag)
			}
			posts[i].Tag = listTag
		}
		if addUser, err := c.bzUser.BzGetUsersById(ctx, posts[i].UserID); err == nil {
			user := &common.UserFormBase{
				Id:        addUser.Id,
				FirstName: addUser.FirstName,
				LastName:  addUser.LastName,
				Phone:     addUser.Phone,
				Avatar:    addUser.Avatar,
			}
			user.Mask()
			posts[i].User = user
		}
	}
}

func (c *BusinessPost) BussinessSearchPost(ctx context.Context, search string) ([]entityPosts.Posts, *common.AppError) {
	search = strings.ToLower(search)
	var tags []string
	var builder strings.Builder
	insideTag := false
	for _, r := range search {
		switch {
		case r == '[':
			insideTag = true
			builder.Reset()
		case r == ']':
			if insideTag {
				tags = append(tags, builder.String())
				insideTag = false
				builder.Reset()
			}
		default:
			if insideTag || (!insideTag && !unicode.IsSpace(r)) {
				builder.WriteRune(r)
			}
		}
	}
	keyword := search
	for _, tag := range tags {
		keyword = strings.ReplaceAll(keyword, "["+tag+"]", "")
	}
	keyword = strings.TrimSpace(keyword)
	var (
		post []entityPosts.Posts
		err  error
	)
	switch {
	case len(tags) > 0 && len(keyword) > 0:
		post, err = c.bzPost.SearchsPost(ctx, tags, keyword)
	case len(tags) > 0:
		fmt.Println("🔍 Search tag only")
		post, err = c.bzPost.GetsPostByTagsName(ctx, tags)
	case len(keyword) > 0:
		post, err = c.bzPost.GetPostByTitles(ctx, keyword)
	default:
		return nil, nil
	}
	if err != nil {
		return nil, common.NewAppError(404, http.StatusText(404), err)
	}
	c.attachTagsAndUsers(ctx, post)
	return post, nil
}
