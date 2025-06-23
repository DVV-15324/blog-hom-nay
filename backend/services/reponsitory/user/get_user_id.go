package user

import (
	entity "bloghomnay-project/services/entity/user"
	"context"
	"database/sql"
)

// thông tin user
func (u *UserServiceSQL) GetUserById(ctx context.Context, id int) (*entity.Users, error) {
	var data entity.Users
	query := "SELECT id, first_name, last_name, phone, avatar, address, email from users where id=@id"
	row := u.db.QueryRowContext(ctx, query, sql.Named("id", id))
	err := row.Scan(&data.Id, &data.FirstName, &data.LastName, &data.Phone, &data.Avatar, &data.Address, &data.Email)
	if err != nil {
		return nil, err
	}
	return &data, nil
}
