package postlikes

import (
	"context"
	"database/sql"
)

func (t *PostLikeServiceSQL) CreatePostLikes(ctx context.Context, userId int, postId int) error {
	query := "INSERT INTO post_likes(user_id, post_id) values (@user_id, @post_id)"
	_, err := t.db.ExecContext(ctx, query, sql.Named("user_id", userId), sql.Named("post_id", postId))
	if err != nil {
		return err
	}
	return nil
}
