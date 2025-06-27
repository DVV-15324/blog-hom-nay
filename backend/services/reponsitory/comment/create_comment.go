package comment

import (
	entity "bloghomnay-project/services/entity/comment"
	"context"
	"database/sql"
)

func (c *CommentServiceSQL) CreateComment(cxt context.Context, comment *entity.CreateComment) (*entity.Comment, error) {
	var newComment entity.Comment

	query := `
        INSERT INTO comments(post_id, user_id, content)
        OUTPUT inserted.id, inserted.post_id, inserted.user_id, inserted.content, inserted.created_at
        VALUES (@post_id, @user_id, @content);
    `

	err := c.db.QueryRowContext(cxt, query,
		sql.Named("post_id", comment.PostID),
		sql.Named("user_id", comment.UserID),
		sql.Named("content", comment.Content),
	).Scan(&newComment.Id, &newComment.PostId, &newComment.UserId, &newComment.Content, &newComment.CreatedAt)

	if err != nil {
		return nil, err
	}

	return &newComment, nil
}
