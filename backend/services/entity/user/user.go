package user

import (
	"bloghomnay-project/common"
	"database/sql"
)

type Users struct {
	Id        int            `json:"-"`
	FakeId    string         `json:"id"`
	FirstName string         `json:"first_name"`
	LastName  string         `json:"last_name"`
	Phone     string         `json:"phone"`
	Avatar    sql.NullString `json:"avatar"`
	Address   sql.NullString `json:"address"`
	Email     string         `json:"email"`
	DeletedAt string         `json:"deleted_at"`
	CreatedAt string         `json:"created_at"`
	UpdatedAt string         `json:"updated_at"`
}

func (u *Users) Table() string {
	return "users"
}

func (u *Users) Mask() {
	uid := common.NewUID(uint32(u.Id), 1).ToBase58()
	u.FakeId = uid
}
