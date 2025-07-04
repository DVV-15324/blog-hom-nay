package posts

import (
	"bloghomnay-project/common"
	"strings"
	"time"
)

type Posts struct {
	Id             int                      `json:"-"`
	FakeId         string                   `json:"id"`
	UserID         int                      `json:"-"`
	FakeUserID     string                   `json:"user_id"`
	CategoryId     int                      `json:"-"`
	FakeCategoryId string                   `json:"category_id"`
	Title          string                   `json:"title"`
	Content        string                   `json:"content"`
	Tag            []common.TagFormBase     `json:"tags"`
	User           *common.UserFormBase     `json:"user"`
	Description    string                   `json:"description"`
	Like           int                      `json:"like"`
	IsLike         bool                     `json:"islike"`
	CountComment   int                      `json:"count_comment,omitempty"`
	Comments       []common.CommentFormBase `json:"comments"`
	Slug           string                   `json:"slug"`
	UpdatedAt      time.Time                `json:"updated_at"`
	CreatedAt      time.Time                `json:"created_at"`
	DeletedAt      time.Time                `json:"deleted_at"`
}

func (p *Posts) TableName() string {
	return "posts"
}

func (p *Posts) Mask() {
	uid := common.NewUID(uint32(p.Id), 3)
	p.Content = strings.ReplaceAll(p.Content, `\"`, `"`)
	category_id := common.NewUID(uint32(p.CategoryId), 2)
	user_uid := common.NewUID(uint32(p.UserID), 1)
	p.FakeId = uid.ToBase58()
	p.FakeUserID = user_uid.ToBase58()
	p.FakeCategoryId = category_id.ToBase58()
}
