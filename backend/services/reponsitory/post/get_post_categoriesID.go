package post

import (
	entity "bloghomnay-project/services/entity/posts"
	"context"
	"database/sql"
)

// Lọc post qua Categories
func (p *PostServiceSQL) GetPostByCategories(ctx context.Context, categoryID int) ([]entity.Posts, error) {
	query := "SELECT id, user_id, category_id, title, content FROM posts where category_id = @category_id"
	rows, err := p.db.QueryContext(ctx, query, sql.Named("category_id", categoryID))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var listPosts []entity.Posts
	for rows.Next() {
		var data entity.Posts
		if err := rows.Scan(&data.Id, &data.UserID, &data.CategoryId, &data.Title, &data.Content); err != nil {
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
