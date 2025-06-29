package auth

import (
	entity "bloghomnay-project/services/entity/auth"
	"context"
	"database/sql"
	//"github.com/microsoft/go-mssqldb"
)

// Kiểm tra Auth tồn tại
func (a *AuthServiceSQL) GetAuthByEmail(ctx context.Context, email string) (*entity.Auth, error) {
	var data entity.Auth
	query := "SELECT email, password, user_id, salt, type_auth FROM auths WHERE email = @email"
	err := a.db.QueryRowContext(ctx, query, sql.Named("email", email)).Scan(
		&data.Email,
		&data.Password,
		&data.UserId,
		&data.Salt,
		&data.TypeAuth,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	return &data, nil
}
