package comment

import (
	"bloghomnay-project/common"
	"time"
)

type Comment struct {
	Id         int                  `json:"-"`
	FakeId     string               `json:"id"`
	PostId     int                  `json:"-"`
	FakePostId string               `json:"post_id"`
	UserId     int                  `json:"-"`
	FakeUserId string               `json:"user_id"`
	Content    string               `json:"content"`
	User       *common.UserFormBase `json:"user"`
	Post       *common.PostFormBase `json:"post,omitempty"`
	CreatedAt  time.Time            `json:"created_at"`
}

func (c *Comment) TableName() string {
	return "comments"
}

func (c *Comment) Mask() {
	uid := common.NewUID(uint32(c.Id), 5)
	c.FakeId = uid.ToBase58()
	uid_post := common.NewUID(uint32(c.PostId), 3)
	c.FakePostId = uid_post.ToBase58()
	uid_user := common.NewUID(uint32(c.UserId), 1)
	c.FakeUserId = uid_user.ToBase58()
}
