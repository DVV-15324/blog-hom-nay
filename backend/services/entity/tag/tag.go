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
<<<<<<< HEAD
	return "tags"
=======
	return "tag"
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
}
func (t *Tag) Mask() {
	uid := common.NewUID(uint32(t.Id), 4)
	t.Fake = uid.ToBase58()
}
