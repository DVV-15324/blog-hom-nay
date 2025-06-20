package post

import (
	entity "bloghomnay-project/services/entity/posts"
	"context"
	"database/sql"
)

// Tìm kiếm post qua tiêu đề title
func (t *PostServiceSQL) GetPostByTitles(ctx context.Context, title string) ([]entity.Posts, error) {
<<<<<<< HEAD
	query := `SELECT 
		p.id, 
		p.user_id, 
		p.category_id, 
		p.description, 
		p.title, 
		p.content,
		COUNT(DISTINCT pl.id) AS [like],
		COUNT(DISTINCT c.id) AS [count_comment]
		FROM posts AS p
		LEFT JOIN post_likes AS pl ON pl.post_id = p.id
		LEFT JOIN comments AS c ON c.post_id = p.id
		WHERE p.content LIKE @title
		GROUP BY 
			p.id, p.user_id, p.category_id, p.description, p.title, p.content;`
=======
	query := "SELECT id, user_id, category_id, title, content FROM posts WHERE title LIKE @title"
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
	likePattern := "%" + title + "%" // Tìm tiêu đề chứa từ khóa

	rows, err := t.db.QueryContext(ctx, query, sql.Named("title", likePattern))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var listPosts []entity.Posts
	for rows.Next() {
		var data entity.Posts
<<<<<<< HEAD
		if err := rows.Scan(&data.Id, &data.UserID, &data.CategoryId, &data.Description, &data.Title, &data.Content, &data.Like, &data.CountComment); err != nil {
=======
		if err := rows.Scan(&data.Id, &data.UserID, &data.CategoryId, &data.Title, &data.Content); err != nil {
>>>>>>> c821afe7457cacaa8d68fb4598eecf76a42272b8
			return nil, err
		}
		listPosts = append(listPosts, data)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return listPosts, nil
}
