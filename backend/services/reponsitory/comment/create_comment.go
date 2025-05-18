package comment

import (
	entity "bloghomnay-project/services/entity/comment"
	"context"
	"database/sql"
)

func (c *CommentServiceSQL) CreateComment(cxt context.Context, comment *entity.CreateComment) error {
	query := `INSERT INTO comment(post_id, user_id, content)
			values(@post_id, @user_id, @content);`
	_, err := c.db.ExecContext(cxt, query,
		sql.Named("post_id", comment.Content),
	)
	if err != nil {
		return err
	}
	return nil
}
