package comment

import (
	entity "bloghomnay-project/services/entity/comment"
	"context"
	"database/sql"
)

<<<<<<< HEAD
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
=======
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
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
}
