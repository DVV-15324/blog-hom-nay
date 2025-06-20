package postlikes

type CreatePostLikes struct {
	UserId int `json:"-"`
	PostId int `json:"-"`
}
