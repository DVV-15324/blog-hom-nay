package user

import (
	"context"
	"database/sql"
)

func (u *UserServiceSQL) DeleteUserById(cxt context.Context, id int) error {
	query := "DELETE users Where id = @id"
	_, err := u.db.ExecContext(cxt, query, sql.Named("id", id))
	if err != nil {
		return err
	}
	return nil
}
