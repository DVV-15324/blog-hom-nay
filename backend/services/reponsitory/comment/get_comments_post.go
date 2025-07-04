package comment

import (
	entity "bloghomnay-project/services/entity/comment"
	"context"
	"database/sql"
)

// Get nhiều comment của 1 post
func (c *CommentServiceSQL) GetCommentByPostId(ctx context.Context, postID int) ([]entity.Comment, error) {
	query := "SELECT id, post_id, user_id, content, created_at FROM comments where post_id = @post_id"
	rows, err := c.db.QueryContext(ctx, query, sql.Named("post_id", postID))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var listComment []entity.Comment
	for rows.Next() {
		var data entity.Comment
		if err := rows.Scan(&data.Id, &data.PostId, &data.UserId, &data.Content, &data.CreatedAt); err != nil {
			return nil, err
		}
		listComment = append(listComment, data)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return listComment, nil
}
