package posts

import (
	"bloghomnay-project/common"
	"time"
)

type Posts struct {
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
}

func (p *Posts) TableName() string {
	return "posts"
}

func (p *Posts) Mask() {
	uid := common.NewUID(uint32(p.Id), 3)
	category_id := common.NewUID(uint32(p.CategoryId), 1)
	user_uid := common.NewUID(uint32(p.UserID), 1)
	p.FakeId = uid.ToBase58()
	p.FakeUserID = user_uid.ToBase58()
	p.FakeCategoryId = category_id.ToBase58()
}
