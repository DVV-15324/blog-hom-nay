package posttag

import (
	"context"
	"database/sql"
)

func (p *PostTagServiceSQL) UpdatePostTags(cxt context.Context, postID int, tagId int) error {
	query := "UPDATE post_tags SET tag_id=@tag_id Where post_id=@post_id"

	_, err := p.db.ExecContext(cxt, query, sql.Named("tag_id", tagId), sql.Named("post_id", postID))
	if err != nil {
		return err
	}
	return nil
}
