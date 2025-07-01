package postlikes

import ()

type PostLike struct {
	Id        int    `json:"-"`
	UserId    int    `json:"-"`
	PostId    int    `json:"-"`
	CreatedAt string `json:"-"`
}

func TableName() string {
	return "post_likes"
}
