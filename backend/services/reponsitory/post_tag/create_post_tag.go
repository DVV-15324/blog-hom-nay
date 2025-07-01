package posttag

import (
	"context"
	"database/sql"
)

func (t *PostTagServiceSQL) CreatePostTag(ctx context.Context, postId int, tagId int) error {
	query := "INSERT INTO post_tags(post_id, tag_id) values (@post_id, @tag_id)"
	_, err := t.db.ExecContext(ctx, query, sql.Named("post_id", postId), sql.Named("tag_id", tagId))
	if err != nil {
		return err
	}
	return nil
}
