package auth

import (
	entity "bloghomnay-project/services/entity/auth"
	"context"
	"database/sql"
)

// tạo auth
func (u *AuthServiceSQL) CreateAuth(cxt context.Context, auth *entity.Auth) error {
	query := `INSERT INTO auths(salt, email, password, user_id, type_auth)
			OUTPUT INSERTED.id
			values(@salt, @email, @password, @user_id, @type_auth);`
	_, err := u.db.ExecContext(cxt, query,
		sql.Named("salt", auth.Salt),
		sql.Named("email", auth.Email),
		sql.Named("password", auth.Password),
		sql.Named("user_id", auth.UserId),
		sql.Named("type_auth", auth.TypeAuth),
	)
	if err != nil {
		return err
	}
	return nil
}
