package post

import (
	entity "bloghomnay-project/services/entity/posts"
	"context"
)

// Get tất cả posts HOT(sắp xếp theo mới nhất)
func (t *PostServiceSQL) GetPosts(ctx context.Context) ([]entity.Posts, error) {
	query := "SELECT id, user_id, category_id, title, content FROM posts"
	rows, err := t.db.QueryContext(ctx, query)
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
