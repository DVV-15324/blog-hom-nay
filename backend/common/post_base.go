package common

import ()

type PostFormBase struct {
	Id       int           `json:"-"`
	FakeId   string        `json:"id"`
	Title    string        `json:"title"`
	UserBase *UserFormBase `json:"user"`
}

func (t *PostFormBase) Mask() {
	uid := NewUID(uint32(t.Id), 3).ToBase58()
	t.FakeId = uid
}
