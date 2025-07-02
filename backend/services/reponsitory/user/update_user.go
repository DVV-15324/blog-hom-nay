package user

import (
	entity "bloghomnay-project/services/entity/user"
	"context"
	"database/sql"
	"fmt"
	"strings"
)

func (u *UserServiceSQL) UpdateUser(cxt context.Context, user *entity.UpdateUserForm, id int) error {
	if user == nil {
		return fmt.Errorf("user is nil")
	}

	var (
		placeholders []string
		args         []interface{}
	)

	if user.FirstName != nil {
		placeholders = append(placeholders, "first_name = @first_name")
		args = append(args, sql.Named("first_name", *user.FirstName))
	}
	if user.LastName != nil {
		placeholders = append(placeholders, "last_name = @last_name")
		args = append(args, sql.Named("last_name", *user.LastName))
	}
	if user.Phone != nil {
		placeholders = append(placeholders, "phone = @phone")
		args = append(args, sql.Named("phone", *user.Phone))
	}
	if user.Address != nil {
		placeholders = append(placeholders, "address = @address")
		args = append(args, sql.Named("address", *user.Address))
	}
	if user.Avatar != nil {
		placeholders = append(placeholders, "avatar = @avatar")
		args = append(args, sql.Named("avatar", *user.Avatar))
	}

	if len(placeholders) == 0 {
		return fmt.Errorf("no fields to update")
	}

	query := fmt.Sprintf("UPDATE users SET %s WHERE id = @id", strings.Join(placeholders, ", "))

	args = append(args, sql.Named("id", id))

	_, err := u.db.ExecContext(cxt, query, args...)
	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}

	return nil
}
