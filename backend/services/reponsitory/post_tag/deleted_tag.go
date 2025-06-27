package posttag

import (
	"context"
	"database/sql"
)

// post deleted
func (t *PostTagServiceSQL) DeletedPostTagByTag(ctx context.Context, tagId int) error {
	query := "DELETE * from post_tags WHERE tag_id=@tag_id"
	_, err := t.db.ExecContext(ctx, query, sql.Named("tag_id", tagId), sql.Named("tag_id", tagId))
	if err != nil {
		return err
	}
	return nil
}
