package common

import "time"

type CommentFormBase struct {
	Id         int           `json:"-"`
	FakeId     string        `json:"id"`
	PostId     int           `json:"-"`
	FakePostId string        `json:"post_id"`
	UserId     int           `json:"-"`
	FakeUserId string        `json:"user_id"`
	Content    string        `json:"content"`
	User       *UserFormBase `json:"user"`
	CreatedAt  time.Time     `json:"created_at"`
}

func (c *CommentFormBase) Mask() {
	uid := NewUID(uint32(c.Id), 5)
	c.FakeId = uid.ToBase58()
	uid_post := NewUID(uint32(c.PostId), 3)
	c.FakePostId = uid_post.ToBase58()
	uid_user := NewUID(uint32(c.UserId), 1)
	c.FakeUserId = uid_user.ToBase58()
}
