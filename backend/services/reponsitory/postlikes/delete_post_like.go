package postlikes

import (
	"context"
	"database/sql"
)

// post deleted
func (t *PostLikeServiceSQL) DeletePostLikeByPost(ctx context.Context, userId int, postId int) error {
	query := "DELETE FROM post_likes WHERE post_id = @post_id AND user_id = @user_id"
	_, err := t.db.ExecContext(ctx, query, sql.Named("post_id", postId), sql.Named("user_id", userId))
	if err != nil {
		return err
	}
	return nil
}
