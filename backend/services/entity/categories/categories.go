package categories

import "bloghomnay-project/common"

type Categories struct {
	Id          int    `json:"-"`
	FakeId      string `json:"id"`
	Name        string `json:"name"`
	Img         string `json:"img"`
	Description string `json:"description"`
	TagId       int    `json:"-"`
	FakeTagId   string `json:"tag_id"`
}

func (c *Categories) TableName() string {
	return "categories"
}
func (c *Categories) Mask() {
	uid := common.NewUID(uint32(c.Id), 2)
	c.FakeId = uid.ToBase58()
	uid_tags := common.NewUID(uint32(c.TagId), 4)
	c.FakeTagId = uid_tags.ToBase58()
}
