package tag

import (
	"bloghomnay-project/common"
)

type Tag struct {
	Id   int    `json:"-"`
	Fake string `json:"id"`
	Name string `json:"name"`
}

func (t *Tag) TableName() string {
	return "tag"
}
func (t *Tag) Mask() {
	uid := common.NewUID(uint32(t.Id), 4)
	t.Fake = uid.ToBase58()
}
