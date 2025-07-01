package postlikes

import (
	entity "bloghomnay-project/services/entity/postlikes"
	"context"
	"database/sql"
)

func (c *PostLikeServiceSQL) GetPostLike(ctx context.Context, userID int, postId int) (bool, error) {
	var data entity.PostLike
	query := `
		SELECT id, user_id, post_id, created_at
		FROM post_likes
		WHERE user_id = @user_id AND post_id = @post_id
	`
	row := c.db.QueryRowContext(ctx, query,
		sql.Named("user_id", userID),
		sql.Named("post_id", postId),
	)

	err := row.Scan(&data.Id, &data.UserId, &data.PostId, &data.CreatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			// Không tìm thấy -> chưa like
			return false, nil
		}
		// Có lỗi khác
		return false, err
	}

	return true, nil
}
