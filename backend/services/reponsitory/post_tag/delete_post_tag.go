package posttag

import (
	"context"
	"database/sql"
)

// post deleted
func (t *PostTagServiceSQL) DeletePostTagByPostAndTag(ctx context.Context, postId int, tagId int) error {
	query := "DELETE * from post_tags WHERE post_id=@post_id and tag_id=@tag_id"
	_, err := t.db.ExecContext(ctx, query, sql.Named("post_id", postId), sql.Named("post_id", postId), sql.Named("tag_id", tagId))
	if err != nil {
		return err
	}
	return nil
}
