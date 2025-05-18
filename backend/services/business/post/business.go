package post

import (
	"bloghomnay-project/common"
	//entityCategories "bloghomnay-project/services/entity/categories"
	entityPost "bloghomnay-project/services/entity/posts"
	entityTag "bloghomnay-project/services/entity/tag"
	entityUser "bloghomnay-project/services/entity/user"
	"context"
)

type ResponsitoryPost interface {
	CreatePost(cxt context.Context, post *entityPost.CreatePost) (int, error)
	DeletePostById(cxt context.Context, id int) error
	GetPostByCategories(ctx context.Context, categoryID int) ([]entityPost.Posts, error)
	GetPostsByID(ctx context.Context, id int) (*entityPost.Posts, error)
	GetsPostByTagsID(ctx context.Context, tagID []int) ([]entityPost.Posts, error)
	GetPostByTitles(ctx context.Context, title string) ([]entityPost.Posts, error)
	GetPosts(ctx context.Context) ([]entityPost.Posts, error)
	GetPostByUserId(ctx context.Context, id int) ([]entityPost.Posts, error)
	UpdatePosts(cxt context.Context, posts *entityPost.UpdatePost, id int) error
}

type BzUser interface {
	BzGetUsersById(ctx context.Context, id int) (*entityUser.Users, *common.AppError)
}

type ResponsitoryPostTag interface {
	CreatePostTag(ctx context.Context, postId int, tagId int) error
	UpdatePostTags(cxt context.Context, postID int, tagId int) error
	DeletePostTagByPostAndTag(ctx context.Context, postId int, tagId int) error
	DeletePostTagByPost(ctx context.Context, postId int) error
}

/*
	type BzCategories interface {
		GetcategoriesByID(ctx context.Context, id int) (*entityCategories.Categories, error)
	}
*/
type BzTag interface {
	BusinessGetTagByPostId(ctx context.Context, postID int) ([]entityTag.Tag, *common.AppError)
}
type BusinessPost struct {
	bzPost    ResponsitoryPost
	bzUser    BzUser
	bzPostTag ResponsitoryPostTag
	//bzCategories BzCategories
	bzTag BzTag
}

func NewBusinessPost(bzPost ResponsitoryPost, bzUser BzUser, bzPostTag ResponsitoryPostTag, bzTag BzTag) *BusinessPost {
	return &BusinessPost{
		bzPost:    bzPost,
		bzUser:    bzUser,
		bzPostTag: bzPostTag,
		bzTag:     bzTag,
	}
}
