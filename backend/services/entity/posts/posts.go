package posts

import (
	"bloghomnay-project/common"
<<<<<<< HEAD
	"strings"
=======
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
	"time"
)

type Posts struct {
<<<<<<< HEAD
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
	UpdatedAt      time.Time                `json:"updated_at"`
	CreatedAt      time.Time                `json:"created_at"`
	DeletedAt      time.Time                `json:"deleted_at"`
=======
	Id             int                  `json:"-"`
	FakeId         string               `json:"id"`
	UserID         int                  `json:"-"`
	FakeUserID     string               `json:"user_id"`
	CategoryId     int                  `json:"-"`
	FakeCategoryId string               `json:"category_id"`
	Title          string               `json:"title"`
	Content        string               `json:"content"`
	Tag            []common.TagFormBase `json:"tags"`
	User           *common.UserFormBase `json:"user"`
	UpdatedAt      time.Time            `josn:"updated_at"`
	CreatedAt      time.Time            `json:"created_at"`
	DeletedAt      time.Time            `json:"deleted_at"`
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
}

func (p *Posts) TableName() string {
	return "posts"
}

func (p *Posts) Mask() {
	uid := common.NewUID(uint32(p.Id), 3)
<<<<<<< HEAD
	p.Content = strings.ReplaceAll(p.Content, `\"`, `"`)
	category_id := common.NewUID(uint32(p.CategoryId), 2)
=======
	category_id := common.NewUID(uint32(p.CategoryId), 1)
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
	user_uid := common.NewUID(uint32(p.UserID), 1)
	p.FakeId = uid.ToBase58()
	p.FakeUserID = user_uid.ToBase58()
	p.FakeCategoryId = category_id.ToBase58()
}
