package img

import (
	"context"
	"database/sql"
)

func (t *ImgServiceSQL) CreateImg(ctx context.Context, linkImg string, userId int) error {
	query := "INSERT INTO img(img, user_id) values (@img, @user_id)"
	_, err := t.db.ExecContext(ctx, query, sql.Named("img", linkImg), sql.Named("user_id", userId))
	if err != nil {
		return err
	}
	return nil
}
