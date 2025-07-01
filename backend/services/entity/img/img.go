package img

import "bloghomnay-project/common"

type Img struct {
	Id     int    `json:"-"`
	FakeId string `json:"id"`
	UserId int    `json:"-"`
	Img    string `json:"img"`
}

func TableName() string {
	return "img"
}

func (i *Img) Mask() {
	uid := common.NewUID(uint32(i.Id), 6)
	i.FakeId = uid.ToBase58()
}
