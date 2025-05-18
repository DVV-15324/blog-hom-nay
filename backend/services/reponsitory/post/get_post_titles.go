package post

import (
	entity "bloghomnay-project/services/entity/posts"
	"context"
	"database/sql"
)

// Tìm kiếm post qua tiêu đề title
func (t *PostServiceSQL) GetPostByTitles(ctx context.Context, title string) ([]entity.Posts, error) {
	query := "SELECT id, user_id, category_id, title, content FROM posts WHERE title LIKE @title"
	likePattern := "%" + title + "%" // Tìm tiêu đề chứa từ khóa

	rows, err := t.db.QueryContext(ctx, query, sql.Named("title", likePattern))
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

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return listPosts, nil
}
