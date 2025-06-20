package common

import ()

type TagFormBase struct {
	Id     int    `json:"-"`
	FakeId string `json:"id"`
	Name   string `json:"name"`
}

func (t *TagFormBase) Mask() {
<<<<<<< HEAD
	uid := NewUID(uint32(t.Id), 4).ToBase58()
=======
	uid := NewUID(uint32(t.Id), 41).ToBase58()
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
	t.FakeId = uid
}
