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
<<<<<<< HEAD
	uid := NewUID(uint32(u.Id), 1).ToBase58()
=======
	uid := NewUID(uint32(u.Id), 11).ToBase58()
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
>>>>>>> 46bb8061e5da0877aec93433ec83d5f5d8b0e033
	u.FakeId = uid
}
