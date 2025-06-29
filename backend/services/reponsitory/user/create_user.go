package user

import (
	entity "bloghomnay-project/services/entity/user"
	"context"
	"database/sql"
)

// Tạo User
func (u *UserServiceSQL) CreateUser(cxt context.Context, user *entity.CreateUserForm) (int, error) {
	query := `INSERT INTO users(first_name, last_name, email, phone)
			OUTPUT INSERTED.id
			values(@first_name, @last_name, @email, @phone);`
	var uid int
	err := u.db.QueryRowContext(cxt, query,
		sql.Named("first_name", user.FirstName),
		sql.Named("last_name", user.LastName),
		sql.Named("email", user.Email),
		sql.Named("phone", user.Phone),
	).Scan(&uid)
	if err != nil {
		return 0, err
	}
	return uid, nil
}
