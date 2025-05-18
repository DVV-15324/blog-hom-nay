package common

import ()

type TagFormBase struct {
	Id     int    `json:"-"`
	FakeId string `json:"id"`
	Name   string `json:"name"`
}

func (t *TagFormBase) Mask() {
	uid := NewUID(uint32(t.Id), 41).ToBase58()
	t.FakeId = uid
}
