package common

import "database/sql"

type UserFormBase struct {
	Id        int            `json:"-"`
	FakeId    string         `json:"user_id"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Phone     string         `json:"phone"`
	Avatar    sql.NullString `json:"avatar"`
}

func (u *UserFormBase) Mask() {
<<<<<<< HEAD
	uid := NewUID(uint32(u.Id), 1).ToBase58()
=======
	uid := NewUID(uint32(u.Id), 11).ToBase58()
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
	u.FakeId = uid
}
