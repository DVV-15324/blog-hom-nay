package comment

import (
	entity "bloghomnay-project/services/entity/comment"
	"context"
	"database/sql"
)

// Get nhiều comment của 1 post
func (c *CommentServiceSQL) GetNotification(ctx context.Context, userID int) ([]entity.Comment, error) {
	query := "SELECT id, user_id ,post_id ,content FROM comments WHERE user_id=@user_id"
	rows, err := c.db.QueryContext(ctx, query, sql.Named("user_id", userID))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var listPostLike []entity.Comment
	for rows.Next() {
		var data entity.Comment
		if err := rows.Scan(&data.Id, &data.UserId, &data.PostId, &data.Content); err != nil {
			return nil, err
		}
		listPostLike = append(listPostLike, data)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return listPostLike, nil
}
