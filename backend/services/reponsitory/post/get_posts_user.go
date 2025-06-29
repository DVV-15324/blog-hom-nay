package post

import (
	entity "bloghomnay-project/services/entity/posts"
	"context"
	"database/sql"
)

// Lấy posts của riêng user
func (t *PostServiceSQL) GetPostByUserId(ctx context.Context, id int) ([]entity.Posts, error) {
	query := `SELECT 
		p.id, 
		p.user_id, 
		p.category_id, 
		p.description, 
		p.title, 
		p.content,
		COUNT(pl.id) AS [like],
		COUNT(c.id) AS [count_comment],
		p.created_at,
		p.updated_at
		FROM posts AS p
		LEFT JOIN post_likes AS pl ON pl.post_id = p.id
		LEFT JOIN comments AS c ON c.post_id = p.id
		WHERE p.user_id = @user_id
		GROUP BY 
			p.id, p.user_id, p.category_id, p.description, p.title, p.content, p.created_at, p.updated_at;`
	rows, err := t.db.QueryContext(ctx, query, sql.Named("user_id", id))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var listPosts []entity.Posts
	for rows.Next() {
		var data entity.Posts
		if err := rows.Scan(&data.Id, &data.UserID, &data.CategoryId, &data.Description, &data.Title, &data.Content, &data.Like, &data.CountComment, &data.CreatedAt, &data.UpdatedAt); err != nil {
			return nil, err
		}
		listPosts = append(listPosts, data)
	}

	// Check lỗi sau khi duyệt xong (ví dụ: lỗi mạng trong lúc đọc)
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return listPosts, nil
}
