package posttag

import (
	"context"
	"database/sql"
)

// post deleted
func (t *PostTagServiceSQL) DeletePostTagByPost(ctx context.Context, postId int) error {
	query := "DELETE * from post_tags WHERE post_id=@post_id"
	_, err := t.db.ExecContext(ctx, query, sql.Named("post_id", postId), sql.Named("post_id", postId))
	if err != nil {
		return err
	}
	return nil
}
