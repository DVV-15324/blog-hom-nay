package categories

import "bloghomnay-project/common"

type Categories struct {
	Id          int    `json:"-"`
	FakeId      string `json:"id"`
	Name        string `json:"name"`
<<<<<<< HEAD
	Img         string `json:"img"`
=======
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
	Description string `json:"description"`
}

func (c *Categories) TableName() string {
	return "categories"
}
func (c *Categories) Mask() {
	uid := common.NewUID(uint32(c.Id), 2)
	c.FakeId = uid.ToBase58()
}
