package post

import (
	"bloghomnay-project/common"
<<<<<<< HEAD
	entityComment "bloghomnay-project/services/entity/comment"
=======
	//entityCategories "bloghomnay-project/services/entity/categories"
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
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
	//GetsPostByTagsID(ctx context.Context, tagID []int) ([]entityPost.Posts, error)
	GetPostByTitles(ctx context.Context, title string) ([]entityPost.Posts, error)
	GetsPostByTagsName(ctx context.Context, tagName []string) ([]entityPost.Posts, error)
	GetPosts(ctx context.Context) ([]entityPost.Posts, error)
	GetPostByUserId(ctx context.Context, id int) ([]entityPost.Posts, error)
	UpdatePosts(cxt context.Context, posts *entityPost.UpdatePost, id int) error
	SearchsPost(ctx context.Context, tagsName []string, title string) ([]entityPost.Posts, error)
}

type BzUser interface {
	BzGetUsersById(ctx context.Context, id int) (*entityUser.Users, *common.AppError)
}

type ResponsitoryPostTag interface {
	CreatePostTag(ctx context.Context, postId int, tagId int) error
	DeletePostTagByPostAndTag(ctx context.Context, postId int, tagId int) error
	DeletePostTagByPost(ctx context.Context, postId int) error
}

type BzComment interface {
	BusinessGetComment(ctx context.Context, postID int) ([]entityComment.Comment, *common.AppError)
}

type BzTag interface {
	BusinessGetTagByPostId(ctx context.Context, postID int) ([]entityTag.Tag, *common.AppError)
}

type BzPostLike interface {
	BusinessGetPostLike(ctx context.Context, userId int, postID int) (bool, *common.AppError)
}
type BusinessPost struct {
	bzPost     ResponsitoryPost
	bzUser     BzUser
	bzPostTag  ResponsitoryPostTag
	bzComment  BzComment
	bzTag      BzTag
	bzPostLike BzPostLike
}

func NewBusinessPost(bzPost ResponsitoryPost, bzUser BzUser, bzPostTag ResponsitoryPostTag, bzTag BzTag, bzPostLike BzPostLike) *BusinessPost {
	return &BusinessPost{
		bzPost:     bzPost,
		bzUser:     bzUser,
		bzPostTag:  bzPostTag,
		bzTag:      bzTag,
		bzComment:  nil,
		bzPostLike: bzPostLike,
	}
}
func (bz *BusinessPost) AddComment(bzComment BzComment) {
	bz.bzComment = bzComment
}
